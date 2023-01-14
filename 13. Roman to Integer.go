package lc

import "strings"

func romanToInt(s string) int {
	var res int
	i := len(s) - 1
	for i > 0 && len(s) > 1 {
		switch s[i] {
		case 'I':
			res += 1
			i--
		case 'V':
			if s[i-1] == 'I' {
				res += 4
				i -= 2
			} else {
				res += 5
				i--
			}
		case 'X':
			if s[i-1] == 'I' {
				res += 9
				i -= 2
			} else {
				res += 10
				i--
			}
		case 'L':
			if s[i-1] == 'X' {
				res += 40
				i -= 2
			} else {
				res += 50
				i--
			}
		case 'C':
			if s[i-1] == 'X' {
				res += 90
				i -= 2
			} else {
				res += 100
				i--
			}
		case 'D':
			if s[i-1] == 'C' {
				res += 400
				i -= 2
			} else {
				res += 500
				i--
			}
		case 'M':
			if s[i-1] == 'C' {
				res += 900
				i -= 2
			} else {
				res += 1000
				i--
			}
		}
	}
	if i == 0 {
		switch s[0] {
		case 'I':
			res += 1
		case 'V':
			res += 5
		case 'X':
			res += 10
		case 'L':
			res += 50
		case 'C':
			res += 100
		case 'D':
			res += 500
		case 'M':
			res += 1000
		}
	}
	return res
}

func romanToIntButcheredHashMap(s string) int {
	var res int
	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	s = strings.Replace(s, "IV", "IIII", -1)
	s = strings.Replace(s, "IX", "VIIII", -1)
	s = strings.Replace(s, "XL", "XXXX", -1)
	s = strings.Replace(s, "XC", "LXXXX", -1)
	s = strings.Replace(s, "CD", "CCCC", -1)
	s = strings.Replace(s, "CM", "DCCCC", -1)
	for _, c := range s {
		res += values[c]
	}
	return res
}
