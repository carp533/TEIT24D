package main

import (
	"fmt"
)

/*Implementierung eines binären Suchbaums.
Je nachdem, in welcher Reihenfolge die Elemente in den Baum geschrieben werden,
ändert sich der Aufbau des Baums, z.B:
		X				Z
       / \ 			   /
      U   Z			  X
					 /
					U
für uns sind nur die Methoden insert und search relevant.
*/

/*Aufgabe 1: programmiere die Methoden
func (t *BinarySearchTree) search(value rune) bool
func (node *Node) search(key rune) bool
(Lösung ist im Kommentar am Ende).
Aufgabe 2: füge die Zeichen '1' bis '7' in einer Reihenfolge in den Baum, sodass er
a) maximal balanciert ist, d.h. die Höhe des Baums ist minimal.
b) "entartet" ist, d.h. die Höhe des Baums ist maximal.
Teste das in der main funktion (Tiefe des Baums = 3 bzw. 7) und zeichne die beiden Bäume.
*/

// Wurzel des Baums
type BinarySearchTree struct {
	root *Node
}

// Struktur für einen Knoten
type Node struct {
	data  rune
	left  *Node
	right *Node
}

func (t *BinarySearchTree) search(value rune) bool {
	if t.root != nil {
		return t.root.search(value)
	}
	return false
}

func (node *Node) search(key rune) bool {
	return false
}

func (t *BinarySearchTree) insert(value rune) {
	if t.root != nil {
		t.root.insert(value)
	} else {
		t.root = &Node{data: value, left: nil, right: nil}
	}
}

func (node *Node) insert(value rune) {
	if value < node.data {
		if node.left == nil {
			node.left = &Node{
				left:  nil,
				right: nil,
				data:  value,
			}
		} else {
			node.left.insert(value)
		}
	}
	if value > node.data {
		if node.right == nil {
			node.right = &Node{
				left:  nil,
				right: nil,
				data:  value,
			}
		} else {
			node.right.insert(value)
		}
	}
}

func inorderTraversal(root *Node) []rune {
	// Traversiere den linken Teilbaum Inorder
	// Besuche die Wurzel
	// Traversiere den rechtenTeilbaum Inorder
	if root == nil {
		return nil
	}
	output := make([]rune, 0)
	return output
}

func preorderTraversal(root *Node) []rune {
	// Besuche die Wurzel
	// Traversiere den linken Teilbaum Preorder
	// Traversiere den rechtenTeilbaum Preorder
	if root == nil {
		return nil
	}
	output := make([]rune, 0)
	return output
}

func postorderTraversal(root *Node) []rune {
	// Traversiere den linken Teilbaum Postorder
	// Traversiere den rechtenTeilbaum Postorder
	// Besuche die Wurzel
	if root == nil {
		return nil
	}
	output := make([]rune, 0)
	return output
}

func levelOrder(root *Node) [][]rune {
	if root == nil {
		return [][]rune{}
	}

	queue := []*Node{}
	queue = append(queue, root)
	result := [][]rune{}

	for len(queue) > 0 {
		sz := len(queue)
		level := []rune{}
		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, rune(node.data))

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		result = append(result, level)
	}
	return result
}

func main() {

	tree := &BinarySearchTree{}
	chars := []rune{'b', 'i', 'n', 'a', 'r', 'y', 't', 'r', 'e', 'e'}

	for _, v := range chars {
		tree.insert(v)
	}

	fmt.Printf("tree.search('x') = %v\n", tree.search('x'))
	fmt.Printf("tree.search('y') = %v\n", tree.search('y'))

	fmt.Printf("Preorder:\n")
	for _, c := range preorderTraversal(tree.root) {
		fmt.Printf("%c ", c)
	}
	fmt.Printf("\nInorder:\n")
	for _, c := range inorderTraversal(tree.root) {
		fmt.Printf("%c ", c)
	}
	fmt.Printf("\nPostorder:\n")
	for _, c := range postorderTraversal(tree.root) {
		fmt.Printf("%c ", c)
	}
	fmt.Printf("\nLevelorder:\n")
	for _, arr := range levelOrder(tree.root) {
		for _, c := range arr {
			fmt.Printf("%c ", c)
		}
		fmt.Printf("\n")
	}
	fmt.Println("Tiefe des Baums:", len(levelOrder(tree.root)))
}

/*














func (t *BinarySearchTree) search(value rune) bool {
	if t.root != nil {
		return t.root.search(value)
	}
	return false
}

func (node *Node) search(key rune) bool {
	if node == nil {
		return false
	} else if key == node.data {
		return true
	} else {
		if node.data > key {
			return node.left.search(key)
		} else {
			return node.right.search(key)
		}
	}
}
*/
