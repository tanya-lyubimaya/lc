package lc

func twoSumBruteForce(nums []int, target int) []int {
	var results []int
	for i, v := range nums {
		for j, v2 := range nums {
			if i == j {
				continue
			}
			if target-v == v2 {
				results = append(results, i)
				results = append(results, j)
				return results
			}
		}
	}
	return results
}

func twoSumHashMapTwoPass(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, v := range nums {
		hashmap[v] = i
	}
	for i, v := range nums {
		complement := target - v
		if index, ok := hashmap[complement]; ok && index != i {
			return []int{i, hashmap[complement]}
		}
	}
	return []int{-1, -1}
}

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, v := range nums {
		complement := target - v
		if index, ok := hashmap[complement]; ok && index != i {
			return []int{i, hashmap[complement]}
		}
		hashmap[v] = i
	}
	return []int{-1, -1}
}
