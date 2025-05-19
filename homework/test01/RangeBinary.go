package main

/*
给你一个非负整数 x ，计算并返回 x 的算术平方根 。由于返回类型是整数，结果只保留 整数部分 ，小数部舍去。
*/
func sqrtRoot(x int) int {
	left := 1
	right := x
	result := 1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
