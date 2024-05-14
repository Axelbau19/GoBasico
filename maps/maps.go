package main

import "fmt"

func main() {

	//Aqui una manera mas usada
	personas := map[string]int{
		"Axel":   22,
		"Alexis": 22,
	}

	fmt.Println(personas)
	//Recorrido de un mapa
	for i, valor := range personas {
		fmt.Println(i, valor)
	}
	//Obtener  o encontrar un valor
	valueAxel, existeValor := personas["Axel"]
	fmt.Println(valueAxel, existeValor)
	valueNoencontrado, existeValor2 := personas["xd"]
	fmt.Print(valueNoencontrado, existeValor2)

}
