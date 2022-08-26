package maths

import (
	"sort"
)

// Number - type constraint for numeric types
type Number interface {
	int | int8 | int16 |int32 | int64 | float32 | float64
}

// Min - return the lowest 'n' numbers from the provided numbers
func Min[TN Number](nums []TN, qualifier int) []TN {
	// should always return at least one number
	if qualifier <= 0 {
		qualifier = 1
	}

	// if requested more numbers than there are, just return all the numbers
	if qualifier > len(nums) {
		return nums
	}

	// sort the numbers in ascending, and return the frist q elements
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	return nums[:qualifier]
}

// Max - return the highest 'n' numbers from the provided numbers 
func Max[TN Number](nums []TN, qualifier int) []TN {
	// should always return at least one number
	if qualifier <= 0 {
		qualifier = 1
	}

	// if requested more numbers than there are, just return all the numbers
	if qualifier > len(nums) {
		return nums
	}

	// sort the numbers in descending, and return the frist q elements
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	return nums[:qualifier]
}

// Avg - the mean of the the provided numbers
func Avg[TN Number](nums []TN) float64 {
	if len(nums) < 1 {
		return 0
	} else if len(nums) == 1 {
		return float64(nums[0])
	}

	var sum TN = 0
	for _, val := range nums {
		sum += val
	}
	return float64(sum)/float64(len(nums))
}

// Median - return the median of the provided numbers
func Median[TN Number](nums []TN) float64 {
	if len(nums) == 0 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	// get the middle index of the array
	m := len(nums)/2
	// if the array length is odd, this will be the value to return
	if len(nums) % 2 != 0 {
		return float64(nums[m])
	}
	// otherwise, get the average of the two middle values
	return Avg(nums[m-1:m+1])
}

// Percentile - return the qth percentile for the provided set of numbers
//  using the nearest-rank method, the returned value will always be one of 
//  the provided numbers
func Percentile[TN Number](nums []TN, q int) TN {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	// get the index of the last value less than or equal to the percenile
	p := ((q * len(nums)) / 100) - 1
	if p < 0 {
		p = 0
	}
	return nums[p]
}
