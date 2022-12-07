package lc

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) == 0 && len(magazine) == 0 {
		return true
	}
	if len(ransomNote) > len(magazine) {
		return false
	}

	letters := make(map[string]int)

	for i := 0; i < len(magazine); i++ {
		letters[string(magazine[i])]++
	}

	for i := 0; i < len(ransomNote); i++ {
		// if a letter not in map or the counter == 0
		if _, ok := letters[string(ransomNote[i])]; !ok || letters[string(ransomNote[i])] == 0 {
			return false
		}
		letters[string(ransomNote[i])]--
	}

	return true
}
