package elasticsearch

import (
	"errors"

	"github.com/elastic/beats/libbeat/common/streambuf"
)

// SAX like json parser. But instead of relying on callbacks, state machine
// returns raw item plus entity. On top of state machine additional helper methods
// like expectDict, expectArray, nextFieldName and nextInt are available for
// low-level parsing/stepping through a json document.
//
// Due to parser simply stepping through the input buffer, almost no additional
// allocations are required.
type jsonReader struct {
	streambuf.Buffer

	// parser state machine
	states       []state // state stack for nested arrays/objects
	currentState state

	// preallocate stack memory for up to 32 nested arrays/objects
	statesBuf [32]state
}

var (
	errFailing             = errors.New("JSON parser failed")
	errUnknownChar         = errors.New("unknown character")
	errQuoteMissing        = errors.New("missing closing quote")
	errExpectColon         = errors.New("expected ':' after map key")
	errUnexpectedDictClose = errors.New("unexpected '}'")
	errUnexpectedArrClose  = errors.New("unexpected ']'")
	errExpectedDigit       = errors.New("expected a digit")
	errExpectedObject      = errors.New("expected JSON object")
	errExpectedArray       = errors.New("expected JSON array")
	errExpectedFieldName   = errors.New("expected JSON object field name")
	errExpectedInteger     = errors.New("expected integer value")
	errExpectedNull        = errors.New("expected null value")
	errExpectedFalse       = errors.New("expected false value")
	errExpectedTrue        = errors.New("expected true value")
	errExpectedArrayField  = errors.New("expected ']' or ','")
)

var (
	nullSymbol  = []byte("null")
	trueSymbol  = []byte("true")
	falseSymbol = []byte("false")
)

type entity uint8

const (
	failEntity entity = iota
	trueValue
	falseValue
	nullValue
	dictStart
	dictEnd
	dictField
	arrStart
	arrEnd
	stringEntity
	mapKeyEntity
	intEntity
	doubleEntity
)

type state uint8

const (
	failedState state = iota
	startState
	arrState
	arrStateNext
	dictState
	dictFieldState
	dictFieldStateEnd
)

func (s state) String() string {
	switch s {
	case failedState:
		return "failed"
	case startState:
		return "start"
	case arrState:
		return "array"
	case arrStateNext:
		return "arrayNext"
	case dictState:
		return "dict"
	case dictFieldState:
		return "dictValue"
	case dictFieldStateEnd:
		return "dictNext"
	}
	return "unknown"
}

func newJSONReader(in []byte) *jsonReader {
	r := &jsonReader{}
	r.init(in)
	return r
}

func (r *jsonReader) init(in []byte) {
	r.Buffer.Init(in, true)
	r.currentState = startState
	r.states = r.statesBuf[:0]
}

var whitespace = []byte(" \t\r\n")

func (r *jsonReader) skipWS() {
	r.IgnoreSymbols(whitespace)
}

func (r *jsonReader) pushState(next state) {
	if r.currentState != failedState {
		r.states = append(r.states, r.currentState)
	}
	r.currentState = next
}

func (r *jsonReader) popState() {
	if len(r.states) == 0 {
		r.currentState = failedState
	} else {
		last := len(r.states) - 1
		r.currentState = r.states[last]
		r.states = r.states[:last]
	}
}

func (r *jsonReader) expectDict() error {
	entity, _, err := r.step()
	if err != nil {
		return err
	}

	if entity != dictStart {
		return r.SetError(errExpectedObject)
	}

	return nil
}

func (r *jsonReader) expectArray() error {
	entity, _, err := r.step()
	if err != nil {
		return err
	}

	if entity != arrStart {
		return r.SetError(errExpectedArray)
	}

	return nil
}

func (r *jsonReader) nextFieldName() (entity, []byte, error) {
	entity, raw, err := r.step()
	if err != nil {
		return entity, raw, err
	}

	if entity != mapKeyEntity && entity != dictEnd {
		return entity, nil, r.SetError(errExpectedFieldName)
	}

	return entity, raw, err
}

func (r *jsonReader) nextInt() (int, error) {
	entity, raw, err := r.step()
	if err != nil {
		return 0, err
	}

	if entity != intEntity {
		return 0, errExpectedInteger
	}

	tmp := streambuf.NewFixed(raw)
	i, err := tmp.AsciiInt(false)
	return int(i), err
}

