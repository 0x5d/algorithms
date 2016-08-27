package sort

// Mergesort sorts an int slice using the mergesort algorithm
func Mergesort(nums []int) []int {
	n := len(nums)
	// An empty slice or a slice with a single element is already sorted.
	if n <= 1 {
		return nums
	}
	// Split the input slice and call Mergesort on each half.
	left := Mergesort(nums[:n/2])
	right := Mergesort(nums[n/2:])
	// sorted will hold the sorted elements.
	var sorted []int
	// While there are still elements left to compare...
	for len(left) > 0 && len(right) > 0 {
		// If the first element of left is less or equal than the first element of right,
		if left[0] <= right[0] {
			// append it to sorted,
			sorted = append(sorted, left[0])
			// remove it from left (assign to left its "last" elements), as we've already checked it.
			left = rest(left)
		} else {
			// Else, append the first element of right and
			sorted = append(sorted, right[0])
			// remove it from right.
			right = rest(right)
		}
	}
	// There might still be unchecked elements in left or right, so append them to sorted.
	for len(left) > 0 {
		sorted = append(sorted, left[0])
		left = rest(left)
	}
	for len(right) > 0 {
		sorted = append(sorted, right[0])
		right = rest(right)
	}
	return sorted
}

func rest(slice []int) []int {
	if len(slice) <= 1 {
		return []int{}
	}
	return slice[1:]
}
