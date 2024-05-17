package mypackage

import "fmt"

type CarPublic struct {
	Name string
	Year int
}

type carPrivate struct {
	name string
	year int
}

//Funcion publica
func PrintMessage(text string) {
	fmt.Println(text)
}
func printMessage(text string) {
	fmt.Println(text)
}
