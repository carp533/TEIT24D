package main

import "fmt"

/* Aufgabe 1:
 *
 * Schreiben Sie eine Funktion, die den Benutzer immer wieder nach einer Zahl fragt,
 * bis er eine 0 eingibt. Anschließend soll die Summe der eingegebenen Zahlen
 * auf die Konsole ausgegeben werden.
 *
 * Hinweis: In der main()-Funktion steht ein Beispiel, wie man eine einzelne Zahl
 *          einliest und wieder ausgibt.
 */
func readZahl_loesung() int {
	sum := 0
	for {
		fmt.Print("Bitte eine Zahl eingeben (0 für Ende): ")
		var input int
		fmt.Scanln(&input)
		sum += input
		if input == 0 {
			break
		}
	}
	return sum
}

/* Aufgabe 2:
 *
 * Schreiben Sie eine Funktion, die den Benutzer immer wieder nach einer Zahl fragt,
 * bis er eine 0 eingibt. Anschließend soll die Anzahl der eingegebenen Fünfen
 * auf die Konsole ausgegeben werden.
 */
func readFives_loesung() int {
	sum_fives := 0
	for {
		fmt.Print("Bitte eine Zahl eingeben (0 für Ende): ")
		var input int
		fmt.Scanln(&input)
		if input == 5 {
			sum_fives += 1
		}
		if input == 0 {
			break
		}
	}
	return sum_fives
}

/* Aufgabe 3:
 *
 * Verallgemeinern Sie die Funktion readFives() so, dass man die zu zählende Zahl
 * als Parameter übergeben kann.
 */
func readSpecZahl_loesung(zahl int) int {
	sum_zahl := 0
	for {
		fmt.Print("Bitte eine Zahl eingeben (0 für Ende): ")
		var input int
		fmt.Scanln(&input)
		if input == zahl {
			sum_zahl += 1
		}
		if input == 0 {
			break
		}
	}
	return sum_zahl
}
func main_loesung() {
	fmt.Println("Lösung Aufgabe 1 (Summe der eingegeben Zahlen):")
	summe := readZahl_loesung()
	fmt.Println("Die Summe der eingebenen Zahlen ist: ", summe)

	fmt.Println("Lösung Aufgabe 2 (Anzahl der eingegeben 5en):")
	summe5 := readFives_loesung()
	fmt.Printf("Sie haben %v mal die Zahl 5 eingegben.\n", summe5)

	fmt.Println("Lösung Aufgabe 3 (Anzahl einer eingegeben Zahl):")
	var input int
	fmt.Print("Bitte geben Sie die zu zählende Zahl ein: ")
	fmt.Scanln(&input)
	summezahl := readSpecZahl_loesung(input)
	fmt.Printf("Sie haben %v mal die Zahl %v eingegben.\n", summezahl, input)
}
