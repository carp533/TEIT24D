package main

import (
	"fmt"
)

/* 1) Pointer Grundlagen*/
func pointerBasic() {
	i, j := 42, 2701

	p := &i                                        // definiert einen Pointer auf die Variable i
	*p = 21                                        // (Schreib-)Zugriff auf den Wert des Pointer (Inhalt Speicher)
	fmt.Printf("Der Wert von    p (*p): %v\n", *p) // (Lese-)Zugriff auf den Wert des Pointer (Inhalt Speicher)
	fmt.Printf("Die Adresse von p (p) : %v\n", p)  //
	fmt.Printf("Der Wert von    i (i) : %v\n", i)  //
	fmt.Printf("Die Adresse von i (&i): %v\n", &i) // Zugriff auf den Speicher, in dem i steht
	p = &j                                         // p zeigt nun auf die Variable j
	*p = *p / 37                                   // ändere den Wert (Inhalt Speicher)
	fmt.Println(j)

	var k *int            // definiert einen Pointer mit Datentyp int
	fmt.Printf("%v\n", k) //ein initialer Pointer ist <nil> -> hiermit kann man nichts machen
	//*k = 42             // -> Programm bricht ab, weil k noch auf keinen Speicherbereich zeigt

}

/* 2) Pointer und Funktionen, Übergabe als Wert/Referenz */
func fw(s string) {
	fmt.Printf("in Funktion fw, Wert von s (vor Änderung): %v\n", s)
	s = "Funktion fw" //die Änderung ist nur innerhalb der Funktion zu sehen!
	fmt.Printf("in Funktion fw, Wert von s: %v\n", s)
}

func fr(p *string) {
	fmt.Printf("in Funktion fr, Wert von s (vor Änderung): %v\n", p)
	*p = "Funktion fr"
	fmt.Printf("in Funktion fr, Wert von p: %v\n", *p)
}

func DemoWertReferenz() {
	str := "test"
	fw(str)
	fmt.Printf("in Funktion main, Wert von str: %v\n", str)

	fr(&str)
	fmt.Printf("in Funktion main, Wert von str: %v\n", str)
}

/*
3. Pointer und Methoden, Übergabe als Wert/Referenz
  - Methoden werden in GO häufig verwendet, um zu signalisieren, dass eine Funktion für einem
  - bestimmten Datentyp definiert ist.
    var x MyStruct
    func noChange(m MyStruct)		-> Funktion
    noChange(x)					    -> Aufruf der Funktion
    func (m MyStruct) noChange()    -> Methode
    x.noChange()					-> Aufruf der Methode
*/
type MyStruct struct {
	i1 int
}

func (m *MyStruct) change() {
	m.i1 += 100
}
func (m MyStruct) noChange() {
	m.i1 += 100 // Hat keinen Effekt, weil m nicht als Pointer vorliegt.
}

func pointerMeth() {
	s1 := MyStruct{10}
	fmt.Println(s1)
	s1.change()
	fmt.Println(s1)
	s1.noChange()
	fmt.Println(s1)
}

func main() {

	pointerBasic()
	DemoWertReferenz()
	pointerMeth()
	aufgabe()

}

func aufgabe() {
	/* schaue dir das Programm an und ergänze die Kommentare. Überprüfe
	*  ob sie stimmen.	*/
	v1 := 42
	p_v1 := &v1
	fmt.Println(v1, p_v1, *p_v1) // Gibt TODO aus.

	v1 = 23
	fmt.Println(v1, p_v1, *p_v1) // TODO hat sich verändert.

	*p_v1 = 77
	fmt.Println(v1, p_v1, *p_v1) // TODO hat sich verändert.

	changePtr(p_v1)              // TODO Beschreibung.
	fmt.Println(v1, p_v1, *p_v1) // TODO hat sich verändert.

	p2 := new(int)       // Einen Pointer mit new erzeugen.
	fmt.Println(p2, *p2) // TODO Bescreibung.
	changePtr(p2)
	fmt.Println(p2, *p2) // TODO Bescreibung.

}

// Beispiel-Funktion: Erwartet einen int-Pointer und ändert den Wert dahinter.
func changePtr(i_ptr *int) {
	*i_ptr += 10
}
