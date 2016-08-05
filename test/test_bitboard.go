package test

import (
	"fmt"
	"go-chess/constants"
	"go-chess/utils"
)

// BitboardTesting tests bitboard structure of pieces
func BitboardTesting() {
	fmt.Println("\nCounting and popping of pawns:")
	var playBitBoard utils.Bitboard
	playBitBoard = 0
	playBitBoard |= (utils.Bitboard(1) << utils.Bitboard(utils.SQ64(constants.D2)))
	playBitBoard |= (utils.Bitboard(1) << utils.Bitboard(utils.SQ64(constants.D3)))
	playBitBoard |= (utils.Bitboard(1) << utils.Bitboard(utils.SQ64(constants.D4)))
	playBitBoard.Print()
	fmt.Println("\nCount: ", playBitBoard.Count())
	i := (&playBitBoard).Pop()
	fmt.Println("Index: ", i)
	playBitBoard.Print()
	fmt.Println("\nCount: ", playBitBoard.Count())
	i = (&playBitBoard).Pop()
	fmt.Println("Index: ", i)
	playBitBoard.Print()
	fmt.Println("\nCount: ", playBitBoard.Count())
	i = (&playBitBoard).Pop()
	fmt.Println("Index: ", i)
	playBitBoard.Print()
	fmt.Println("\nCount: ", playBitBoard.Count())

	fmt.Println("Setting and clearing of bits on the chess board:")
	playBitBoard = 0
	(&playBitBoard).Set(61)
	playBitBoard.Print()
	(&playBitBoard).Clear(61)
	playBitBoard.Print()
}
