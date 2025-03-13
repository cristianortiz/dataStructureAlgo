package twoarrayintersection

//given two integer arrays, nums1, nums2 return another integer array with all the common elements
// between nums1 and nums2 without duplicates
// ej: entering: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]

func TwoArrayIntersection(nums1, nums2 []int) []int {
	interMap := make(map[int]bool) // Map to save nums1 elements
	var inter []int
	seen := make(map[int]bool) // Mapa to avoid duplicates in result array

	// save nums1 elements in a map
	for _, num1 := range nums1 {
		interMap[num1] = true
	}

	// looking in nums2 the elements from nums1 added to the map
	for _, num2 := range nums2 {
		if interMap[num2] && !seen[num2] {
			inter = append(inter, num2)
			seen[num2] = true
		}
	}

	return inter
}
