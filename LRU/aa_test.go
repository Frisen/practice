package lru

import (
	"testing"
)

func TestPut(t *testing.T) {
	l := new(LRU)
	l.Cap = 3
	l.MapData = make(map[int]*Node)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Put(3, 3)
	// l.Put(4, 4)
	for _, v := range l.MapData {
		t.Log(v.Value)
	}
}
