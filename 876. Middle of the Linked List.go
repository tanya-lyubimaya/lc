package lc

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNodeNaive(head *ListNode) *ListNode {
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

func middleNodeClassic(head *ListNode) *ListNode {
	nodes := []*ListNode{head}

	for nodes[len(nodes)-1].Next != nil {
		nodes = append(nodes, nodes[len(nodes)-1].Next)
	}

	if len(nodes)%2 == 0 {
		return nodes[len(nodes)/2]
	}
	return nodes[(len(nodes)-1)/2]
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == nil {
			return slow
		}
	}

	return slow
}
