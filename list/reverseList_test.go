package list_test

import (
	"GinLearn/list"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := new(list.ListNode)
	head.Val = 1
	head.Next = &list.ListNode{Val: 2, Next: nil}
	t.Log(head.Val, head.Next.Val)
	r := list.ReverseList(head)
	t.Log(r.Val, r.Next.Val)
}
