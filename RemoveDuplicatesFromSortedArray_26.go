package main

func removeDuplicates(nums []int) int {
	putPosition := 0
	i := 0
	for i < len(nums) {
		nums[putPosition] = nums[i]
		putPosition++
		i++
		for i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
	}
	return putPosition
}
