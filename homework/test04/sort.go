package main

// 快速排序
func quickSort(array []int) []int {
	if array == nil || len(array) == 0 {
		return array
	}
	quickSortHelper(array, 0, len(array)-1)
	return array

}

// 快速排序递归辅助函数
func quickSortHelper(array []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(array, left, right)
	quickSortHelper(array, left, pivot-1)
	quickSortHelper(array, pivot+1, right)
}

// 快排分区
func partition(array []int, left, right int) int {
	i := left
	j := right
	pivot := array[left]
	for i <= j {
		if array[i] <= pivot {
			i += 1
		} else {
			swap(array, i, j)
			j--
		}
	}
	swap(array, i, right)
	return i
}

// 交换数组元素
func swap(array []int, i, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}

// 归并排序
func mergeSort(array []int) []int {
	if array == nil || len(array) == 0 {
		return array
	}
	helper := make([]int, len(array))
	mergeSortFunc(array, helper, 0, len(array)-1)
	return array
}

func mergeSortFunc(array []int, helper []int, left int, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSortFunc(array, helper, left, mid)
	mergeSortFunc(array, helper, mid+1, right)
	merge(array, helper, left, mid, right)
}

func merge(array []int, helper []int, left int, mid int, right int) {
	for i := left; i <= right; i++ {
		helper[i] = array[i]
	}
	i := left
	j := mid + 1
	index := left

	for i <= mid && j <= right {
		if helper[i] <= helper[j] {
			array[index] = helper[i]
			i++
		} else {
			array[index] = helper[j]
			j++
		}
		index++
	}
	for i <= mid {
		array[index] = helper[i]
		i++
		index++
	}
	for j <= right {
		array[index] = helper[j]
		j++
		index++
	}
}
