package main

import "fmt"

func main() {

	const nihongo = "ABC 日本語"
	fmt.Println("UTF-8 String:")
	fmt.Printf("%v\n", (nihongo))
	fmt.Println("String als byte-Folge:")
	fmt.Printf("%v\n", []byte(nihongo))
	fmt.Println("String als byte-Folge (hexadezimal):")
	fmt.Printf("% x\n", nihongo)
	fmt.Println("String als rune-Folge:")
	fmt.Printf("%v\n", []rune(nihongo))
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	for i := 0; i < len(nihongo); i++ {
		fmt.Printf("%v ", nihongo[i])
	}
	fmt.Println()
	for i := 0; i < len([]rune(nihongo)); i++ {
		fmt.Printf("%v ", nihongo[i])
	}

}
