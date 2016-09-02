package sort

import "math/rand"

// Quicksort sorts an int slice (in-place) using the quicksort algorithm.
func Quicksort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	i, j := 0, n-1
	// Pick a random pivot to ensure O(n*log(n)) time.
	pivot := nums[rand.Intn(n)]
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
	Quicksort(nums[:j+1])
	Quicksort(nums[j+1 : n])
}
