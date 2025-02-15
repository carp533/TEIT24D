```go
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
	if n.IsEmpty() {
		n.Left = NewNode()
		n.Right = NewNode()
		n.Value = value
		return
	}
	if value < n.Value {
		n.Left.AddValue(value)
		return
	}
	n.Right.AddValue(value)
}

func (n *Node) SumAll() int {
	if n.IsEmpty() {
		return 0
	}
	return n.Value + n.Left.SumAll() + n.Right.SumAll()
}
```