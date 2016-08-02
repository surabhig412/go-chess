package main

import (
	"fmt"
	"go-chess/engine"
	"go-chess/models"
)

func main() {
	fmt.Println()
	AllInit()

	var board models.SBoard
	i := engine.ParseFEN("4KQB", &board)
	fmt.Println("Parsing successful: ", i)
	fmt.Println("Board:")
	for j := 0; j < 120; j++ {
		fmt.Printf("%d ", board.Pieces[j])
	}
}
