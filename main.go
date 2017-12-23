package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println()
	AllInit()
	// Test()
	fmt.Println("Let us play chess!")
	var board SBoard
	err := ParseFEN(StartFEN, &board)
	if err != nil {
		fmt.Println("Error in parsing fen: ", err)
	}
	board.Print()
	fmt.Printf("Please enter a move > ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		if command == "q" {
			break
		} else if command == "t" {
			if board.Ply == 0 {
				fmt.Println("There are no moves made on the board")
			} else {
				err = TakeMove(&board)
				if err != nil {
					fmt.Errorf("Error in taking move back", err)
				}
				board.Print()
			}
			fmt.Printf("Please enter a move > ")
		} else {
			move, err := ParseMove(strings.ToLower(command), &board)
			if err != nil {
				fmt.Errorf("Error in parsing move", err)
			}
			if move != NoMove {
				_, err = MakeMove(move, &board)
				if err != nil {
					fmt.Errorf("Error in making move", err)
				}
			}
			board.Print()
			fmt.Printf("Please enter a move > ")
		}
	}
}
