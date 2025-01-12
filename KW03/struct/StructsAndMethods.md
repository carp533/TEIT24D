# Strukturen und Methoden in Go

## Strukturen (Structs)
In Go werden Strukturen verwendet, um verschiedene Datentypen zu einer Einheit zusammenzufassen. Eine Struktur ist wie ein Container für zusammengehörige Daten. Die Komponenten einer Struktur bestehen aus einem Namen und einem Datentyp.

```go
// Definition einer Struktur
type Student struct {
    Name    string
    Alter   int
    Matrikel int
}

// Erstellen einer neuen Struktur
student1 := Student{
    Name: "Maria",
    Alter: 20,
    Matrikel: 12345,
}

// Alternative Erstellung
var student2 Student
student2.Name = "Peter"
student2.Alter = 22
student2.Matrikel = 67890
```

## Methoden
Methoden sind Funktionen, die an eine Struktur gebunden sind. Sie können auf die Daten der Struktur zugreifen und diese verarbeiten. 

```go
// Methode zur Struktur Student hinzufügen
func (s Student) Vorstellen() string {
    return fmt.Sprintf("Hallo, ich bin %s und %d Jahre alt.", s.Name, s.Alter)
}

//vergleiche hierzu:
func Vorstellen_(s Student) string {...}

// Methode mit Referenz zum Ändern von Daten
func (s *Student) AlterErhoehen() {
    s.Alter++
}

// Verwendung der Methoden
func main() {
    student1 := Student{Name: "Maria", Alter: 20, Matrikel: 12345}
    
    // Methode aufrufen
    fmt.Println(student1.Vorstellen())
    // Ausgabe: Hallo, ich bin Maria und 20 Jahre alt.
    
    student1.AlterErhoehen()
    fmt.Printf("Neues Alter: %d\n", student1.Alter)
    // Ausgabe: Neues Alter: 21
}
```

## Wichtige Konzepte:

1. Strukturen gruppieren zusammengehörige Daten
2. Methoden können auf zwei Arten definiert werden:
   - Mit Wert-Empfänger `func (s Student)` - können Daten nur lesen
   - Mit Zeiger-Empfänger `func (s *Student)` - können Daten auch ändern
3. Methoden machen den Code übersichtlicher 

