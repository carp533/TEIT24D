# Funktionen

## Variadic Functions
variadische Funktionen können eine variable Zahl von Eingaben empfangen.
Dies wird in der Signatur mit ...Datentyp gemacht. In der Funktion kann ist
die Eingabe dann eine Liste: []Datentyp.
```fmt.Println``` ist eine variadische Funktionen.
Mit dem ```...``` Operator kann man eine Liste in "entpacken".

```go
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
```

# anonyme Funktionen
anonyme Funktionen sind Funktionen ohne Namen. Sie können einer Variablen
zugewießen werden oder als Ein- oder Ausgabeparameter verwedent werden.
Beachte die Klammern: mit "()" wird die Funktion direkt aufgerufen.
```go
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
	}("TEIT24") //wird direkt mit ("TEIT24") aufgerufen

}
```
# Closures
ein closure ist eine anonyme Funktion, die noch auf einen Wert außerhalb der
Funktion zugreift. Jede closure hält ihre Variablen.
```go
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

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
```
# Funktionen als Ein-/Ausgabe
```go
func funcInput(a func(x, y int) int) {
	fmt.Println("Funktion als Eingabe. ", a(4, 5))
}

func funcOutput(x int) func(a, b int) int {
	f := func(a, b int) int {
		return x * (a + b)
	}
	return f
}
```
