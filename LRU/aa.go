package lru

/*
type Node struct {
	Value    int
	LastNode *Node
	NextNode *Node
}
type Nodes struct {
	HeadNode *Node
	TailNode *Node
}
type LRU struct {
	Cap      int
	MapData  map[int]*Node
	ListData Nodes
}

func (l *LRU) Get(k int) int {
	if node, ok := l.MapData[k]; ok {
		if node.LastNode != nil {
			node.LastNode.NextNode = node.NextNode
		}
		node.NextNode = l.ListData.HeadNode
		l.ListData.HeadNode = node
		return node.Value
	}
	return -1
}

func (l *LRU) Put(k, v int) {
	node := new(Node)
	node.Value = v
	if len(l.MapData) >= l.Cap {
		l.ListData.TailNode.LastNode.NextNode = nil
	}
	node.NextNode = l.ListData.HeadNode
	l.ListData.HeadNode = node
	l.MapData[k] = node
	return

} */
