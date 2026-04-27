package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var result strings.Builder
	result.Grow(len(input))

	var buf rune
	var hasBuf bool
	var isEscaped bool

	setBuf := func(r rune) {
		if hasBuf {
			result.WriteRune(buf)
		}
		buf = r
		hasBuf = true
	}

	for _, r := range input {
		if isEscaped {
			if (r < '0' || r > '9') && r != '\\' {
				return "", ErrInvalidString
			}
			setBuf(r)
			isEscaped = false
			continue
		}

		if r == '\\' {
			isEscaped = true
			continue
		}

		num, err := strconv.Atoi(string(r))
		if err != nil {
			setBuf(r)
			continue
		}

		if !hasBuf {
			return "", ErrInvalidString
		}

		result.WriteString(strings.Repeat(string(buf), num))
		hasBuf = false
	}

	if isEscaped {
		return "", ErrInvalidString
	}

	if hasBuf {
		result.WriteRune(buf)
	}

	return result.String(), nil
}
