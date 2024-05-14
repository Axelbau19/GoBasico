package main

import "fmt"

//Declaramos la clase
type carro struct {
	nombre string
	anio   int
}

func main() {
	supra := carro{
		nombre: "Supra",
		anio:   2024,
	}
	fmt.Println(supra)

	//Otra manera
	var gt carro
	gt.nombre = "Gran turismo"
	gt.anio = 2024
	fmt.Println(gt)
}
