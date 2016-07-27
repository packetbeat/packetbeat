package reader

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/elastic/beats/libbeat/common"
)

// MultiLine reader combining multiple line events into one multi-line event.
//
// Lines to be combined are matched by some configurable predicate using
// regular expression.
//
// The maximum number of bytes and lines to be returned is fully configurable.
// Even if limits are reached subsequent lines are matched, until event is
// fully finished.
//
// Errors will force the multiline reader to return the currently active
// multiline event first and finally return the actual error on next call to Next.
type Multiline struct {
	reader    Reader
	pred      matcher
	maxBytes  int // bytes stored in content
	maxLines  int
	separator []byte
	ts        time.Time
	content   []byte
	last      []byte
	readBytes int // bytes as read from input source
	numLines  int
	fields    common.MapStr
	err       error // last seen error
	state     func(*Multiline) (Message, error)
	done      chan struct{}
}

const (
	// Default maximum number of lines to return in one multi-line event
	defaultMaxLines = 500

	// Default timeout to finish a multi-line event.
	defaultMultilineTimeout = 5 * time.Second
)

// Matcher represents the predicate comparing any two lines
// to find start and end of multiline events in stream of line events.
type matcher func(last, current []byte) bool

var (
	errMultilineTimeout = errors.New("multline timeout")
)

// NewMultiline creates a new multi-line reader combining stream of
// line events into stream of multi-line events.
func NewMultiline(
	r Reader,
	separator string,
	maxBytes int,
	config *MultilineConfig,
) (*Multiline, error) {
	types := map[string]func(*regexp.Regexp) (matcher, error){
		"before": beforeMatcher,
		"after":  afterMatcher,
	}

	matcherType, ok := types[config.Match]
	if !ok {
		return nil, fmt.Errorf("unknown matcher type: %s", config.Match)
	}

	matcher, err := matcherType(config.Pattern)
	if err != nil {
		return nil, err
	}

	if config.Negate {
		matcher = negatedMatcher(matcher)
	}

	maxLines := defaultMaxLines
	if config.MaxLines != nil {
		maxLines = *config.MaxLines
	}

	timeout := defaultMultilineTimeout
	if config.Timeout != nil {
		timeout = *config.Timeout
		if timeout < 0 {
			return nil, fmt.Errorf("timeout %v must not be negative", config.Timeout)
		}
	}

	if timeout > 0 {
		r = NewTimeout(r, errMultilineTimeout, timeout)
	}

	mlr := &Multiline{
		reader:    r,
		pred:      matcher,
		state:     (*Multiline).readFirst,
		maxBytes:  maxBytes,
		maxLines:  maxLines,
		separator: []byte(separator),
		done:      make(chan struct{}),
	}
	return mlr, nil
}

// Next returns next multi-line event.
func (mlr *Multiline) Next() (Message, error) {
	return mlr.state(mlr)
}

func (r *Multiline) Close() error {
	close(r.done)
	return r.reader.Close()
}

func (mlr *Multiline) readFirst() (Message, error) {
	for {
		select {
		case <-mlr.done:
			return Message{}, ErrReaderStopped
		default:
		}
		p, err := mlr.reader.Next()
		if err == nil {
			if p.Bytes == 0 {
				continue
			}

			mlr.startNewLine(p)
			mlr.state = (*Multiline).readNext
			return mlr.readNext()
		}

		// no lines buffered -> ignore timeout
		if err == errMultilineTimeout {
			continue
		}

		// something is wrong here
		return p, err
	}
}

