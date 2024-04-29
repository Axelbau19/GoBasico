package main

import "fmt"

func main() {
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)
	//Declarar variables
	base := 2
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)
	//Zero Values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)
	//Area Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)
	//Aritmetrica
	x := 5
	y := 10
	//suma
	resutado := x + y
	fmt.Println(resutado)
	//resta
	resutado = y - x
	fmt.Println(resutado)
	//multiplicacion
	resutado = x * y
	fmt.Println(resutado)
	//Divison
	resutado = y / x
	fmt.Println(resutado)
	//Modulo
	resutado = y % x
	fmt.Println(resutado)
	//Area Rectangulo
	const baseRectangulo float64 = 10
	const alturaRectangulo float64 = 20
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println(areaRectangulo)
}
