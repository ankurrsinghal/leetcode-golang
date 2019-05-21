package add_two_numbers

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var newList *ListNode
	var cursor *ListNode

	head1 := l1
	head2 := l2
	carry := 0

	for head1 != nil || head2 != nil || carry != 0 {
		var n1 int
		var n2 int

		if head1 != nil {
			n1 = head1.Val
			head1 = head1.Next
		} else {
			n1 = 0
		}

		if head2 != nil {
			n2 = head2.Val
			head2 = head2.Next
		} else {
			n2 = 0
		}

		sum := n1 + n2 + carry
		n := sum % 10
		carry = sum / 10

		node := &ListNode{n, nil}

		if newList == nil {
			newList = node
			cursor = newList
		} else {
			cursor.Next = node
			cursor = cursor.Next
		}
	}

	return newList
}
