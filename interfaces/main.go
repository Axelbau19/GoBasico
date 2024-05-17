package main

import "fmt"

type FiguraGeometrica interface {
	area() float64
}

type Cuadrado struct {
	Lado float64
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

func (cuadradoObj Cuadrado) area() float64 {
	return cuadradoObj.Lado * cuadradoObj.Lado
}

func (rectanguloObj Rectangulo) area() float64 {
	return rectanguloObj.Base * rectanguloObj.Altura
}

func calcularArea(figura FiguraGeometrica) {
	fmt.Println("Area: ", figura.area())
}
func main() {
	miCuadrado := Cuadrado{Lado: 5}
	miRectangulo := Rectangulo{Base: 10, Altura: 20}

	calcularArea(miCuadrado)
	calcularArea(miRectangulo)
	//Lista interfaces
	myInterface := []interface{}{"Hola", 12, 2.1}
	fmt.Println(myInterface...)

}
