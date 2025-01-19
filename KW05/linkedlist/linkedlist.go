package main

import (
	"fmt"
)

/* Eine einfach verkettete Liste ist eine lineare Datensammlung,
in der jedes Element (oder Knoten, engl. Node) einen Wert speichert
und auf das nächste Element in der Liste verweist.
Das letzte Element hat keine "nächstes" Element, der Zeiger ist <nil>
(nicht referenzierter Nullwert).
*/

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) AddAtEnd(data int) {
	newNode := &Node{data: data}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	ll.head.add(newNode)
}
func (n *Node) add(node *Node) {
	if n.next == nil {
		n.next = node
		return
	}
	n.next.add(node)

}

func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (ll *LinkedList) Size() int {
	return 0
}
func (node *Node) size() int {
	return 0
}

func main() {
	ll := LinkedList{}
	ll.AddAtEnd(1)
	ll.AddAtEnd(2)
	ll.AddAtEnd(3)
	fmt.Println("ll.Size()=", ll.Size())
	ll.AddAtEnd(3)
	ll.AddAtEnd(3)
	ll.AddAtEnd(3)
	fmt.Println("ll.Size()=", ll.Size())
	ll.Print() // Output: 1 2 3 3 3 3
}
