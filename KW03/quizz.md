# Go Programming Quiz

## 1. Was ist der Unterschied zwischen einem Array und einem Slice in Go?
<ol type="a">
 <li>Arrays haben eine feste Länge, Slices eine variable Länge</li>
 <li>Arrays können nur Zahlen speichern, Slices beliebige Datentypen</li>
 <li>Arrays können nicht als Funktionsparameter übergeben werden</li>
</ol>

## 2. Welche Aussage über die Übergabe eines Arrays als Funktionsparameter ist korrekt?
<ol type="a">
 <li>Arrays werden immer als Referenz übergeben</li>
 <li>Arrays werden standardmäßig als Kopie übergeben</li>
 <li>Die Übergabeart hängt von der Größe des Arrays ab</li>
 <li>Arrays können nicht als Parameter übergeben werden</li>
</ol>

## 3. Wie verhält sich ein Slice bei der Übergabe als Funktionsparameter?
<ol type="a">
 <li>Der Slice wird komplett kopiert</li>
 <li>Nur die Slice-Header-Struktur wird kopiert, die zugrundeliegenden Daten werden geteilt</li>
 <li>Die Übergabeart hängt von der Größe des Slice ab</li>
 <li>Slices können nicht als Parameter übergeben werden</li>
</ol>

## 4. Was ist die korrekte Syntax für eine For-Schleife über einen Slice?
```go
numbers := []int{1, 2, 3, 4, 5}
a. for i := 0; i < numbers.length; i++ { }
b. for (i = 0; i < len(numbers); i++) { }
c. for i := 0; i < len(numbers); i++ { }
d. for i to len(numbers) { }
```

## 5. Was bewirkt die folgende Operation?
```go
s1 := []int{1, 2, 3}
s2 := make([]int, len(s1), cap(s1)*2)
copy(s2, s1)
```

<ol type="a">
<li>Erstellt eine exakte Referenz von s1</li>
<li>Erstellt eine neue Slice mit doppelter Kapazität und kopiert die Elemente von s1</li>
<li>Vergrößert s1 um das Doppelte</li>
<li>Führt zu einem Kompilierfehler</li>
</ol>

## 6. Was ist das Ergebnis des folgenden Codes?
```go
s := make([]int, 3, 5)
fmt.Println(len(s), cap(s))
```
<ol type="a">
<li>3, 3</li>
<li>5, 5</li>
<li>3, 5</li>
<li>5, 3</li>
</ol>

## 7. Welche Aussage über li = append(li,1) ist korrekt?
<ol type="a">
<li>append() modifiziert den originalen Slice</li>
<li>append() gibt einen neuen Slice zurück</li>
<li>append() führt zu einem Fehler, wenn die Kapazität überschritten wird</li>
</ol>

## 8. Wie ist die Signatur dieser Funktion (Ein-&Ausgabe Parameter)?

```go
func x(int64, b int, int float64 ) (n int, err error) {
```
<ol type="a">
<li>Eingabe: x Typ int64, b Typ int, int Typ float64, Ausgabe: n Typ int, err Typ error</li>
<li>Eingabe: x Typ int64, b Typ int, int Typ float64, Ausgabe: (n Typ int, err Typ error)</li>
<li>Eingabe: int64 Typ int, b Typ int, int Typ float64, Ausgabe: n Typ int, err Typ error</li>
<li>Eingabe: int64 Typ int, b Typ int, int Typ int, float64 Typ float64, Ausgabe: n Typ int, err Typ error</li>
</ol>

