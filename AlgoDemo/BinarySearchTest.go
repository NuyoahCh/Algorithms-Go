package main

import (
	"fmt"
)

type BinarySearchTest struct {
	nums   []int
	target int
	x      int
}

func binarySearch(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

func sqrtRoot(x int) int {
	left := 1
	right := x
	result := 1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			left = mid + 1
			result = mid
		} else {
			right = mid - 1
		}
	}
	return result
}

func findTargetInRotatedArray(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if nums[mid] > target {
				if nums[left] > target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target {
				if nums[right] < target {
					right = mid + 1
				} else {
					left = mid + 1
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Println("Hello Binary Search")

}