// ignore type of next element and return raw content.
func (r *jsonReader) ignoreNext() (raw []byte, err error) {
	r.skipWS()

	snapshot := r.Snapshot()
	before := r.Len()

	var ignoreKind func(*jsonReader, entity) error
	ignoreKind = func(r *jsonReader, kind entity) error {

		for {
			entity, _, err := r.step()
			if err != nil {
				return err
			}

			switch entity {
			case kind:
				return nil
			case arrStart:
				return ignoreKind(r, arrEnd)
			case dictStart:
				return ignoreKind(r, dictEnd)
			}
		}
	}

	entity, _, err := r.step()
	if err != nil {
		return nil, err
	}

	switch entity {
	case dictStart:
		err = ignoreKind(r, dictEnd)
	case arrStart:
		err = ignoreKind(r, arrEnd)
	}
	if err != nil {
		return nil, err
	}

	after := r.Len()
	r.Restore(snapshot)

	bytes, _ := r.Collect(before - after)
	return bytes, nil
}

// step continues the JSON parser state machine until next entity has been parsed.
func (r *jsonReader) step() (entity, []byte, error) {
	r.skipWS()
	switch r.currentState {
	case failedState:
		return r.stepFailing()
	case startState:
		return r.stepStart()
	case arrState:
		return r.stepArray()
	case arrStateNext:
		return r.stepArrayNext()
	case dictState:
		return r.stepDict()
	case dictFieldState:
		return r.stepDictValue()
	case dictFieldStateEnd:
		return r.stepDictValueEnd()
	default:
		return r.failWith(errFailing)
	}
}

func (r *jsonReader) stepFailing() (entity, []byte, error) {
	return failEntity, nil, r.Err()
}

func (r *jsonReader) stepStart() (entity, []byte, error) {
	c, err := r.PeekByte()
	if err != nil {
		return r.failWith(err)
	}

	return r.tryStepPrimitive(c)
}

func (r *jsonReader) stepArray() (entity, []byte, error) {
	return r.doStepArray(true)
}

func (r *jsonReader) stepArrayNext() (entity, []byte, error) {
	c, err := r.PeekByte()
	if err != nil {
		return r.failWith(errFailing)
	}

	switch c {
	case ']':
		return r.endArray()
	case ',':
		r.Advance(1)
		r.skipWS()
		r.currentState = arrState
		return r.doStepArray(false)
	default:
		return r.failWith(errExpectedArrayField)
	}
}

func (r *jsonReader) doStepArray(allowArrayEnd bool) (entity, []byte, error) {
	c, err := r.PeekByte()
	if err != nil {
		return r.failWith(err)
	}

	if c == ']' {
		if !allowArrayEnd {
			return r.failWith(errUnexpectedArrClose)
		}
		return r.endArray()
	}

	r.currentState = arrStateNext
	return r.tryStepPrimitive(c)
}

func (r *jsonReader) stepDict() (entity, []byte, error) {
	return r.doStepDict(true)
}

func (r *jsonReader) doStepDict(allowEnd bool) (entity, []byte, error) {
	c, err := r.PeekByte()
	if err != nil {
		return r.failWith(err)
	}

	switch c {
	case '}':
		if !allowEnd {
			return r.failWith(errUnexpectedDictClose)
		}
		return r.endDict()
	case '"':
		r.currentState = dictFieldState
		return r.stepMapKey()
	default:
		return r.failWith(errExpectedFieldName)
	}
}

func (r *jsonReader) stepDictValue() (entity, []byte, error) {
	c, err := r.PeekByte()
	if err != nil {
		return r.failWith(err)
	}

	r.currentState = dictFieldStateEnd
	return r.tryStepPrimitive(c)
}

func (r *jsonReader) stepDictValueEnd() (entity, []byte, error) {
	c, err := r.PeekByte()
	if err != nil {
		return r.failWith(err)
	}

	switch c {
	case '}':
		return r.endDict()
	case ',':
		r.Advance(1)
		r.skipWS()
		r.currentState = dictState
		return r.doStepDict(false)
	default:
		return r.failWith(errUnknownChar)
	}
}

