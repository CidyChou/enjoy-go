/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	res := []int{0, 1, 2}

	for i := 3; i <= n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res[n]
}
