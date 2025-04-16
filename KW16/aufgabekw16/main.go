package main

/**********************************************************************
* Aufgabe A (10 Punkte):
* AUFGABENSTELLUNG:


Definiere ein Interface in Go, das Methoden enthält, um die linken und
rechten Kinder eines Knotens in einem Heap zurückzugeben. Implementiere
dieses Interface für einen Max-Heap, der als Array dargestellt wird.

Beschreibung:
Erstelle ein Interface HeapNode mit den Methoden Left() und Right().
Implementiere dieses Interface für eine Struktur Heap so, dass die Methoden
die Indizes der linken und rechten Kinder eines Knotens im Heap zurückgeben.
Die Kinder-Indizes sollten für den Fall, dass ein Kind existiert,
zurückgegeben werden; andernfalls sollt -1 (ein ungültiger Index) zurückgegeben
werden.
Beispiel:
Input:

Max-Heap: [42, 29, 18, 14, 7, 18, 12, 11, 5, 13]
Knoten-Index: 1 (Element: 29)
Output:

Linkes Kind: 14
Rechtes Kind: 7

**********************************************************************/

// Heap Struktur, die das HeapNode Interface implementiert
type Heap struct {
	elements []int
}

// NewHeap initialisiert einen neuen Heap
func NewHeap(elements []int) *Heap {
	return &Heap{elements: elements}
}

func main() {
	// heap := NewHeap([]int{42, 29, 18, 14, 7, 18, 12, 11, 5, 13})
	// index := 1

	// leftChild := heap.Left(index)
	// rightChild := heap.Right(index)

	// fmt.Printf("Knoten %d (Wert: %d)\n", index, heap.elements[index])
	// fmt.Printf("Linkes Kind: %d\n", leftChild)
	// fmt.Printf("Rechtes Kind: %d\n", rightChild)
}
