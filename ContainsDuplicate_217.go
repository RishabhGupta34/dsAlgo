package main

func containsDuplicate(nums []int) bool {
	exist := make(map[int]bool, 0)
	for _, num := range nums {
		if _, ok := exist[num]; ok {
			return true
		}
		exist[num] = true
	}
	return false
}
