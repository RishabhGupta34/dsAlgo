package main

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	if k <= 0 {
		return
	}
	count := 0
	for i := 0; count < l; i++ {
		ind := i
		val := nums[ind]
		firstPass := false
		for (!firstPass) || (ind != i) {
			firstPass = true
			tempInd := (ind + k) % l
			temp := nums[tempInd]
			nums[tempInd] = val
			val = temp
			ind = tempInd
			count++
		}
	}
}
