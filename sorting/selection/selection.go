package selection

/*Bei SelectionSort suchen wir in jedem Schritt das kleinste(größte) der
noch ungeordneten Elemente und ordnen es am rechten Ende der
bereits sortierten Elemente ein. In einem Array mit Indexbereich
[Lo..Hi] sei k die erste Position und i die Position des kleinsten
Elementes im noch unsortierten Bereich.
Wiederholen wir diesen Vorgang so lange, bis k = Hi gilt, so ist am
Ende das ganze Array sortiert.*/

func Sort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
	return a
}

// LÖSUNG:

// func Sort(a []int) []int {
// 	for i := 0; i < len(a)-1; i++ {
// 		minIndex := i
// 		for j := i + 1; j < len(a); j++ {
// 			if a[j] < a[minIndex] {
// 				minIndex = j
// 			}
// 		}
// 		a[i], a[minIndex] = a[minIndex], a[i]
// 	}
// 	return a
// }
