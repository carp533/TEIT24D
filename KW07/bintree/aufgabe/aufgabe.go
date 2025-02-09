package aufgabe

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) IsEmpty() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node) AddValue(value int) {
}

func (n Node) SumAll() int {
	return 0
}
