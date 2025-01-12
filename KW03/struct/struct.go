package main

import (
	"fmt"
)

type Adresse struct {
	Strasse string
	Stadt   string
}
type Hobbies []string
type Kunde struct {
	name    string
	vorname string
	kdnr    int
	adresse Adresse
	hobbies Hobbies
}

func (k Kunde) String_() string {
	return fmt.Sprintf(
		"Kunde (%v):\n"+
			"  Name:    %v\n"+
			"  Vorname: %v\n",
		k.kdnr,
		k.name,
		k.vorname)
}
func main() {
	fmt.Print()

	// Zwei einzelne Variablen vom Typ Kunde definieren.
	k2 := Kunde{
		"Beispiel",
		"Barbara",
		77,
		Adresse{"Hauptstr.1", "Mannheim"},
		[]string{"Go-programmieren", "Kochen"},
	}
	//VIEL besser ist es, die Struktur mit Komponentennamen zu erzeugen (weniger Fehleranfällig)
	k3 := Kunde{
		vorname: "Monika",
		kdnr:    456,
		name:    "Musterfrau",
	}

	fmt.Println(k2)
	fmt.Println(k3)
	fmt.Printf("Ausgabe Struktur mit Komponentennamen:\n%+v\n", k3)

	// Eine Liste von Kunden:
	l1 := []Kunde{k2, k3, {
		vorname: "Manfred",
		kdnr:    456,
		name:    "MusterMann",
	}}

	fmt.Println(l1)

	// Kundennummer des 2. Kunden ausgeben:
	kunde2 := l1[1]
	kdnr_kunde2 := kunde2.kdnr
	fmt.Println(kdnr_kunde2)

	//Methodenaufruf für k3:
	fmt.Println(k3.String_())

}
