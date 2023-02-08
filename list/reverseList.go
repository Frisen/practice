package list

/* Definition for singly-linked list. */
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := new(ListNode)
	cur := head
	left := head
	for cur != nil {
		left = cur.Next
		cur.Next = pre
		pre = cur
		cur = left
	}
	return pre
	// if head == nil {
	// 	return head
	// }
	// waitList := head
	// pre := new(ListNode)
	// flag := true
	// for flag {
	// 	if waitList.Next == nil {
	// 		flag = false
	// 	}
	// 	temp := waitList.Next
	// 	waitList.Next = pre
	// 	pre = waitList
	// 	waitList = temp
	// }
	// return pre
}
