package main

import "fmt"

func main() {
	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))
	//Slice no tiene limitacion de datos a guardar en el arreglo
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))
	//Slicing sirve para interactuar
	//Metodos en el slice
	fmt.Println(slice[0])
	//hasta indice 3
	fmt.Println(slice[:3])
	//desde indice 2 hasta 4
	fmt.Println(slice[2:4])
	//Desde indice 4 hasta final
	fmt.Println(slice[4:])
	//Append
	slice = append(slice, 7)
	fmt.Println(slice)
	//Agregar otra lista
	//Append nueva lista
	newSlice := []int{8, 9, 10}
	//Lo que va esos ... los van a descomprimir de forma indepediente
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
