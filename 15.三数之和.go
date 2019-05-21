/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
func threeSum(nums []int) [][]int {
	var res = make([][]int, 0)
	index := 0
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-2 {
			if nums[i]+nums[i+1]+nums[i+2] == 0 {
				res[index] = []int{nums[i], nums[i+1], nums[i+2]}
				index++
			}
		}
	}
	return res
}
