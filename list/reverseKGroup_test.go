package list_test

import (
	"GinLearn/list"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	l1 := &list.ListNode{Val: 1}
	l2 := &list.ListNode{Val: 2}
	l3 := &list.ListNode{Val: 3}
	l4 := &list.ListNode{Val: 4}
	l5 := &list.ListNode{Val: 5}
	l6 := &list.ListNode{Val: 6}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	r := list.ReverseKGroup(l1, 2)
	t.Log(r.Val, r.Next.Val, r.Next.Next.Val, r.Next.Next.Next.Val, r.Next.Next.Next.Next.Val, r.Next.Next.Next.Next.Next.Val)
}
