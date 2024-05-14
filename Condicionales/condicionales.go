package main

import "fmt"

func isParImpar(number int) {
	if number%2 == 0 {
		fmt.Println("El numero", number, " es par")
	} else {
		fmt.Println("Es numero ", number, "es impar")
	}
}
func verificAccount(username, password string) {
	if username == "Axel" && password == "gordito21" {
		fmt.Println("Acceso Permitido")
	} else {
		fmt.Println("Acceso Denegado")
	}
}

func main() {
	isParImpar(8)
	isParImpar(3)
	verificAccount("Axel", "gordito21")
	verificAccount("Axel", "gorditxd")
	verificAccount("Alexis", "nosexd")
}
