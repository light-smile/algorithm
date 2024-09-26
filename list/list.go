package list

type Node struct {
	Key   int
	Value string
	Next  *Node
}

type List struct {
	Head   *Node
	Length int
}

func NewList() *List {
	return &List{
		Head:   &Node{},
		Length: 0,
	}
}
func (l *List) Get(key int) *Node {
	node := l.Head
	if node.Next == nil {
		return nil
	}
	for node != nil {
		if node.Key == key {
			return node
		}
		node = node.Next
	}
	return nil

}
func (l *List) Append(key int, value string) {
	node := &Node{
		key, value, nil,
	}
	cur := l.Head

	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	l.Length++
}
func (l *List) Delete(key int) {
	node := l.Head
	var pre *Node
	for node != nil {
		if node.Key == key {
			pre.Next = node.Next
			l.Length--
		}
		pre = node
		node = node.Next
	}
}
