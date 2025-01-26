# Aufgaben 1: einfach verkettete Liste
(1) implementiere die Methode "Size" rekursiv

	func (ll LinkedList) Size() int
	Hilfsfunktion func (node Node) size() int

(2) implementiere die "IsInList" Methode

	wie sieht die Signatur dieser Methode aus?

(3) implementiere die Methode "RemoveValue"

	wie sieht die Signatur dieser Methode aus?
		a) lösche den ersten gefundenen Wert
		b) lösche alle gefundenen Werte

# Aufgabe 2: "Example-Funktionen"
Schreibe zu Size, IsInList, RemoveValue jeweils eine Example Funktionen, 
in der du die Methode testest. Beachte, dass die Datei mit ```_test.go```
enden, der Funktionsname mit ```Example``` beginnen und ```// Output:```
der letzte Kommentar sein muss.

Siehe hierzu [https://pkg.go.dev/testing#hdr-Examples](https://pkg.go.dev/testing#hdr-Examples) (nur den Abschnitt Examples!)
```go
func ExampleLinkedList_Size() {
	// Output:
	// 3
}
```

# Aufgabe 3: Stacks
Versuche die Aufgaben zu lösen, OHNE in die Musterlösung zu schauen.

1. Vergleiche den Datentyp Stack mit dem Datentyp verkettete Liste.
2. Programmiere den Stack mit
	1. der Methode Length ->Push +1, Pop -1
	2. allgemeinen Daten statt int -> interface{}. Teste das, indem du Push() mit verschiedenen Datentypen aufrufst.
        
 - teste das Programm, einmal mit und einmal ohne obige String Methode
 - programmiere die rekursive Methode String für die Ausgabe der Knoten.        
