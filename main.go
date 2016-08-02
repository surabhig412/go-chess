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
	err := engine.ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq a4", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Board:")
	for j := 0; j < 120; j++ {
		fmt.Printf("%d ", board.Pieces[j])
	}
}
