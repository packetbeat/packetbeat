package xml

import (
	"bytes"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
)

// A SafeReader escapes UTF control characters in the io.Reader
// it wraps, so that it can be fed to Go's xml parser.
// Characters for which `unicode.IsControl` returns true will be output as
// an hexadecimal unicode escape sequence "\\uNNNN".
type SafeReader struct {
	inner   io.Reader
	backing [256]byte
	buf     []byte
	code    []byte
}

var _ io.Reader = (*SafeReader)(nil)

func output(n int) (int, error) {
	if n == 0 {
		return 0, io.EOF
	}
	return n, nil
}

// Read implements the io.Reader interface.
func (r *SafeReader) Read(d []byte) (n int, err error) {
	if len(r.code) > 0 {
		n = copy(d, r.code)
		r.code = r.code[n:]
		return output(n)
	}
	if len(r.buf) == 0 {
		n, _ = r.inner.Read(r.backing[:])
		r.buf = r.backing[:n]
	}
	for i := 0; i < len(r.buf); {
		code, size := utf8.DecodeRune(r.buf[i:])
		if !unicode.IsSpace(code) && unicode.IsControl(code) {
			n = copy(d, r.buf[:i])
			r.buf = r.buf[n+1:]
			r.code = []byte(fmt.Sprintf("\\u%04x", code))
			m := copy(d[n:], r.code)
			r.code = r.code[m:]
			return output(n + m)
		}
		i += size
	}
	n = copy(d, r.buf)
	r.buf = r.buf[n:]
	return output(n)
}

func NewSafeReader(rawXML []byte) *SafeReader {
	return &SafeReader{inner: bytes.NewReader(rawXML)}
}
