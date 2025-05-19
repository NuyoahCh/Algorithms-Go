package main

func findTargetInRotatedArray(nums []int, target int) int {
	// 特殊判断
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// 初始化搜索区间的左右边界
	left, right := 0, len(nums)-1

	// 当搜索区间有效时，即左边界不超过右边界时
	for left <= right {
		// 计算中间位置
		mid := left + (right - left)
		// 如果中间值就是目标值
		if nums[mid] == target {
			return mid
		}
		// 判断左半部分是否有序
		if nums[left] <= nums[mid] {
			// 目标值在左半部分范围
			if nums[left] <= target && target < nums[mid] {
				// 缩小范围到左半部分
				right = mid - 1
			} else {
				// 否则定位右半部分
				left = mid + 1
			}
			// 否则右半部分一定有效
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
