package main

import (
	"fmt"
	"strings"
)

/*Hier sind einige Übungaufgaben zur Klausurvorbereitung.
In der Klausur wird es vier Aufgaben geben. */

func main() {
	ExampleCommonMultiples()
	ExampleArraySums()
	ExamplePower2()
	ExampleFindAll()
	ExampleDictionary_SetEntry()
	ExampleEntry_Translations()
	ExampleDiv()
}

/*
	AUFGABE 1: Listen -  Erreichbare Punkte: 10

// AUFGABENSTELLUNG:
// Erwartet drei Zahlen m,n,max > 0 und liefert eine Liste mit allen
// gemeinsamen Vielfachen von m und n, die nicht größer als max sind.
*/
func CommonMultiples(m, n, max int) []int {
	result := []int{}
	// TODO
	return result
}
func ExampleCommonMultiples() {
	fmt.Println(CommonMultiples(3, 5, 50))
	fmt.Println(CommonMultiples(2, 10, 100))
	fmt.Println(CommonMultiples(1, 1, 10))

	// Output:
	// [15 30 45]
	// [10 20 30 40 50 60 70 80 90 100]
	// [1 2 3 4 5 6 7 8 9 10]
}

/*
	AUFGABE 2: Listen - Erreichbare Punkte: 10

// AUFGABENSTELLUNG:
// Erwartet eine int-Slice list.
// Liefert eine int-Slice, die an Stelle n die Summe
// der Elemente aus list bis zu Stelle n enthält.
*/
func ArraySums(list []int) []int {
	result := []int{}
	// TODO
	return result
}

func ExampleArraySums() {
	fmt.Println(ArraySums([]int{1, 3, 5, 7}))
	fmt.Println(ArraySums([]int{1, 1, 2, 80}))
	fmt.Println(ArraySums([]int{7, 3, 1, 2}))
	fmt.Println(ArraySums([]int{3, 3, 0, 2}))

	// Output:
	// [1 4 9 16]
	// [1 2 4 84]
	// [7 10 11 13]
	// [3 6 6 8]
}

/*
	AUFGABE 3: Rekursion - Erreichbare Punkte: 10

// AUFGABENSTELLUNG:
// Liefert die Potenz "2 hoch x".
// Zusatzanforderung: Die Funktion muss rekursiv sein.
*/
func Power2(x int) float64 {
	// TODO
	return 0.0
}
func ExamplePower2() {

	fmt.Println(Power2(2))
	fmt.Println(Power2(5))
	fmt.Println(Power2(0))
	fmt.Println(Power2(-2))

	// Output:
	// 4
	// 32
	// 1
	// 0.25
}

/*
	AUFGABE 4: Listen -  Erreichbare Punkte: 10

// AUFGABENSTELLUNG:
// Liefert eine Liste mit allen Stellen in list,
// an denen x vorkommt.
*/
func FindAll(list []int, x int) []int {
	result := []int{}
	// TODO
	return result
}
func ExampleFindAll() {
	l1 := []int{1, 4, 17, 2, 1, 5, 10, 5, 2}
	fmt.Println(FindAll(l1, 1))
	fmt.Println(FindAll(l1, 4))
	fmt.Println(FindAll(l1, 5))
	fmt.Println(FindAll(l1, 42))

	// Output:
	// [0 4]
	// [1]
	// [5 7]
	// []
}

/* AUFGABE 5: Structs - Erreichbare Punkte: 20
// AUFGABENSTELLUNG:
// Reparieren Sie die folgende Methode.
// Es müssen alle Tests grün werden.

// Fügt ein Element zu dict hinzu,
// das das Wortpaar (de,en) enthält.
// Falls dict schon einen Eintrag für de enthält,
// soll dessen en ersetzt werden. */

func (dict *Dictionary) SetEntry(de, en string) {
	//TODO
}

type Entry struct {
	De string
	En string
}

func (entry Entry) String() string {
	return fmt.Sprintf("%s : %s", entry.De, entry.En)
}

type Dictionary struct {
	Entries []Entry
}

func (dict Dictionary) String() string {
	lines := []string{}
	for i, entry := range dict.Entries {
		lines = append(lines, fmt.Sprintf("%d - %s", i, entry))
	}
	return strings.Join(lines, "\n")
}
func ExampleDictionary_SetEntry() {
	d1 := Dictionary{[]Entry{}}

	d1.SetEntry("Himmel", "sky")
	fmt.Println(d1)
	fmt.Println()

	d1.SetEntry("Erde", "earth")
	fmt.Println(d1)
	fmt.Println()

	d1.SetEntry("Himmel", "heaven")
	fmt.Println(d1)
	fmt.Println()

	// Output:
	// 0 - Himmel : sky
	//
	// 0 - Himmel : sky
	// 1 - Erde : earth
	//
	// 0 - Himmel : heaven
	// 1 - Erde : earth
}

/* AUFGABE 6: Structs - Erreichbare Punkte: 20 */

// AUFGABENSTELLUNG:
// Im Folgenden ist ein Datentyp für die Einträge eines Wörterbuchs gegeben.
// Ändern Sie diesen Datentyp so ab, dass er anstelle eines einzelnen englischen
// Worts eine Liste von Wörtern enthält.
// Passen Sie die Funktionen NewEntry(), Translations() und AddTranslation() an.

// Datentyp für Einträge eines Wörterbuchs.
type Eintrag struct {
	De string
	En []string
}

// Liefert einen neuen Eintrag.
func NewEntry(de string, en []string) Eintrag {
	return Eintrag{De: de, En: en}
}

// Liefert eine String-Repräsentation eines Eintrags.
func (entry Eintrag) String() string {
	return fmt.Sprintf("%s : %v", entry.De, entry.Translations())
}

// Liefert einen String mit allen englischen Wörtern aus entry.
// Die einzelnen Wörter sollen mit Kommata getrennt sein.
func (entry Eintrag) Translations() []string {
	return entry.En
}

// Fügt eine neue Übersetzung zu entry hinzu.
func (entry *Eintrag) AddTranslation(newEn string) {
	entry.En = append(entry.En, newEn)
}

func ExampleEntry_Translations() {
	e1 := NewEntry("Himmel", []string{"sky"})
	fmt.Println(e1)

	e1.AddTranslation("heaven")
	fmt.Println(e1)

	// Output:
	// Himmel : sky
	// Himmel : sky,heaven
}

/* AUFGABE 7 : Rekursion - Erreichbare Punkte: 15 */
// AUFGABENSTELLUNG:
// Reparieren Sie die folgende Funktion.
// Es müssen alle Tests grün werden.
// Zusatzanforderung: Die Funktion muss rekursiv bleiben.

// Berechnet den ganzzahligen Quotienten x / y.
// D.h. die Funktion ignoriert Nachkommastellen bzw. den Rest.
func Div(x, y int) int {
	return 0
}
func ExampleDiv() {

	fmt.Println(Div(3, 2))
	fmt.Println(Div(2, 3))
	fmt.Println(Div(20, 2))
	fmt.Println(Div(-3, 2))
	fmt.Println(Div(3, -2))
	fmt.Println(Div(-3, -2))

	// Output:
	// 1
	// 0
	// 10
	// -1
	// -1
	// 1
}
