package lc

func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[string]int)

	for i := 0; i < len([]rune(magazine)); i++ {
		letters[string(magazine[i])]++
	}

	for i := 0; i < len([]rune(ransomNote)); i++ {
		// if a letter not in map or the counter == 0
		if _, ok := letters[string(ransomNote[i])]; !ok || letters[string(ransomNote[i])] == 0 {
			return false
		}
		letters[string(ransomNote[i])]--
	}

	return true
}
