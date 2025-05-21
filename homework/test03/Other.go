package main

import "container/list"

/*
括号匹配
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	// 定义映射关系
	pair := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := list.New()

	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := pair[c]; ok {
			stack.PushBack(c)
		} else if stack.Len() > 0 && pair[stack.Back().Value.(byte)] == c {
			stack.Remove(stack.Back())
		} else {
			return false
		}
	}
	return stack.Len() == 0
}

/**
优先级括号匹配
给定一个只包括 '('，')'，'['，']'， '{'，'}' 的字符串 s ，括号的优先级排序由高到低依次为()，[]，{}，判断字符串是否有效。
有效字符串需满足：​
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
内层括号的优先级不低于外层括号。
*/
// priIsValid 判断带优先级的括号字符串是否有效
func priIsValid(s string, priorityOrder string) bool {
	if len(s)%2 != 0 {
		return false
	}

	// 括号配对映射
	pair := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	// 优先级映射（数值越小优先级越高）
	priorityMap := make(map[rune]int)
	for i, ch := range priorityOrder {
		priorityMap[ch] = i // 数值越小优先级越高
	}

	stack := list.New()

	for _, c := range s {
		if _, ok := pair[c]; ok {
			// 是左括号
			if stack.Len() > 0 {
				top := stack.Back().Value.(rune)
				// 检查当前括号优先级是否低于栈顶括号
				if priorityMap[c] < priorityMap[top] {
					return false
				}
			}
			stack.PushBack(c)
		} else {
			if stack.Len() == 0 {
				return false
			}
			top := stack.Back().Value.(rune)
			if pair[top] != c {
				return false
			}
			stack.Remove(stack.Back())
		}
	}

	return stack.Len() == 0
}

/**
下一个最大元素
给定一个非负整数组，对于每个元素，找出它右边第一个比它大的元素。若没有最大的元素返回 -1.
*/
// nextGreaterValue 返回每个元素右边第一个比它大的元素值，若无则为 -1
func nextGreaterValue(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
	}
	stack := list.New()

	for i := 0; i < n; i++ {
		// 当栈不为空且当前元素大于栈顶索引对应的值时
		for stack.Len() > 0 && nums[stack.Back().Value.(int)] < nums[i] {
			idx := stack.Remove(stack.Back()).(int)
			result[idx] = nums[i]
		}
		stack.PushBack(i)
	}

	return result
}
