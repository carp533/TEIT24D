package main

import (
	"testing"
)

// Hilfsfunktion zum Erstellen einer LinkedList aus einer Slice
func createLinkedList(values []int) *LinkedList {
	ll := &LinkedList{}
	if len(values) == 0 {
		return ll
	}
	ll.head = &Node{data: values[0]}
	current := ll.head
	for _, v := range values[1:] {
		current.next = &Node{data: v}
		current = current.next
	}
	return ll
}

// Hilfsfunktion zum Umwandeln einer LinkedList in eine Slice (zum Vergleich)
func (ll *LinkedList) toSlice() []int {
	var result []int
	current := ll.head
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}

// Testfälle für RemoveValues
func TestRemoveValue(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		remove   int
		expected []int
	}{
		{"Leere Liste", []int{}, 5, []int{}},
		{"Ein einziges Element entfernen", []int{5}, 5, []int{}},
		{"Kopf entfernen", []int{5, 1, 2, 3}, 5, []int{1, 2, 3}},
		{"Mittleres Element entfernen", []int{1, 2, 5, 3}, 5, []int{1, 2, 3}},
		{"Letztes Element entfernen", []int{1, 2, 3, 5}, 5, []int{1, 2, 3}},
		{"Mehrere gleiche Elemente entfernen", []int{5, 1, 5, 2, 5, 3, 5}, 5, []int{1, 2, 3}},
		{"Nicht vorhandenes Element entfernen", []int{1, 2, 3, 4}, 5, []int{1, 2, 3, 4}},
		{"Alle Elemente entfernen", []int{5, 5, 5, 5}, 5, []int{}},
		{"Nur ein Element behalten", []int{5, 5, 1, 5, 5}, 5, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := createLinkedList(tt.input)
			ll.RemoveValue(tt.remove)
			result := ll.toSlice()
			if !equal(result, tt.expected) {
				t.Errorf("Fehler: erwartetes Ergebnis %v, erhalten %v", tt.expected, result)
			}
		})
	}
}

// Vergleichsfunktion für Slices
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
