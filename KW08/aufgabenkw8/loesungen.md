package main

import "fmt"

/* Alle Tests sollen funktionieren - Testfunktionen nicht ändern!

Aufgabe 1a:
berechne die "Hailstone sequence", definiert durch:
Starte mit einer positiven Zahl n:
Wenn n == 1 Ende.
Wenn n   gerade ist: n -> n/2
Wenn n ungerade ist: n -> (3 * n) + 1
*/
func Hs(n int) []int {
	res := make([]int, 0)
	res = append(res[:0], n)
	for n > 1 {
		if n&1 == 0 { //<- warum funktioniert das? (& ist der Bitweise UND Operator)
			n = n / 2
		} else {
			n = 3*n + 1
		}
		res = append(res, n)
	}
	return res
}

func ExampleHs() {
	//Output:
	//[5 16 8 4 2 1]
	seq := Hs(5)
	fmt.Println(seq)
}

/*Aufgabe 1b:
berechne alle Hailstone Sequenzen < 100.000 und gib den Startwert und Länge der
längsten Sequenz aus.
*/
func HsMax(max int) (maxN, maxLen int) {
	maxN, maxLen = 1, 1
	for i := 1; i <= max; i++ {
		l := len(Hs(i))
		if l > maxLen {
			maxN, maxLen = i, l
		}
	}
	return maxN, maxLen
}

func ExampleHsMax() {
	//Output:
	//hs(77031): 351 Elemente
	maxN, maxLen := HsMax(100000)
	fmt.Printf("hs(%d): %d Elemente\n", maxN, maxLen)
}


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


package main

import (
	"fmt"
)

/* Aufgabe 3a:
implementiere eine Queue
wir benötigen die grundlegenden Operationen Enqueue, Dequeue, IsEmpty und Size.
Die Enqueue-Methode fügt ein Element am Ende der Queue hinzu,
Dequeue entfernt und gibt das erste Element aus der Queue zurück,
IsEmpty prüft, ob die Queue leer ist, und
Size gibt die Anzahl der Elemente in der Queue zurück.
verwende eine int-Liste zum verwalten der Elemente der Queue
*/
/*Aufgabe 3b:
baue die Queue um, sodass beliebige Elemente verwendet werden können.
welchen Datentyp musst du hierfür benutzen?
*/
type Queue struct {
	items []int
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the first element from the queue
func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// IsEmpty returns true if the queue is empty, false otherwise
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return len(q.items)
}

func ExampleQueue() {
	//Output:
	// Dequeued item: 10
	// Dequeued item: 20
	// Queue size: 1
	// Is queue empty? false
	// queue is empty
	queue := Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	v, err := queue.Dequeue()
	fmt.Println("Dequeued item:", v)
	v, err = queue.Dequeue()
	fmt.Println("Dequeued item:", v)
	fmt.Println("Queue size:", queue.Size())
	fmt.Println("Is queue empty?", queue.IsEmpty())
	v, err = queue.Dequeue()
	v, err = queue.Dequeue()
	fmt.Println(err)
}

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
	primesiev := make([]bool, N+1)
	for i := range primesiev {
		primesiev[i] = true
	}
	for i := 2; i*i <= N; i++ {
		if primesiev[i] {
			for j := i * 2; j <= N; j = j + i {
				primesiev[j] = false
			}
		}
	}
	result := make([]int, 0)
	for i, isprime := range primesiev {
		if isprime && i > 1 {
			result = append(result, i)
		}
	}
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
