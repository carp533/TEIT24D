package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Printf("fact(7)=%v\n", fact(7))
	fmt.Printf("fib(7)=%v\n", fib(7))
	fmt.Printf("ackermann(3, 2)=%v\n", ackermann(3, 2)) //->29
	//fmt.Printf("ackermann(4, 1)=%v\n", ackermann(4, 1)) //->65533
}

// die Ackermann Funktion
func ackermann(m, n int) int {
	//TODO
	return 0
}
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// Prüft, ob eine Zahl gerade ist
func isEven(n int) bool {
	return false
}

// Prüft, ob eine Zahl ungerade ist
func isOdd(n int) bool {
	return false
}

func fibList(n int) []int {
	return nil
}

func fibBig(n uint64) *big.Int {
	return nil
}
