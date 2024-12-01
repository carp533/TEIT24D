// Angeben, zu welchem Package diese Datei gehört.
// Dateien mit einer main()-Funktion werden zu ausführbaren Dateien.
// Diese müssen immer zum Package "main" gehören.
package main

// Import-Anweisung. Wenn man Funktionen aus anderen Packages benutzen will,
// muss man das hier angeben. Wir verwenden unten die Funktion Println() aus dem Package
// fmt. Deshalb wird hier das Package fmt importiert.
import (
	"fmt"
)

// Die main()-Funktion.
// Jedes ausführbare Programm muss eine main()-Funktion haben.
// Sie ist der Einstiegspunkt, hier startet das Programm.
func main() {

	// Den String "Hallo Welt auf die Konsole ausgeben."
	// Println gibt den String aus und schreibt danach einen Zeilenumbruch( "Print line.").
	fmt.Println("Hello, World!")
}
