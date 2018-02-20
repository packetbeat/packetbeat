package common

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
)

// Byte order utilities

func BytesNtohs(b []byte) uint16 {
	return uint16(b[0])<<8 | uint16(b[1])
}

func BytesNtohl(b []byte) uint32 {
	return uint32(b[0])<<24 | uint32(b[1])<<16 |
		uint32(b[2])<<8 | uint32(b[3])
}

func BytesHtohl(b []byte) uint32 {
	return uint32(b[3])<<24 | uint32(b[2])<<16 |
		uint32(b[1])<<8 | uint32(b[0])
}

func BytesNtohll(b []byte) uint64 {
	return uint64(b[0])<<56 | uint64(b[1])<<48 |
		uint64(b[2])<<40 | uint64(b[3])<<32 |
		uint64(b[4])<<24 | uint64(b[5])<<16 |
		uint64(b[6])<<8 | uint64(b[7])
}

// Ipv4_Ntoa transforms an IP4 address in it's dotted notation
func IPv4Ntoa(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16),
		byte(ip>>8), byte(ip))
}

// ReadString extracts the first null terminated string from
// a slice of bytes.
func ReadString(s []byte) (string, error) {
	i := bytes.IndexByte(s, 0)
	if i < 0 {
		return "", errors.New("No string found")
	}
	res := string(s[:i])
	return res, nil
}

// RandomBytes return a slice of random bytes of the defined length
func RandomBytes(length int) ([]byte, error) {
	r := make([]byte, length)
	_, err := rand.Read(r)

	if err != nil {
		return nil, err
	}

	return r, nil
}
