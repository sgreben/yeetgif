package parse

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

func scanFloat(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading space
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}
	buf := &bytes.Buffer{}
	isNumber := false
	longestMatch := 0
	// Scan until no longer a number
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		buf.WriteRune(r)
		if expFloat.Match(buf.Bytes()) {
			isNumber = true
			longestMatch = i + width
		}
	}
	if isNumber {
		return longestMatch, data[start:longestMatch], nil
	}
	if atEOF && len(data) > start {
		return 0, nil, fmt.Errorf("could not scan float")
	}
	// Request more data.
	return start, nil, nil
}

func scanCommand(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading non-command-codes.
	start := 0
	lastWidth := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if isCommandCode[unicode.ToLower(r)] {
			lastWidth = width
			break
		}
	}
	// Scan until next command-code
	for width, i := 0, start+lastWidth; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if isCommandCode[unicode.ToLower(r)] {
			return i, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final command. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}
