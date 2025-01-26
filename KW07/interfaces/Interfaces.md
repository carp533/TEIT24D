# Interfaces
Ein Interface I ist definiert durch eine Menge von Methodensignaturen.
Das Interface ist nicht direkt mit einem Datentypen gekoppelt.
Wir sagen: ein Datentyp T ist vom Typ Interface I, wenn T alle Methoden des Interfaces I bereit stellt.

In go ist es üblich, möglichst wenig Methoden in einem Interface zusammen zu fassen (durch "kleine" interfaces fördert man das Schreiben modularer, wiederverwendbarer und flexibler Quelltexte).

Das leere interface (```interface{}```, siehe [https://go.dev/tour/methods/14](https://go.dev/tour/methods/14)) kann jeden Datentyp enthalten.

typische Beispiele:

```go
//https://go.dev/ref/spec#Predeclared_identifiers
//immer vorhanden:
type error interface {
	Error() string
}

//in package sort: 
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//in package io: 
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}
```
