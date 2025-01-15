## Antworten:
1. a
2. b
3. b
4. c
5. b
6. c
7. a
8. c

## Erklärungen:

1. Arrays haben eine feste Länge, die Teil des Typs ist (z.B. [5]int). Slices sind dynamisch und können wachsen.

2. In Go werden Arrays als Wert übergeben, was bedeutet, dass eine komplette Kopie erstellt wird. Um dies zu vermeiden, kann man einen Pointer übergeben.

3. Bei Slices wird nur die Slice-Header-Struktur kopiert (enthält Pointer zum Array, Länge und Kapazität). Die zugrundeliegenden Daten werden geteilt.

4. Die korrekte For-Schleife in Go verwendet den := Operator für die Initialisierung und len() für die Länge.

5. make() erstellt einen neuen Slice mit angegebener Länge und Kapazität. copy() kopiert dann die Elemente von s1 nach s2.

6. make([]int, 3, 5) erstellt einen Slice mit Länge 3 und Kapazität 5. len() gibt die Länge zurück, cap() die Kapazität.

