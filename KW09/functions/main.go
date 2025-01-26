package main

import "fmt"

func main() {
	//DEMO_Variadic()
	//DEMO_AnonymeFunktionen()
	//DEMO_FunktionenAlsParameter()
	DEMO_Closures()
}

/*Variadic Functions
variadische Funktionen können eine variable Zahl von Eingaben empfangen.
Dies wird in der Signatur mit ...Datentyp gemacht.
fmt.Println ist eine variadische Funktionen*/
func variadic(i ...int) {
	//i ist kein int, sondern eine Liste von int
	fmt.Printf("i:%v, %T\n", i, i)
	for _, v := range i {
		fmt.Printf("%v ", v)
	}
}

func DEMO_Variadic() {
	liste := []int{1, 2, 3, 4, 5}
	variadic(1, 2, 3)  //Aufruf mit drei int
	variadic(liste...) //liste... "entpackt" die Liste in einzelne Werte
}

/*Anonyme Funktion*/
func DEMO_AnonymeFunktionen() {

	// achte auf die Klammern ()
	func() {
		fmt.Println("ich bin eine anonyme Funktion")
	}() //<- die Funktion wird sofort aufgerufen

	// anonyme Funktion wird einer Variablen zugewiesen
	a := func() {
		fmt.Println("ich bin eine anonyme Funktion II")
	}
	a() //<- hier wird die Funktion aufgerufen
	fmt.Printf("Signatur der Funktion a: %T\n", a)

	// anonyme Funktion mit Eingabeparameter
	func(n string) {
		fmt.Println("anonyme Funktion mit Eingabeparameter. Hello", n)
	}("TEL23B") //wird direkt mit ("TEL23B") aufgerufen

}

/*Funktionen als Ein- und Ausgabeparameter*/
// die Funktion funcInput hat eine Funktion als Eingabe
func funcInput(a func(x, y int) int) {
	fmt.Println("Funktion als Eingabe. ", a(4, 5))
}

// die Funktion funcOutput hat eine Funktion als Rückgabeparameter
func funcOutput(x int) func(a, b int) int {
	f := func(a, b int) int {
		return x * (a + b)
	}
	return f
}

func DEMO_FunktionenAlsParameter() {
	// Funktionen als Eingabeparameter
	f := func(a, b int) int {
		return a * b
	}
	// die Funktion funcInput hat eine Funktion (f) als Eingabe
	funcInput(f)

	f2 := funcOutput(2)
	fmt.Println("Funktion als Ausgabe. ", f2(6, 7))

}

/*Closure*/
func DEMO_Closures() {

	//closure: eine anonyme Funktion, die auf eine "äußere" Variable zugreift
	z := 5
	func() {
		fmt.Println("z =", z)
	}()


	//closure -> jede closure "hält" ihre Variablen
	c1 := appendStr()
	c2 := appendStr()

	fmt.Printf("Datentyp c1: %T\n", c1)
	fmt.Printf("Datentyp appendStr: %T\n", appendStr)

	fmt.Println(c1("World"))
	fmt.Println(c2("Everyone"))

	fmt.Println(c1("Gopher"))
	fmt.Println(c2("!"))

}

// Beispiel für closure & Gültigkeit Variablen
func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
