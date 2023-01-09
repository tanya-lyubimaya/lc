package lc

import (
	"regexp"
	"strings"
)

func isPalindromeRegexpRunes(s string) bool {
	s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(strings.ToLower(s), "")
	r := []rune(s)
	i := 0
	j := len(r) - 1
	for i < j {
		if r[i] != r[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindromeRegexpString(s string) bool {
	s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(strings.ToLower(s), "")
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindrome(s string) bool {
	var sTrimmed []rune
	for _, c := range []rune(strings.ToLower(s)) {
		if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') {
			sTrimmed = append(sTrimmed, c)
		}
	}
	i, j := 0, len(sTrimmed)-1
	for i < j {
		if sTrimmed[i] != sTrimmed[j] {
			return false
		}
		i++
		j--
	}
	return true
}
