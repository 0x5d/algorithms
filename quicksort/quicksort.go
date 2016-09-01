package sort

// Quicksort sorts an int slice (in-place) using the quicksort algorithm.
func Quicksort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	i, j := 0, n-1
	pivot := nums[i]
	for {
		for nums[i] < pivot {
			i++
		}
		for nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	if j > 0 {
		Quicksort(nums[:j+1])
	}
	if j+1 < n-1 {
		Quicksort(nums[j+1 : n])
	}
}
