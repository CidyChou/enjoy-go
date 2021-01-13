import "math"

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
func reverse(x int) int {
	var i int

	if x > int(math.Pow(2, 31)-1) || x < -int((math.Pow(2, 31))) {
		return 0
	}

	for x != 0 {
		i = i*10 + x%10
		x = x / 10
	}

	if i > int(math.Pow(2, 31)-1) || i < -int((math.Pow(2, 31))) {
		return 0
	}
	
	return i

}