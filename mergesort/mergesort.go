package sort

// Mergesort sorts an int slice using the mergesort algorithm.
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
	sorted := make([]int, n)
	// i and j are indices for left and right, respectively.
	var i, j int
	// While there are still elements left to compare...
	for i < len(left) && j < len(right) {
		// If the current element of 'left' is less or equal than the current element of 'right',
		if left[i] <= right[j] {
			// it goes into 'sorted' first.
			sorted[i+j] = left[i]
			i++
		} else {
			// Else, the element from 'right' goes first.
			sorted[i+j] = right[j]
			j++
		}
	}
	// There might still be unchecked elements in 'left' or 'right', so add them to sorted.
	for i < len(left) {
		sorted[i+j] = left[i]
		i++
	}
	for j < len(right) {
		sorted[i+j] = right[j]
		j++
	}
	return sorted
}