func (mlr *Multiline) readNext() (Message, error) {
	for {
		select {
		case <-mlr.done:
			return Message{}, ErrReaderStopped
		default:
		}
		p, err := mlr.reader.Next()

		if err != nil {
			// handle multiline timeout signal
			if err == errMultilineTimeout {
				// no lines buffered -> ignore timeout
				if mlr.numLines == 0 {
					continue
				}

				// return collected multiline event and
				// empty buffer for new multiline event
				p := mlr.pushLine()
				mlr.reset()
				return p, nil
			}

			// handle error without any bytes returned from reader
			if p.Bytes == 0 {
				// no lines buffered -> return error
				if mlr.numLines == 0 {
					return Message{}, err
				}

				// lines buffered, return multiline and error on next read
				p := mlr.pushLine()
				mlr.err = err
				mlr.state = (*Multiline).readFailed
				return p, nil
			}

			// handle error with some content being returned by reader and
			// line matching multiline criteria or no multiline started yet
			if mlr.readBytes == 0 || mlr.pred(mlr.last, p.Content) {
				mlr.addLine(p)

				// return multiline and error on next read
				l := mlr.pushLine()
				mlr.err = err
				mlr.state = (*Multiline).readFailed
				return l, nil
			}

			// no match, return current multline and retry with current line on next
			// call to readNext awaiting the error being reproduced (or resolved)
			// in next call to Next
			l := mlr.startNewLine(p)
			return l, nil
		}

		// if predicate does not match current multiline -> return multiline event
		if mlr.readBytes > 0 && !mlr.pred(mlr.last, p.Content) {
			l := mlr.startNewLine(p)
			return l, nil
		}

		// add line to current multiline event
		mlr.addLine(p)
	}
}

func (mlr *Multiline) readFailed() (Message, error) {
	// return error and reset line reader
	err := mlr.err
	mlr.err = nil
	mlr.reset()
	return Message{}, err
}

func (mlr *Multiline) startNewLine(l Message) Message {
	retLine := mlr.pushLine()
	mlr.addLine(l)
	mlr.ts = l.Ts
	mlr.fields = l.Fields
	return retLine
}

func (mlr *Multiline) pushLine() Message {
	content := mlr.content
	sz := mlr.readBytes
	fields := mlr.fields

	mlr.content = nil
	mlr.last = nil
	mlr.readBytes = 0
	mlr.numLines = 0
	mlr.err = nil
	mlr.fields = nil

	return Message{Ts: mlr.ts, Content: content, Fields: fields, Bytes: sz}
}

func (mlr *Multiline) addLine(m Message) {
	if m.Bytes <= 0 {
		return
	}

	sz := len(mlr.content)
	addSeparator := len(mlr.content) > 0 && len(mlr.separator) > 0
	if addSeparator {
		sz += len(mlr.separator)
	}

	space := mlr.maxBytes - sz
	spaceLeft := (mlr.maxBytes <= 0 || space > 0) &&
		(mlr.maxLines <= 0 || mlr.numLines < mlr.maxLines)
	if spaceLeft {
		if space < 0 || space > len(m.Content) {
			space = len(m.Content)
		}

		tmp := mlr.content
		if addSeparator {
			tmp = append(tmp, mlr.separator...)
		}
		mlr.content = append(tmp, m.Content[:space]...)
		mlr.numLines++
	}

	mlr.last = m.Content
	mlr.readBytes += m.Bytes
}

func (mlr *Multiline) reset() {
	mlr.state = (*Multiline).readFirst
}

// matchers

func afterMatcher(regex *regexp.Regexp) (matcher, error) {
	return genPatternMatcher(regex, func(last, current []byte) []byte {
		return current
	})
}

func beforeMatcher(regex *regexp.Regexp) (matcher, error) {
	return genPatternMatcher(regex, func(last, current []byte) []byte {
		return last
	})
}

func negatedMatcher(m matcher) matcher {
	return func(last, current []byte) bool {
		return !m(last, current)
	}
}

func genPatternMatcher(
	regex *regexp.Regexp,
	sel func(last, current []byte) []byte,
) (matcher, error) {
	matcher := func(last, current []byte) bool {
		line := sel(last, current)
		return regex.Match(line)
	}
	return matcher, nil
}
