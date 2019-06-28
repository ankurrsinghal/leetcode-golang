package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	i := l1
	j := l2
	var start, cur *ListNode
	for i != nil && j != nil {
		if i.Val >= j.Val {
			if start == nil {
				start = j
				cur = start
			} else {
				cur.Next = j
			}
			j = j.Next
		} else {
			if start == nil {
				start = i
				cur = start
			} else {
				cur.Next = i
			}
			i = i.Next
		}
		cur = cur.Next
	}

	if i != nil {
		for i != nil {
			cur.Next = i
			cur = cur.Next
			i = i.Next
		}
	}

	if j != nil {
		for j != nil {
			cur.Next = j
			cur = cur.Next
			j = j.Next
		}
	}

	return start
}
