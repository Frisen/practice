package list

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	oldHead := head
	end := head
	for i := 0; i < k; i++ {
		if i < (k-1) && end.Next == nil {
			break
		}
		end = end.Next
	}
	newHead := reverse(oldHead, end)
	oldHead.Next = ReverseKGroup(end, k)
	return newHead
}

func reverse(head, end *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := new(ListNode)
	cur := head
	left := head
	for cur != end {
		left = cur.Next
		cur.Next = pre
		pre = cur
		cur = left
	}
	return pre
}
