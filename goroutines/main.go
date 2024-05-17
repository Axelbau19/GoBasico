package main

import (
	"fmt"
	"sync"
	"time"
)

func returnText(text string, wg *sync.WaitGroup) {
	//Espera la ultima instruccion
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	//Aprender mas de go routine
	//Crear un wait group
	var wg sync.WaitGroup

	fmt.Println("xd")
	//se agrega
	wg.Add(1)

	//Aqui no ejecuta la siguiente instruccion
	go returnText("Ni idea", &wg)
	wg.Wait()
	//No es recomendable
	go func(text string) {
		fmt.Println(text)
	}("adios")
	time.Sleep(time.Second * 1)

}
