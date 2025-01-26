package main

import "fmt"

func main() {
	// Eine Map erstellen, um Altersinformationen zu speichern
	ages := make(map[string]int)

	// Werte zur Map hinzufügen
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 22

	// Auf Werte in der Map zugreifen
	fmt.Println("Das Alter von Alice:", ages["Alice"])
	fmt.Println("Das Alter von Bob:", ages["Bob"])
	fmt.Println("Das Alter von Charlie:", ages["Charlie"])

	// Überprüfen, ob ein Schlüssel existiert
	if age, ok := ages["Dave"]; ok {
		fmt.Println("Das Alter von Dave:", age)
	} else {
		fmt.Println("Dave existiert nicht in der Map.")
	}

	// Iteration über die Map
	fmt.Println("Iteration über die Map:")
	for name, age := range ages {
		fmt.Printf("%s ist %d Jahre alt\n", name, age)
	}

	// Ein Element aus der Map entfernen
	delete(ages, "Charlie")
	fmt.Println("Iteration über die Map, nach dem Entfernen von Charlie:")
	for name, age := range ages {
		fmt.Printf("%s ist %d Jahre alt\n", name, age)
	}
}
