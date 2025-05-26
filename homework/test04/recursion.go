package main

// 计算阶乘
func factorial(n int) int {
	// 基本情况
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 斐波那契数列
func fibonacci(n int) int {
	// 基本情况
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 斐波那契数列优化版 -- 记忆化搜索
func fibonacciOptimize(n int) int {
	memo := make(map[int]int)
	return fibonacciOptimizeHelper(memo, n)
}

// 辅助函数
func fibonacciOptimizeHelper(memo map[int]int, n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	result := fibonacciOptimizeHelper(memo, n-1) + fibonacciOptimizeHelper(memo, n-2)
	memo[n] = result
	return result
}
