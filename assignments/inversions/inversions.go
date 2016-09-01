package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile := flag.String("input", "./input.txt", "The path to the input file.")
	nums, err := readInputArray(*inputFile)
	if err != nil {
		panic(err)
	}
	_, inversions := sortAndCountInversions(nums)
	fmt.Println("Number of inversions:", inversions)
}

// readInputArray scans the input file, which should have an int per line, and returns an []int.
func readInputArray(path string) ([]int, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nums, err
		}
		nums = append(nums, num)
	}
	return nums, scanner.Err()
}

// sortAndCountInversions is basically a mergesort used to count inversions along the way.
// For a detailed explanation of mergesort, check /mergesort/mergesort.go
func sortAndCountInversions(nums []int) ([]int, int) {
	n := len(nums)
	if n <= 1 {
		return nums, 0
	}
	sortedLeft, leftInversions := sortAndCountInversions(nums[:n/2])
	sortedRight, rightInversions := sortAndCountInversions(nums[n/2:])

	totalInversions := leftInversions + rightInversions

	sorted := make([]int, n)
	var i, j int
	for i < len(sortedLeft) && j < len(sortedRight) {
		if sortedLeft[i] <= sortedRight[j] {
			sorted[i+j] = sortedLeft[i]
			i++
		} else {
			// If we copy an element from 'sortedRight' over to 'sorted', it means we have found an
			// inversion, and that there are inversions for all the remaining elements interface
			// 'sortedLeft' (len(sortedLeft) - i).
			sorted[i+j] = sortedRight[j]
			totalInversions += len(sortedLeft) - i
			j++
		}
	}
	for i < len(sortedLeft) {
		sorted[i+j] = sortedLeft[i]
		i++
	}
	for j < len(sortedRight) {
		sorted[i+j] = sortedRight[j]
		j++
	}

	return sorted, totalInversions
}
