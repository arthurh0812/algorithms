package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(nums ...int) *ListNode {
	for i, n := range nums {
		
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil { // base case
		return nil
	}
	digit := 0
	next := l1
	if l1 == nil {
		digit = l2.Val
		next = l2
	} else if l2 == nil {
		digit = l1.Val
		next = l1
	} else {
		digit = l1.Val + l2.Val
	}

	if digit >= 10 {
		digit = digit - 10
		if next.Next != nil {
			next.Next.Val++
		} else {
			next.Next = &ListNode{Val: 1}
		}
	}

	if l1 == nil {
		return &ListNode{Val: digit, Next: addTwoNumbers(nil, l2.Next)}
	} else if l2 == nil {
		return &ListNode{Val: digit, Next: addTwoNumbers(l1.Next, nil)}
	}
	return &ListNode{Val: digit, Next: addTwoNumbers(l1.Next, l2.Next)}
}

func main() {

}