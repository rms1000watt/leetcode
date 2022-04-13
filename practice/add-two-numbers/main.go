// https://leetcode.com/problems/add-two-numbers/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (out *ListNode) {
	sum := 0
	carry := 0

	out = &ListNode{}
	first := out

	for {
		sum = l1.Val + l2.Val + carry

		carry = 0
		if sum > 9 {
			carry = 1
			out.Val = sum % 10
		} else {
			out.Val = sum
		}

		if l1.Next == nil && l2.Next == nil && carry == 0 {
			break
		}

		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1.Val = 0
		}

		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2.Val = 0
		}

		out.Next = &ListNode{}
		out = out.Next
	}

	out = first
	return
}
