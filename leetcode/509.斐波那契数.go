/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */
func fib(N int) int {

	if N > 1 {
		return fib(N-1) + fib(N-2)
	}

	if N == 1 {
		return 1
	}

	return 0
}
