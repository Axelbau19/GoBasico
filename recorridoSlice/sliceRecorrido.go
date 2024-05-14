package main

import (
	"fmt"
	"strings"
)

// Funcion palidromos
func isPalidromo(word string) {
	var text string
	for i := len(word) - 1; i >= 0; i-- {
		text += string(word[i])
	}
	if strings.ToLower(word) == strings.ToLower(text) {
		fmt.Println("Es palidromo")
	} else {
		fmt.Println("No lo es")
	}
}

func main() {
	slice := []string{"hola", "como", "estas"}
	for _, valor := range slice {
		fmt.Println(valor)
	}
	isPalidromo("ama")
	isPalidromo("amor a roma")
	isPalidromo("Ama")
	isPalidromo("Amor a Roma")
}
