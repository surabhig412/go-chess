package main

import (
	"fmt"
	"go-chess/test"
)

// Test tests functions of go-chess
func Test() {
	fmt.Println("\nTesting Boards: ")
	test.BoardTesting()
	fmt.Println("\nTesting Bitboards: ")
	test.BitboardTesting()
	fmt.Println("\nTesting Parsing of FEN notation: ")
	test.ParseFENTesting()
}
