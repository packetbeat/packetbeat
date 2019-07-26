package http

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"unicode/utf8"

	"github.com/elastic/beats/heartbeat/eventext"
	"github.com/elastic/beats/libbeat/common"

	"github.com/elastic/beats/heartbeat/reason"
	"github.com/elastic/beats/libbeat/beat"
)

func handleRespBody(event *beat.Event, resp *http.Response, responseConfig responseConfig, errReason reason.Reason) error {
	defer resp.Body.Close()

	sampleMaxBytes := responseConfig.IncludeBodyMaxBytes

	includeSample := responseConfig.IncludeBody == "always" || (responseConfig.IncludeBody == "on_error" && errReason != nil)

	// No need to return any actual body bytes if we'll discard them anyway. This should save on allocation
	if !includeSample {
		sampleMaxBytes = 0
	}

	sampleStr, bodyBytes, bodyHash, err := readResp(resp, sampleMaxBytes)
	if err != nil {
		return err
	}

	if includeSample {
		addRespBodyFields(event, sampleStr, bodyBytes, bodyHash)
	}

	return nil
}

func addRespBodyFields(event *beat.Event, sampleStr string, bodyBytes int64, bodyHash string) {
	body := common.MapStr{"bytes": bodyBytes}
	if sampleStr != "" {
		body["content"] = sampleStr
	}
	if bodyHash != "" {
		body["hash"] = bodyHash
	}

	eventext.MergeEventFields(event, common.MapStr{"http": common.MapStr{
		"response": common.MapStr{
			"body": body,
		},
	}})
}

// readResp reads the first sampleSize bytes from the httpResponse,
// then closes the body (which closes the connection). It doesn't return any errors
// but does log them. During an error case the return values will be (nil, -1).
// The maxBytes params controls how many bytes will be returned in a string, not how many will be read.
// We always read the full response here since we want to time downloading the full thing.
// This may return a nil body if the response is not valid UTF-8
func readResp(resp *http.Response, maxSampleBytes int) (bodySample string, bodySize int64, hashStr string, err error) {
	if resp == nil {
		return "", -1, "", fmt.Errorf("cannot readResp of nil HTTP response")
	}

	respSize, bodySample, hash, err := readPrefixAndHash(resp.Body, maxSampleBytes)

	return bodySample, respSize, hash, err
}

func readPrefixAndHash(body io.ReadCloser, maxPrefixSize int) (respSize int64, prefix string, hashStr string, err error) {
	hash := sha256.New()
	// Function to lazily get the body of the response
	rawBuf := make([]byte, 1024)

	// Buffer to hold the prefix output along with tracking info
	prefixBuf := make([]byte, maxPrefixSize)
	prefixRemainingBytes := maxPrefixSize
	prefixWriteOffset := 0
	for {
		readSize, readErr := body.Read(rawBuf)

		respSize += int64(readSize)
		hash.Write(rawBuf[:readSize])

		if prefixRemainingBytes > 0 {
			if readSize >= prefixRemainingBytes {
				copy(prefixBuf[prefixWriteOffset:maxPrefixSize-1], rawBuf[:prefixRemainingBytes])
			} else {
				copy(prefixBuf[prefixWriteOffset:prefixWriteOffset+readSize], rawBuf[:readSize])
			}
			prefixRemainingBytes -= readSize
			prefixWriteOffset += readSize
		}

		if readErr == io.EOF {
			break
		}

		if readErr != nil {
			return 0, "", "", readErr
			break
		}
	}
	if utf8.Valid(prefixBuf[:prefixWriteOffset]) {
		prefix = string(prefixBuf[:prefixWriteOffset])
	}
	return respSize, prefix, hex.EncodeToString(hash.Sum(nil)), nil
}
