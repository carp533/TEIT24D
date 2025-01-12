Kommt in KW3 ... 
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
	for {
		var input int
		fmt.Scanln(&input)

		//mit Fehlerbehabdlung:
		// _, err := fmt.Scanln(&input)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		if input == 0 {
			break
		}
		result = append(result, input)
	}

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
	for _, v := range readNumbers() {
		sum += v
	}
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
	//return 0
	var count int
	for _, v := range readNumbers() {
		if v == 5 {
			count += 1
		}
	}
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
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}
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
	// return true
	return len(divisors(n)) == 2
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
	// return 0, 0
	if len(list) == 0 {
		return 0, -1
	}
	val, pos = list[0], 0
	for index, value := range list[1:] {
		if value > val {
			val = value
			pos = index + 1
		}
	}
	return val, pos
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
