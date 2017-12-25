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
	(&board).PvTable.Init()
	fmt.Printf("Please enter a move > ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		switch command {
		case "q":
			os.Exit(1)
		case "t":
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
		case "p":
			PerftTest(3, &board)
			fmt.Printf("Please enter a move > ")
		case "pv":
			count, _ := GetPvLine(4, &board)
			for i := 0; i < count; i++ {
				move := board.PvArray[i]
				fmt.Println(PrintMove(move))
			}
			fmt.Printf("Please enter a move > ")
		default:
			move, err := ParseMove(strings.ToLower(command), &board)
			if err != nil {
				fmt.Errorf("Error in parsing move", err)
			}
			if move != NoMove {
				StorePvMove(&board, move)
				_, err = MakeMove(move, &board)
				if err != nil {
					fmt.Errorf("Error in making move", err)
				}
			} else {
				fmt.Println("Move not parsed")
			}
			board.Print()
			fmt.Printf("Please enter a move > ")
		}

	}
}
