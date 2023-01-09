package lc

import "sort"

func isAnagramHashMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	smap := make(map[rune]int)
	for _, c := range s {
		smap[c]++
	}
	for _, c := range t {
		if _, ok := smap[c]; ok {
			smap[c]--
		}
	}
	for _, v := range smap {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagramSort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sRunes, tRunes := []rune(s), []rune(t)
	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] < sRunes[j]
	})
	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] < tRunes[j]
	})
	for i := 0; i < len(sRunes); i++ {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sRunes, tRunes := []rune(s), []rune(t)
	var sum int
	for i := 0; i < len(sRunes); i++ {
		sum += 262144/int(sRunes[i]) - 262144/int(tRunes[i])
	}
	if sum != 0 {
		return false
	}
	return true
}
