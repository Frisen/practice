package lru_test

import (
	lru "GinLearn/LRU"
	"testing"
)

func TestPut(t *testing.T) {
	l := new(lru.LRU)
	l.Init(3)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Put(3, 3)
	l.Put(4, 4)
	t.Log(l.List.Tail.Pre.V, l.List.Tail.Pre.Pre.V, l.List.Tail.Pre.Pre.Pre.V)
	t.Log(l.List.Head.Next.V, l.List.Head.Next.Next.V, l.List.Head.Next.Next.Next.V)
	r := l.Get(3)
	t.Log(r)
	t.Log(l.List.Head.Next.V, l.List.Head.Next.Next.V, l.List.Head.Next.Next.Next.V)
	l.Put(5, 5)
	t.Log(l.List.Head.Next.V, l.List.Head.Next.Next.V, l.List.Head.Next.Next.Next.V)
}
