package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}
func tripeArgumento(a, b int, c string) {
	fmt.Print(a, b, c)
}
func valorDoble(numero int) int {
	return numero * 2
}

//Para funciones para devolver 1 o mas respuestas
func doubleReturn(a int) (b, c int) {
	return a, a * 2
}

//Area Cuadrado
func areaCuadrado(lado float64) float64 {
	return lado * lado
}

//Area de un rectangulo
func areaRectangulo(ancho, altura float64) float64 {
	return ancho * altura
}

func main() {
	normalFunction("Hola mundo")
	tripeArgumento(1, 2, "Hola mundo")
	value := valorDoble(2)
	fmt.Println("value: ", value)
	//Para guardar 2 variables que retorna 2 resultados
	value1, value2 := doubleReturn(2)
	fmt.Println("Value1 y Value 2: ", value1, value2)
	//Si una funcion regresa mas 1 resultados pero solamente ocupa 1 (o bueno la cantidad que requieras), se puede guardar como
	_, value3 := doubleReturn(1)
	fmt.Println(value3)
	areaCuadrado := areaCuadrado(2)
	fmt.Println("Area de un cuadrado:", areaCuadrado)
	areaRectangulo1 := areaRectangulo(10, 2)
	fmt.Println(areaRectangulo1)
}
