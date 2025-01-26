package main

import (
	"fmt"
	"testing"
)

/* Aufgabe 4a:
implementiere das Sieb des Eratosthenes  (Primzahl Sieb) -> siehe Wikipedia
1. definiere einen booleschen array A der Länge N+1 und initialisiere ab index 2 alles mit true
   setze A[0] = A[1] = false
2. schleife mit i:=2, i<wurzel(N)
	2.1 wenn A[i], dann setze für alle Vielfachen J von i A[J] auf false
3. erstelle die Ergebnisliste aus allen i mit A[i]=true
Wie kann man i<Wurzel(N) umformulieren, sodass keine Wurzel benötigt wird?
Aufgabe 4b:
schaue dir die beiden benchmark-Funktionen an und führe diese aus.
Aufgabe 4c (optional):
schreibe einen "naiven" Primezahl Algorithmus, indem du dir (beginnend mit 2) eine Liste mit
Primzahlen erstellt und in einer Schleife für jede neue Zahl prüfst, ob diese ein Vielfaches
einer der bereits gefundenen Primzahlen ist (%); falls nein, füge die Zahl zur Liste der Primzahlen
hinzu. Schreibe für diese Funktion einen Benchmark und vergleiche das Ergebnis aus 4c.

*/

func PrimeNumbers(N int) []int {
	result := make([]int, 0)
	return result
}

func ExamplePrimeNumbers() {
	//Output:
	//[2 3 5 7 11 13 17 19 23 29]
	primes := PrimeNumbers(30)
	fmt.Println(primes)
}

var num = 1000

func BenchmarkPrimeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeNumbers(num)
	}
}

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1_000},
	{input: 10_000},
	{input: 100_000},
}

func BenchmarkPrimeNumbers_2(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				PrimeNumbers(v.input)
			}
		})
	}
}
