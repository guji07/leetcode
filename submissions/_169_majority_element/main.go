package _169_majority_element

func majorityElement(nums []int) int {
	hashmap := make(map[int]int)

	for _, v := range nums {
		hashmap[v]++
	}

	for i, v := range hashmap {
		if v > len(nums)/2 {
			return i
		}
	}

	return -1
}
