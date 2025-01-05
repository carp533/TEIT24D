package main

import (
	"fmt"
)

/* Aufgabe 1:
 *
 * Schreiben Sie eine Funktion, die den Benutzer immer wieder nach einer Zahl fragt,
 * bis er eine 0 eingibt. Anschließend soll eine Liste zurückgegeben werden, die die
 * eingegebenen Zahlen enthält.
 * Testen Sie die Funktion in der main Funktion
 *
 */
func readNumbers() []int {
	result := make([]int, 0)
	// TODO
	return result
}

/* Aufgabe 2:
 *
 * Schreiben Sie die Funktion `readSum()`. Die Funktion soll den Benutzer immer wieder nach einer Zahl fragen,
 * bis er eine 0 eingibt. Anschließend soll die Summe aller Zahlen zurückgegeben werden .
 * Verwenden Sie dabei die obige Funktion `readNumbers()`.
 * Testen Sie die Funktion in der main Funktion
 */
func readSum() int {
	// TODO
	var sum int
	return sum
}

/* Aufgabe 3:
 *
 * Schreiben Sie die Funktion `readFives()`. Die Funktion soll die Anzahl der eingegebenen 5en zurück geben.
 * Verwenden Sie dabei die obige Funktion `readNumbers()`.
 * Testen Sie die Funktion in der main Funktion
 */
func readFives() int {
	// TODO
	var count int
	return count
}

/* Aufgabe 4:
 *
 * Schreiben Sie eine Funktion `divisors()`.
 * Die Funktion soll eine Zahl `n` erwarten und eine Liste liefern, die alle Teiler
 * von `n` enthält.
 *
 * Beispiele:
 * - divisors(10) == {1,2,5,10}
 * - divisors(60) == {1,2,3,4,5,6,10,12,15,20,30,60}
 * - divisors(37) == {1,37}
 * Testen Sie die Funktion mit den gegebenen Testfunktionen
 */
func divisors(n int) []int {
	result := make([]int, 0)
	return result
}

/* Aufgabe 4:
 *
 * Schreiben Sie (erneut) eine Funktion `isPrime()`.
 * Die Funktion soll eine Zahl `n` erwarten und true zurückliefern,
 * wenn n eine Primzahl ist.
 *
 * Nutzen Sie die Funktion `divisors()`
 * Testen Sie die Funktion mit den gegebenen Testfunktionen
 */
func isPrime(n int) bool {
	// TODO
	return true
}

/* Aufgabe 5:
 *
 * Schreiben Sie eine Funktion `listMax()`, die eine Liste erwartet und die
 * sowohl Position als auch Wert des größten Elements bestimmt.
 *
 * Hinweis: Aufruf in main() z.B. mit `x,y := listMax(...)`
 * Testen Sie die Funktion mit den gegebenen Testfunktionen
 */
func listMax(list []int) (val int, pos int) {
	// TODO
	return 0, 0

}

/*Hilfsfunktion für Ausgabe*/
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {

	// Testen Sie hier Ihre Funktionen.
	fmt.Println("Test readNumbers: bitte Zahlen eingeben:")
	li := readNumbers()
	fmt.Printf("Sie haben %v eingegeben.\n", li)

	fmt.Println("Test readFives: bitte Zahlen eingeben:")
	fives := readFives()
	fmt.Printf("Anzahl der eingegebenen 5en: %v.\n", fives)

	fmt.Println("Test readSum: bitte Zahlen eingeben:")
	sum := readSum()
	fmt.Printf("Summe der eingegeben Zahlen: %v.\n", sum)

}
