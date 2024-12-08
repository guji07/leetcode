package submissions

func removeDuplicates(nums []int) int {
	i := 0
	j := 0

	for ; i < len(nums); i++ {
		nums[j] = nums[i]
		for ; i < len(nums)-1 && nums[i] == nums[i+1]; i++ {
		}
		j++
	}

	return j
}
