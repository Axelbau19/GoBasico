// Puntero guarda el lugar de memoria y el dato
package main

import "fmt"

type Computadora struct {
	Cpu          string
	Ram          int
	TarjetaVideo string
}

func (myPc Computadora) ping() {
	fmt.Println(myPc.Cpu, "Pong")
}

func (myPc *Computadora) duplicateRam() {
	myPc.Ram *= 2
}

func main() {
	a := 50
	//Se obtiene la direccion de memoria con el &
	b := &a
	//y conviertes o obtienes el valor el *
	fmt.Println(a)
	//Direccion memoria
	fmt.Println(b)
	//Para obtener el valor de memoria
	fmt.Println(*b)
	//Reasignar un valor
	*b = 100
	//tambien cambia el valor de la a
	fmt.Println(a)
	fmt.Println(*b)
	//Instancia valor
	myPc := Computadora{Cpu: "Intel i5", Ram: 32, TarjetaVideo: "RTX 4090 TI"}
	fmt.Println(myPc)
	myPc.ping()
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
}