func (r *jsonReader) tryStepPrimitive(c byte) (entity, []byte, error) {
	switch c {
	case '{': // start dictionary
		return r.startDict()
	case '[': // start array
		return r.startArray()
	case 'n': // null
		return r.stepNull()
	case 'f': // false
		return r.stepFalse()
	case 't': // true
		return r.stepTrue()
	case '"':
		return r.stepString()
	default:
		// parse number?
		if c == '-' || c == '+' || c == '.' || ('0' <= c && c <= '9') {
			return r.stepNumber()
		}

		err := r.Err()
		if err == nil {
			err = r.SetError(errUnknownChar)
		}
		return r.failWith(err)
	}
}

func (r *jsonReader) stepNull() (entity, []byte, error) {
	return stepSymbol(r, nullValue, nullSymbol, errExpectedNull)
}

func (r *jsonReader) stepTrue() (entity, []byte, error) {
	return stepSymbol(r, trueValue, trueSymbol, errExpectedTrue)
}

func (r *jsonReader) stepFalse() (entity, []byte, error) {
	return stepSymbol(r, falseValue, falseSymbol, errExpectedFalse)
}

func stepSymbol(r *jsonReader, e entity, symb []byte, fail error) (entity, []byte, error) {
	ok, err := r.AsciiMatch(symb)
	if err != nil {
		return failEntity, nil, err
	}
	if !ok {
		return failEntity, nil, fail
	}

	r.Advance(len(symb))
	return e, nil, nil
}

func (r *jsonReader) stepMapKey() (entity, []byte, error) {
	entity, key, err := r.stepString()
	if err != nil {
		return entity, key, err
	}

	r.skipWS()
	c, err := r.ReadByte()
	if err != nil {
		return failEntity, nil, err
	}

	if c != ':' {
		return r.failWith(r.SetError(errExpectColon))
	}

	if err := r.Err(); err != nil {
		return r.failWith(err)
	}
	return mapKeyEntity, key, nil
}

func (r *jsonReader) stepString() (entity, []byte, error) {
	start := 1
	for {
		idxQuote := r.IndexByteFrom(start, '"')
		if idxQuote == -1 {
			return failEntity, nil, r.SetError(errQuoteMissing)
		}

		if b, _ := r.PeekByteFrom(idxQuote - 1); b == '\\' { // escaped quote?
			start = idxQuote + 1
			continue
		}

		// found string end
		str, err := r.Collect(idxQuote + 1)
		str = str[1 : len(str)-1]
		return stringEntity, str, err
	}
}

func (r *jsonReader) startDict() (entity, []byte, error) {
	r.Advance(1)
	r.pushState(dictState)
	return dictStart, nil, nil
}

func (r *jsonReader) endDict() (entity, []byte, error) {
	r.Advance(1)
	r.popState()
	return dictEnd, nil, nil
}

func (r *jsonReader) startArray() (entity, []byte, error) {
	r.Advance(1)
	r.pushState(arrState)
	return arrStart, nil, nil
}

func (r *jsonReader) endArray() (entity, []byte, error) {
	r.Advance(1)
	r.popState()
	return arrEnd, nil, nil
}

func (r *jsonReader) failWith(err error) (entity, []byte, error) {
	r.currentState = failedState
	return failEntity, nil, r.SetError(err)
}

func (r *jsonReader) stepNumber() (entity, []byte, error) {
	snapshot := r.Snapshot()
	lenBefore := r.Len()
	isDouble := false

	if err := r.Err(); err != nil {
		return failEntity, nil, err
	}

	// parse '+', '-' or '.'
	if b, _ := r.PeekByte(); b == '-' || b == '+' {
		r.Advance(1)
	}
	if b, _ := r.PeekByte(); b == '.' {
		r.Advance(1)
		isDouble = true
	}

	// parse digits
	buf, _ := r.CollectWhile(isDigit)
	if len(buf) == 0 {
		return failEntity, nil, r.SetError(errExpectedDigit)
	}

	if !isDouble {
		// parse optional '.'
		if b, _ := r.PeekByte(); b == '.' {
			r.Advance(1)
			isDouble = true

			// parse optional digits
			r.CollectWhile(isDigit)
		}
	}

	lenAfter := r.Len()
	r.Restore(snapshot)
	total := lenBefore - lenAfter - 1
	if total == 0 {
		return failEntity, nil, r.SetError(errExpectedDigit)
	}

	raw, _ := r.Collect(total)
	state := intEntity
	if isDouble {
		state = doubleEntity
	}

	return state, raw, nil
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
