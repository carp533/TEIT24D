package main

import (
	"fmt"
)

/* Aufgabe 3a:
implementiere eine Queue (Warteschlange)
wir benötigen die grundlegenden Operationen Enqueue, Dequeue, IsEmpty und Size.
Die Enqueue-Methode fügt ein Element am Ende der Queue hinzu,
Dequeue entfernt und gibt das erste Element aus der Queue zurück,
IsEmpty prüft, ob die Queue leer ist, und
Size gibt die Anzahl der Elemente in der Queue zurück.
verwende eine int-Liste zum verwalten der Elemente der Queue
*/
/*Aufgabe 3b:
baue die Queue um, sodass beliebige Elemente verwendet werden können.
welchen Datentyp musst du hierfür benutzen?
*/
type Queue struct {
}

func (q *Queue) Enqueue(item int) {
}

func (q *Queue) Dequeue() (int, error) {
	return 0, nil
}

func (q *Queue) IsEmpty() bool {
	return true
}

func (q *Queue) Size() int {
	return 42
}

func ExampleQueue() {
	//Output:
	// Dequeued item: 10
	// Dequeued item: 20
	// Queue size: 1
	// Is queue empty? false
	// queue is empty
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	v, err := queue.Dequeue()
	fmt.Println("Dequeued item:", v)
	v, err = queue.Dequeue()
	fmt.Println("Dequeued item:", v)
	fmt.Println("Queue size:", queue.Size())
	fmt.Println("Is queue empty?", queue.IsEmpty())
	v, err = queue.Dequeue()
	v, err = queue.Dequeue()
	fmt.Println(err)
}
