package main

import (
	"fmt"
)

func main() {
	// Einige Beispiele, wie man mit If-Then-Else Dinge überprüft.

	// Eine Variable definieren und ihr sofort einen Wert geben.
	// x ist eine Variable vom Typ int, weil 42 vom Typ int ist.
	x := 42
	// Alternative Deklaration einer Variablen:
	//var y int
	//y = 42

	// Die folgende Anweisung unterscheidet, ob x größer oder kleiner ist als 100
	// und es wird dann ein entsprechender String ausgegeben.
	if x > 100 {
		fmt.Println("Die Zahl ist groß.")
	} else {
		fmt.Println("Die Zahl ist winzig.")
	}

	// Die folgende Anweisung prüft, ob x durch 5 teilbar ist und gibt das Ergebnis aus.
	if x%5 == 0 {
		fmt.Println("Die Zahl ist durch 5 teilbar.")
	} else {
		fmt.Println("Die Zahl ist nicht durch 5 teilbar.")
	}
	// Anmerkung: Der Operator % heißt "modulo". Er liefert den Rest beim Dividieren.
	// Beispiele:
	//   5 % 3 == 2     "5 geteilt durch 3 lässt den Rest 2."
	//   25 % 4 == 1    "25 geteilt durch 4 lässt den Rest 1."
	//   Es gilt: x ist durch y teilbar, wenn x % y == 0, d.h. wenn beim Dividieren kein Rest bleibt.
}
