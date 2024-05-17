package videogame

import "fmt"

type VideoGame struct {
	Name  string
	Price float64
}

func (videoGame VideoGame) String() string {
	return fmt.Sprintf("El juego es %s y tiene un costo de: $ %0.2f USD", videoGame.Name, videoGame.Price)
}
