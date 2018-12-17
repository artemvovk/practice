package main

import (
	"github.com/kierachell/practice/data"
)

func main() {
}

func twoSum(arr []int, target int) []int {
	indicies := make([]int, 2)
	complements := make(map[int]int)
	for idx, val := range arr {
		if compl, ok := complements[val]; ok {
			indicies[0] = idx
			indicies[1] = compl
			return indicies
		}
		complements[target-val] = idx
	}
	return indicies
}

func addTwoNumbers(l1 *data.ListNode, l2 *data.ListNode) *data.ListNode {
	carry := 0
	sumNode := data.ListNode{
		Val: 0,
	}
	next := &sumNode
	for {
		next.Val += carry
		carry = 0
		if l2 != nil {
			next.Val += l2.Val
			l2 = l2.Next
		}
		if l1 != nil {
			next.Val += l1.Val
			l1 = l1.Next
		}
		if next.Val >= 10 {
			carry = 1
			next.Val = next.Val % 10
		}
		if l1 != nil || l2 != nil {
			next.Next = &data.ListNode{}
			next = next.Next
		} else {
			break
		}
	}
	if carry > 0 {
		next.Next = &data.ListNode{
			Val: carry,
		}
	}
	return &sumNode
}
