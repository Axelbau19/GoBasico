package main

import "fmt"

func main() {
	helloMessage := "Hello"
	worldMessage := "World"
	//Salto de linea
	fmt.Println(helloMessage, worldMessage)
	//PrintF
	nombre := "Axel"
	galletas := 3
	fmt.Printf("%s tiene %d galletas", nombre, galletas)
	//Sprintf
	message := fmt.Sprintf("%s tiene %d galletas", nombre, galletas)
	fmt.Println(message)
}
