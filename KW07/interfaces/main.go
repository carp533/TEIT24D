package main

import (
	"fmt"
	"math"
)

/*Aufgabe 1: erstelle einen Datentyp, der ein Sechseck repräsentiert
und implementiere das interfage Geometer.
*/

/*Beispiel 1: geometrische Form*/
// Definition des Interface Geometer
type Geometer interface {
	Fläche() float64
	Umfang() float64
}

// Konkreter Typ Kreis
type Kreis struct {
	Radius float64
}

// Methode, um die Fläche eines Kreises zu berechnen
func (k Kreis) Fläche() float64 {
	return math.Pi * k.Radius * k.Radius
}

// Methode, um den Umfang eines Kreises zu berechnen
func (k Kreis) Umfang() float64 {
	return 2 * math.Pi * k.Radius
}

// Konkreter Typ Rechteck
type Rechteck struct {
	Länge, Breite float64
}

// Methode, um die Fläche eines Rechtecks zu berechnen
func (r Rechteck) Fläche() float64 {
	return r.Länge * r.Breite
}

// Methode, um den Umfang eines Rechtecks zu berechnen
func (r Rechteck) Umfang() float64 {
	return 2 * (r.Länge + r.Breite)
}

// Funktion, die die Eigenschaften eines Geometers ausgibt
func zeigeEigenschaften(g Geometer) {
	fmt.Printf("Fläche: %.2f\n", g.Fläche())
	fmt.Printf("Umfang: %.2f\n", g.Umfang())
	fmt.Println()
}

func Beispiel1() {
	// Erstellen eines Kreises
	kreis := Kreis{Radius: 5}
	// Erstellen eines Rechtecks
	rechteck := Rechteck{Länge: 4, Breite: 3}

	// Aufruf der Funktionen mit einem Interface
	// dieser Aufruf ist möglich, da die Datentypen Kreis und Recheck
	// die Methoden Fläche und Umfang bereit stellen
	// dies ist grundlegender Unterschied zu anderen Programmiersprachen!
	zeigeEigenschaften(kreis)
	zeigeEigenschaften(rechteck)
}

// Beispiel2: das leere interface: interface{}
// Funktion, die ein leeres Interface als Parameter akzeptiert
func anzeigen(obj interface{}) {
	fmt.Printf("Typ: %T, Wert: %v\n", obj, obj)
	//das geht nicht:
	//_ = obj + obj
}

func Beispiel2() {
	// Verschiedene Datentypen werden dem leeren Interface übergeben
	anzeigen(42)
	anzeigen("Hallo")
	anzeigen(3.14)
	anzeigen([]string{"A", "B", "C"})
	anzeigen(map[string]int{"Eins": 1, "Zwei": 2, "Drei": 3})
	fmt.Println()
}

/*Beispiel 3: Verwendung des eingebauten interface Stringer() und GoStringer() in go */
type Student struct {
	name  string
	matnr string
}

func (i Student) String() string {
	res := "da der Datentyp Student das interface Stringer() des Paketes fmt implementiert\n"
	res += "wird dieser Text angezeigt.\n"
	res += fmt.Sprintf("Der Name = %v\n", i.name)
	res += fmt.Sprintf("Matrikelnummer = :%v\n", i.matnr)
	return res
}
func (i Student) GoString() string {
	res := "da der Datentyp Student das interface GoStringer() des Paketes fmt implementiert\n"
	res += "wird für Formatangabe '%#v' ein anderer Text angezeigt.\n"
	return res
}

func Beispiel3() {
	stud1 := Student{name: "NoName", matnr: "12345"}
	fmt.Println(stud1)
	fmt.Printf("%#v\n", stud1)

}

func main() {
	Beispiel1()
	Beispiel2()
	Beispiel3()
}
