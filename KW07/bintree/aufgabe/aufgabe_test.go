package aufgabe

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	node := NewNode()
	if node == nil {
		t.Fatalf("NewNode() returned nil")
	}
	if node.Left != nil || node.Right != nil {
		t.Errorf("New node should have nil Left and Right children")
	}
}

func TestIsEmpty(t *testing.T) {
	node := NewNode()
	if !node.IsEmpty() {
		t.Errorf("New node should be empty")
	}
	node.AddValue(10)
	if node.IsEmpty() {
		t.Errorf("Node should not be empty after adding a value")
	}
}

func TestAddValue(t *testing.T) {
	node := NewNode()
	node.AddValue(10)
	if node.Value != 10 {
		t.Errorf("Expected root value to be 10, got %d", node.Value)
	}
	if node.Left == nil || node.Right == nil {
		t.Errorf("Left and Right children should not be nil after adding a value")
	}
}

func TestAddMultipleValues(t *testing.T) {
	node := NewNode()
	values := []int{10, 5, 15, 3, 7, 13, 17}
	for _, v := range values {
		node.AddValue(v)
	}

	if node.Left.Value != 5 || node.Right.Value != 15 {
		t.Errorf("Expected left child value 5 and right child value 15, got %d and %d", node.Left.Value, node.Right.Value)
	}
	if node.Left.Left.Value != 3 || node.Left.Right.Value != 7 {
		t.Errorf("Left subtree values incorrect")
	}
	if node.Right.Left.Value != 13 || node.Right.Right.Value != 17 {
		t.Errorf("Right subtree values incorrect")
	}
}

func TestSumAll(t *testing.T) {
	node := NewNode()
	values := []int{10, 5, 15, 3, 7, 13, 17}
	for _, v := range values {
		node.AddValue(v)
	}
	sum := node.SumAll()
	if node.SumAll() != 70 {
		t.Errorf("Expected Sum = 70, got %v", sum)
	}

}
