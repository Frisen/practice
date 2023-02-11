package lru

type Node struct {
	K    int
	V    int
	Pre  *Node
	Next *Node
}
type NodeList struct {
	Head *Node
	Tail *Node
}
type LRU struct {
	Cap  int
	Map  map[int]*Node
	List NodeList
}

func (l *LRU) Init(cap int) {
	l.Cap = cap
	l.Map = make(map[int]*Node)
	head := new(Node)
	tail := new(Node)
	head.Next = tail
	tail.Pre = head
	list := NodeList{Head: head, Tail: tail}
	l.List = list
}

func (l *LRU) Get(k int) (v int) {
	//node:=new(Node) 先声明 再赋值跟下行的区别？
	node, ok := l.Map[k]
	if ok {
		v = node.V
	} else {
		v = -1
		return
	}
	//删除该节点
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	delete(l.Map, node.K)
	//将该节点移动到头部
	node.Next = l.List.Head.Next
	l.List.Head.Next.Pre = node
	l.List.Head.Next = node
	return
}

func (l *LRU) Put(k, v int) {
	node := &Node{K: k, V: v}
	//如果大小等于设置的容量，删除尾部节点
	if len(l.Map) == l.Cap {
		delete(l.Map, l.List.Tail.Pre.K)
		l.List.Tail.Pre.Pre.Next = l.List.Tail
		l.List.Tail.Pre = l.List.Tail.Pre.Pre
	}
	node.Next = l.List.Head.Next
	l.List.Head.Next.Pre = node
	l.List.Head.Next = node
	node.Pre = l.List.Head
	l.Map[k] = node
}
