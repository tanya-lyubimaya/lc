package lc

func twoSum(nums []int, target int) []int {
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

// TODO: Add two-pass hash table solution
// TODO: Add one-pass hash table solution
