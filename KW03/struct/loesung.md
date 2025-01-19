/*
AUFGABE: Structs für Notenverwaltung
Im Folgenden ist ein Datentyp für die Verwaltung von Noten gegeben.
Ändern Sie diesen Datentyp so ab, dass er anstelle einer einzelnen Note
eine Liste von Noten für verschiedene Teilleistungen enthält.
Passen Sie die Funktionen NewStudent(), Grades() und AddGrade() an.
*/

package main

import "fmt"

// Datentyp für einen Studenten mit Noten
type Student struct {
	Name   string
	Grades []float64
}

// Liefert eine neue Variable vom Typ Student
func NewStudent(name string, grades []float64) Student {
	return Student{Name: name, Grades: grades}
}

// Liefert eine String-Repräsentation eines Studenten
func (student Student) String() string {
	//Hinweis: verwende fmt.Sprintf
	return fmt.Sprintf("%s : %v", student.Name, student.GetGrades())
}

// Liefert ein Slice mit allen Noten des Studenten
func (student Student) GetGrades() []float64 {
	return student.Grades
}

// Fügt eine neue Note hinzu
func (student *Student) AddGrade(newGrade float64) {
	student.Grades = append(student.Grades, newGrade)
}

func (student Student) AverageGrade() float64 {
	sum := 0.0
	for _, v := range student.Grades {
		sum += v
	}
	return sum / float64(len(student.Grades))
}

func ExampleStudent_Grades() {
	s1 := NewStudent("Max Mustermann", []float64{1.3})
	fmt.Println(s1)

	s1.AddGrade(2.3)
	fmt.Println(s1)

	fmt.Printf("Durchschnitt: %v\n", s1.AverageGrade())
	fmt.Printf("Durchschnitt: %.2f\n", s1.AverageGrade())
	// Output:
	// Max Mustermann : [1.3]
	// Max Mustermann : [1.3 2.3]
	// Durchschnitt: 1.80
}
