package main

func intersect(nums1 []int, nums2 []int) []int {
	exists := make(map[int]int, 0)
	if len(nums1) < len(nums2) {
		for _, val := range nums1 {
			exists[val]++
		}
		k := 0
		for _, val := range nums2 {
			if _, ok := exists[val]; ok && exists[val] > 0 {
				nums1[k] = val
				k++
				exists[val]--
			}
		}
		return nums1[:k]
	} else {
		for _, val := range nums2 {
			exists[val]++
		}
		k := 0
		for _, val := range nums1 {
			if _, ok := exists[val]; ok && exists[val] > 0 {
				nums2[k] = val
				k++
				exists[val]--
			}
		}
		return nums2[:k]
	}
}
