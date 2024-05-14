package main

import "fmt"

func main() {
	//Switch Tradicional en go
	switch expression := 53 % 2; expression {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}
	//Sin codicion
	value := 90

	switch {
	case value > 100:
		fmt.Println("Mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No hay codicional")
	}

}
