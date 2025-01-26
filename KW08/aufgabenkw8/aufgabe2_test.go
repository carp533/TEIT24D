package main

import (
	"fmt"
	"sort"
	"strings"
)

/*Aufgabe 2a: analysiere einen String auf doppelte Zeichen
Hinweis: caste den String in eine Liste mit rune:
chars := []rune(s)
du kannst
return fmt.Errorf(" %q is duplicated at positions %d and %d.\n", chars[pos1], pos1+1, pos2+1)
zurück geben, falls ein Zeichen doppelt ist.
*/

func Analyze(s string) error {
	charPositions := make(map[rune]int) // Karte für Zeichen und ihre erste Position
	chars := []rune(s)
	for i, char := range chars {
		if pos, exists := charPositions[char]; exists {
			// Zeichen wurde bereits gefunden, gebe Fehler zurück
			return fmt.Errorf(" %q is duplicated at positions %d and %d.\n", char, pos+1, i+1)
		}
		charPositions[char] = i // Speichere die Position des Zeichens
	}

	return nil // Keine doppelten Zeichen gefunden
}

func ExampleAnalyze() {
	strings := []string{
		"",
		"abcABC",
		"XYZ 12X",
		"01234567890ABCDEFGHIJKLMN0PQRSTUVWXYZ0X",
		"😍😀🙌💃😍🙌🐬🐳🐋🐡",
	}
	for _, s := range strings {
		fmt.Printf("Analyze(%q): ", s)
		if err := Analyze(s); err != nil {
			fmt.Print(err)
		} else {
			fmt.Printf("ok\n")
		}

	}
	//Output:
	//Analyze(""): ok
	//Analyze("abcABC"): ok
	//Analyze("XYZ 12X"):  'X' is duplicated at positions 1 and 7.
	//Analyze("01234567890ABCDEFGHIJKLMN0PQRSTUVWXYZ0X"):  '0' is duplicated at positions 1 and 11.
	//Analyze("😍😀🙌💃😍🙌🐬🐳🐋🐡"):  '😍' is duplicated at positions 1 and 5.
}

/*
Aufgabe 2b:
zähle die Wort-Häufigkeiten eines Strings. Als Rückgabe verwendest du eine map,
mit den Wörten als Schlüssel und Anzahl als Wert.
Hinweis: verwende die Funktion strings.Fields und strings.ToLower
*/
func WordFrequencyCounter(input string) map[string]int {
	input = strings.ToLower(input)
	words := strings.Fields(input)
	result := make(map[string]int)

	for _, word := range words {
		result[word]++
	}

	return result
}
func ExampleWordFrequencyCounter() {
	//Output:
	// Word Frequencies:
	// hello: 2
	// is: 2
	// this: 4
	// world: 2
	text := "Hello world this is this this hello this is world"
	frequencies := WordFrequencyCounter(text)

	words := make([]string, 0, len(frequencies))
	for word := range frequencies {
		words = append(words, word)
	}

	sort.Strings(words)

	fmt.Println("Word Frequencies:")
	for _, word := range words {
		fmt.Printf("%s: %d\n", word, frequencies[word])
	}
}
