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
	if ll.head == nil {
		return 0
	}
	return ll.head.size()
}
func (node *Node) size() int {
	if node.next == nil {
		return 1
	}
	return 1 + node.next.size()
}

/*
Aufgaben:
(1) implementiere die Methode "Size" rekursiv

	func (ll LinkedList) Size() int
	Hilfdfunktion func (node Node) size() int

(2) implementiere die "IsInList" Methode

	wie sieht die Signatur dieser Methode aus?

(3) implementiere die Methode "RemoveValue"

	wie sieht die Signatur dieser Methode aus?
		a) lösche den ersten gefundenen Wert
		b) lösche allegefundene Werte
*/
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
//Lösung für 3b:
func (ll *LinkedList) RemoveValue(n int) {
	if ll.head == nil {
		return
	}
	if ll.head.data == n {
		ll.head = ll.head.next
		ll.RemoveValue(n)
		return
	}
	ll.head.removeValues(n)
}

func (node *Node) removeValues(n int) {
	if node.next == nil {
		return
	}
	if node.next.data == n {
		node.next = node.next.next
		node.removeValues(n)
		return
	}
	node.next.removeValues(n)

}