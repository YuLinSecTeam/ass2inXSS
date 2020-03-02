package core

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

func SubStr1(s string, length int) string {
	n := 0
	for i := 0; i < length && n < len(s); i++ {
		_, size := utf8.DecodeRuneInString(s[n:])
		n += size
	}

	return s[n:]
}

func Implode(glue string, pieces []string) string {
	var buf bytes.Buffer
	l := len(pieces)
	for _, str := range pieces {
		buf.WriteString(str)
		if l--; l > 0 {
			buf.WriteString(glue)
		}
	}
	return buf.String()
}

func Stripos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	haystack = haystack[offset:]
	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack, needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}

func Substr(str string, start int, length int) string {
	if start < 0 || length < -1 {
		return str
	}
	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}
	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}

