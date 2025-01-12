/*
AUFGABE: Structs für Notenverwaltung
Im Folgenden ist ein Datentyp für die Verwaltung von Noten gegeben.
Ändern Sie diesen Datentyp so ab, dass er anstelle einer einzelnen Note
eine Liste von Noten für verschiedene Teilleistungen enthält.
Passen Sie die TODO's an.
*/

package main

import "fmt"

// Datentyp für einen Studenten mit Noten
type Student struct {
	Name  string
	Grade float64 //TODO
}

// Liefert eine neue Variable vom Typ Student
func NewStudent(name string, grades []float64) Student {
	//TODO
	return Student{}
}

// Liefert eine String-Repräsentation eines Studenten
func (student Student) String() string {
	//Hinweis: verwende fmt.Sprintf
	return "" //TODO
}

// Liefert ein Slice mit allen Noten des Studenten
func (student Student) GetGrades() []float64 {
	return []float64{} //TODO
}

// Fügt eine neue Note hinzu
func (student *Student) AddGrade(newGrade float64) {
	//TODO
}

func (student Student) AverageGrade() float64 {
	//TODO
	return 0.0
}

func ExampleStudent_Grades() {
	s1 := NewStudent("Max Mustermann", []float64{1.3})
	fmt.Println(s1)

	s1.AddGrade(2.3)
	fmt.Println(s1)

	//TODO korrigiere die Ausgabe, sodass nicht 1.7999999999999998 sondern 1.80 ausgegeben wird
	//Hinweis: verwende das richtige Format 'verb' , siehe https://pkg.go.dev/fmt
	fmt.Printf("Durchschnitt: %v\n", s1.AverageGrade())
	// Output:
	// Max Mustermann : [1.3]
	// Max Mustermann : [1.3 2.3]
	// Durchschnitt: 1.80
}
