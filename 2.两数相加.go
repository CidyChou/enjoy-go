/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var array []int
	bol := true
	i := 0

	for !(l1 == nil && l2 == nil) || i == 1 {
		v1 := 0
		v2 := 0

		if l1 != nil {
			v1 = l1.Val
		}

		if l2 != nil {
			v2 = l2.Val
		}

		s := v1 + v2 + i
		if i == 1 {
			i = 0
		}

		if s > 9 {
			s -= 10
			i = 1
		}

		array = append(array, s)

		if i == 1 && bol == false {
			l1.Val = 0
			l2.Val = 0
		} else {
			if l1 != nil {
				l1 = l1.Next
			}

			if l2 != nil {
				l2 = l2.Next
			}
		}
	}

	var sumLn *ListNode = &ListNode{}
	var lastNode *ListNode = sumLn

	for _, a := range array {
		lastNode.Next = &ListNode{Val: a, Next: nil}
		lastNode = lastNode.Next
	}

	return sumLn.Next
}
