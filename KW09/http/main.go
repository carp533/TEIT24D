package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

/*Programm starten und im Browser http://localhostlocalhost:5000/hello aufrufen
ggf. muss man noch eine Sicherheitsabfrage der Windows-Firewall/Virenscanner bestätigen
*/

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":5000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
	fmt.Println("hello was called...")
	time.Sleep(time.Millisecond * time.Duration(100*(rand.Intn(5))))
}

func protokolliereZeit(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: rufe f auf und protokolliere die Zeit
	}
}

/*
	Wir wollen nun die Funktion hello "dekorieren",d.h.

Aufruf in main:

	http.HandleFunc("/hello", protokolliereZeit(hello))

func protokolliereZeit(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: rufe f auf und protokolliere die Zeit
	}
}
*/
