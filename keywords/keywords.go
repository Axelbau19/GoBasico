package main

import "fmt"

func main() {
	//Defer cuando acabe una tarea , enseguida va esta(es decir hasta el ultimo) por ejemplo:
	defer fmt.Println("Hola")
	fmt.Println("Mundo")
	//Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//Continue se utiliza en ciclos para obtener un valor y quiere que siga el ciclo
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		//Break para terminar el ciclo
		if i == 8 {
			fmt.Println("Aqui para")
			break
		}
	}
}
