package main

import "fmt"

func sayText(text string, c chan<- string) {
	c <- text
}
func main() {
	//Declarar el channel
	c := make(chan string, 1)
	fmt.Println("Hola")
	go sayText("Bye", c)
	fmt.Println(<-c)
}
