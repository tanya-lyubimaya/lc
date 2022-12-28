package lc

func removeDuplicatesNaive(nums []int) int {
	i := 1
	k := len(nums)
	for i < len(nums) && k > i {
		if nums[i] == nums[i-1] {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
			k--
		} else {
			i++
		}
	}
	return k
}

func removeDuplicatesClassic(nums []int) int {
	k := len(nums)
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == nums[i+1] {
			for j := i + 1; j < len(nums); j++ {
				nums[j-1] = nums[j]
			}
			k--
		}
	}
	return k
}

func removeDuplicates(nums []int) int {
	// Check for edge cases.
	if len(nums) == 0 {
		return 0
	}

	// Use the two pointer technique to remove the duplicates in-place.
	// The first element shouldn't be touched; it's already in its correct place.
	writePointer := 1
	// Go through each element in the Array.
	for readPointer := 1; readPointer < len(nums); readPointer++ {
		// If the current element we're reading is *different* to the previous
		// element...
		if nums[readPointer] != nums[readPointer-1] {
			// Copy it into the next position at the front, tracked by writePointer.
			nums[writePointer] = nums[readPointer]
			// And we need to now increment writePointer, because the next element
			// should be written one space over.
			writePointer++
		}
	}

	// This turns out to be the correct length value.
	return writePointer
}
