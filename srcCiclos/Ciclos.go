package main

import "fmt"

func main() {
	//For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Print("\n")
	//For while
	counter := 0
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}
	//For forever
	/*
		counterForever:=0
		for{
			fmt.print(counterForever)
			counterForever++
		}
	*/
	//Reto
	for j := 20; j >= 0; j-- {
		fmt.Println(j)
	}

}
