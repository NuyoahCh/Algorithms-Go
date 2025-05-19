package main

/*
基础二分查找算法模版
要去注意边界的处理，在脑中存在两个指针不断移动
*/
func binarySearch(nums []int, target int) int {
	// 左边界
	left := 0
	// 右边界
	right := len(nums) - 1
	// 当左边界小于等于右边界时
	for left <= right {
		// 取二分中间节点，折中查找
		mid := left + (right-left)/2
		// 相等直接返回中间 mid
		if nums[mid] == target {
			return mid
			// 折中小于目标值
		} else if nums[mid] < target {
			// 左边界右移
			left = mid + 1
			// 折中大于目标值
		} else if nums[mid] > target {
			// 右边界左移
			right = mid - 1
		}
	}
	return -1
}
