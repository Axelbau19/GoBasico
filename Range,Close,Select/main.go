package main

import "fmt"

func messague(text string, c chan string) {
	c <- text
}

func main() {
	//make(chan tipoDato, numero de channels )
	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"
	fmt.Println(len(c), cap(c))

	//Range y  close
	//Va cerrar el canal y ya no va dejar insertar valor aunque tenga espacio
	close(c)
	for mesaje := range c {
		fmt.Println(mesaje)
	}
	//Select
	email := make(chan string)
	email2 := make(chan string)
	go messague("mensaje1", email)
	go messague("mensaje2", email2)
	//Tomar en cuenta cuanta cantidad tiene cada channel
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email:
			fmt.Println("El email se recibio de email1", m1)
		case m2 := <-email2:
			fmt.Println("El email se recibio de email2", m2)
		}
	}
}
