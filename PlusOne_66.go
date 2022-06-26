package main

func plusOne(digits []int) []int {
	ind := len(digits) - 1
	for ind >= 0 {
		if digits[ind] == 9 {
			digits[ind] = 0
			ind--
		} else {
			digits[ind]++
			ind = -2
		}
	}
	if ind == -1 {
		digits = append([]int{1}, digits...)
	}
	return digits

}
