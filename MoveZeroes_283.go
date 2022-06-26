package main

func moveZeroes(nums []int) {
	movePtr := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != movePtr {
				nums[movePtr] = nums[i]
				nums[i] = 0
			}
			movePtr++
		}
	}
}
