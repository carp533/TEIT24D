package insertion

/*
Bei InsertionSort nehmen wir immer ein
beliebiges Element der noch nicht sortierten
Daten auf und sortieren es an der richtigen
Stelle ein.
Leider lässt sich in einem Array nicht einfach
ein neues Element zwischen zwei benachbarte
„einschieben“, wir müssen zunächst Platz
schaffen, indem wir die störenden Elemente alle
um eins nach rechts verschieben und dann das
Element an den freigewordenen Platz
einsetzen.
*/
func Sort(a []int) []int {
	for j := 1; j < len(a)-1; j++ {
		elem := a[j]
		i := j - 1
		for i > -1 && a[i] > elem {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = elem
	}
	return a
}
