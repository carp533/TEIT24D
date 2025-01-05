# Arrays

Der Typ ```[n]T``` ist ein Array mit n Werten des Typs T.

Der Ausdruck

```go
var a [10]int
```

deklariert eine Variable a als ein Array von zehn int. Die Indizes der Elemente sind 0,1,...8,9

Die Länge eines Arrays ist Teil seines Typs, so dass Arrays nicht in ihrer Größe verändert werden können. Dies scheint einschränkend zu sein, aber keine Sorge, Go bietet eine bequeme Möglichkeit, mit Arrays zu arbeiten.

Arrays in Go haben eine feste Länge und einen festgelegten Datentyp. Sie werden als Werttyp behandelt, das bedeutet bei der Übergabe an Funktionen wird eine Kopie erstellt:

```go
// Array mit fester Länge 3
var arr [3]int = [3]int{1, 2, 3}

// Arrays werden als Werte übergeben - arr bleibt unverändert
func modify(a [3]int) {
    a[0] = 999 // Ändert nur die lokale Kopie
}
```

# Slices (Listen)
Ein Array hat eine feste Größe. Ein Slice hingegen ist eine dynamische, flexible Sicht auf die Elemente eines Arrays. In der Praxis sind Slices viel häufiger anzutreffen als Arrays. **Achtung: Listen sind nicht zu verwechseln mit einfach und doppelt verketteten Listen** (diese behandeln wir später).

Der Typ ```[]T``` ist ein Slice mit Elementen des Typs T.

Ein Slice wird durch die Angabe von zwei Indizes gebildet, einer unteren und einer oberen Grenze, die durch einen Doppelpunkt getrennt sind:

a[low : high]
Damit wird ein halboffener Bereich ausgewählt, der das erste Element einschließt, aber das letzte ausschließt.

Der folgende Ausdruck erzeugt ein Slice, das die Elemente 1 bis 3 von a einschließt:

```go
a[1:4]
```

## Slices sind Referenzen auf Arrays

Ein Slice speichert keine Daten, sondern beschreibt lediglich einen Abschnitt eines zugrunde liegenden Arrays.

Werden die Elemente eines Slices geändert, so werden die entsprechenden Elemente des zugrunde liegenden Arrays geändert.

Andere Slices, die dasselbe zugrunde liegende Array nutzen, sehen diese Änderungen.

Sie bestehen intern aus:
- Eine Referenz auf den zugrundeliegenden Array
- Der Länge (len)
- Der Kapazität (cap)

```go
// Slice erstellen
numbers := []int{1, 2, 3, 4, 5}

// make() zum Erstellen mit definierter Länge/Kapazität
slice := make([]int, 3, 5) // Länge 3, Kapazität 5

// Mehrere Slices können auf denselben Array zeigen
slice1 := numbers[1:3]  // [2,3]
slice2 := numbers[1:4]  // [2,3,4]
// Änderungen in slice1 sind auch in slice2 sichtbar!

// range zum Iterieren
for i, v := range numbers {
    fmt.Printf("Index: %d, Wert: %d\n", i, v)
}

// append() zum Hinzufügen von Elementen
numbers = append(numbers, 6, 7)
```

Wichtige Eigenschaften von Slices:
- Sie werden als Referenz übergeben - Funktionen können den Inhalt ändern
- append() erstellt bei Bedarf einen neuen, größeren Array
- Slices können wachsen und schrumpfen, anders als Arrays
- Die eingebaute len() Funktion gibt die aktuelle Länge zurück
- Mit cap() erhält man die Kapazität des zugrundeliegenden Arrays

```go
// Slices werden als Referenz übergeben
func modifySlice(s []int) {
    s[0] = 999 // Ändert das Original!
}

slice := []int{1, 2, 3}
modifySlice(slice)
// slice ist jetzt [999, 2, 3]
```
