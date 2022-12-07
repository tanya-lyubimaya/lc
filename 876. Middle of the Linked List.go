package lc

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	cnt := 1
	curr := head

	for curr.Next != nil {
		cnt++
		curr = curr.Next
	}

	curr = head

	if cnt%2 == 0 {
		for i := 0; i < cnt/2; i++ {
			curr = curr.Next
		}
	} else {
		for i := 0; i < (cnt-1)/2; i++ {
			curr = curr.Next
		}
	}

	return curr
}
