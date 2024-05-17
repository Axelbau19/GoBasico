package main

import (
	"fmt"
	pk "learngo/src/mypackage"
)

func main() {
	var tsurito pk.CarPublic
	tsurito.Name = "Tsurito"
	tsurito.Year = 1994
	fmt.Println(tsurito)
	pk.PrintMessage("RUN!!!!!!")
	//Es privada
	//pk.printMessage("RUN!!!!!!!!")
}
