/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */
func checkPossibility(nums []int) bool {
	var index int
	var position []int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			index++
			position = append(position, i)
		}
	}

	if index >= 2 {
		return false
	} else if index == 1 {
		p := position[0]
		if p != 0 && p != len(nums) && (nums[p+1] != nums[p] || nums[p-1] != nums[p]) {
			if nums[p+1]-nums[p-1] < -1 {
				return false
			}
		}
	}
	return true
}
