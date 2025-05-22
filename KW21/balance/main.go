package main

// Aufgabe: Prüfe auf balancierte Markup-Tags
// Ziel:
// Implementiere eine Funktion in Go, die prüft, ob eine gegebene Zeichenkette korrekt ausbalancierte Markup-Tags enthält.
// Die Zeichenkette kann Tags wie <open> und </open>, <start> und </start> enthalten. Welche Datenstruktur ist geeignet,
// um die Aufgabe zu lösen?

// Beschreibung:
// Schreibe eine Funktion IsBalancedTags, die eine Zeichenkette als Eingabeparameter nimmt und einen booleschen Wert zurückgibt.
// Die Funktion soll true zurückgeben, wenn die Markup-Tags in der Zeichenkette korrekt ausbalanciert sind, andernfalls false.
// Ein Markup-Tag ist korrekt ausbalanciert, wenn jedes öffnende Tag ein entsprechendes schließendes Tag hat und die Tags richtig verschachtelt sind.
// Beispiel:
// Input:

// Zeichenkette: "<open><start></start></open>"
// Output:

// true
// Input:

// Zeichenkette: "<open><start></open></start>"
// Output:

// false

//Hinweis: prüfe den String Zeichen für Zeichen. verwende z.B.
//strings.TrimSuffix(strings.TrimPrefix(tag, "</"), ">")
//um den Namen des schließenden Tag zu ermitteln
import (
	"fmt"
)

func IsBalancedTags(s string) bool {
	stack := []string{}
	return len(stack) == 0 // Alle Tags müssen geschlossen sein
}

func main() {
	// Beispielzeichenketten
	examples := []string{
		"<open><start></start></open>",
		"<open><start></open></start>",
		"<open></open>",
		"<start><open></open></start>",
		"<start><open></start></open>",
		"<open><start></start>",
		"</open><start></start>",
	}

	for _, example := range examples {
		fmt.Printf("Zeichenkette: %s, Ausbalanciert: %v\n", example, IsBalancedTags(example))
	}
}
