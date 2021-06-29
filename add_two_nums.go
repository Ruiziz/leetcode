package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = new(ListNode)
	var head = result
	var total int
	count := 0
	for {
		// case 1: both empty, stop
		if l1 == nil && l2 == nil && count == 0{
			break
		} else if l1.Next == nil {
			// case 2.2: either empty
			total = l2.Val + count
			result.Val = total % 10
			count = total / 10
			l2 = l2.Next
		} else if l2.Next == nil {
			// case 2.2
			total = l1.Val + count
			result.Val = total % 10
			count = total / 10
			l1 = l1.Next
		} else {
			// case 3: both not empty
			total = l1.Val + l2.Val + count
			result.Val = total % 10
			count = total / 10
			l1 = l1.Next
			l2 = l2.Next
		}
		result.Next = &ListNode{0, nil}
		result = result.Next
	}
	return head
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	var result *ListNode = dummy
	//result := new(ListNode)
	//var dummy *ListNode = result
	val := 0
	for {
		if (l1 == nil && l2 == nil && val == 0) {break}
		if (l1 != nil) {
			val += l1.Val
			l1 = l1.Next
		}
		if (l2 != nil) {
			val += l2.Val
			l2 = l2.Next
		}
		result = &ListNode{val % 10, new(ListNode)}
		val = val / 10
		result = result.Next
	}
	result = dummy
	for{
		if (result == nil) {break}
		fmt.Printf(string(result.Val))
		result = result.Next
	}
	return dummy
}
func main() {
	fmt.Printf("hello world")
	var l1 = &ListNode{2, &ListNode {4, &ListNode{3, nil}}}
	var l2 = &ListNode{5, &ListNode {6, &ListNode{4, nil}}}
	addTwoNumbers2(l1,l2)
	//fmt.Printf(string(addTwoNumbers2(l1, l2).Val))
}
