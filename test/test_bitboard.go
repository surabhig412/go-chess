package test

import (
	"fmt"
	"go-chess/constants"
	"go-chess/utils"
)

// BitboardTesting tests bitboard structure of pieces
func BitboardTesting() {
	fmt.Println("\nCounting and popping of pawns:")
	var playBitBoard utils.Uint64Utils
	playBitBoard = 0
	playBitBoard |= (utils.Uint64Utils(1) << utils.Uint64Utils(utils.SQ64(constants.D2)))
	playBitBoard |= (utils.Uint64Utils(1) << utils.Uint64Utils(utils.SQ64(constants.D3)))
	playBitBoard |= (utils.Uint64Utils(1) << utils.Uint64Utils(utils.SQ64(constants.D4)))
	playBitBoard.PrintBitBoard()
	fmt.Println("\nCount: ", playBitBoard.CountBits())
	i := (&playBitBoard).PopBit()
	fmt.Println("Index: ", i)
	playBitBoard.PrintBitBoard()
	fmt.Println("\nCount: ", playBitBoard.CountBits())
	i = (&playBitBoard).PopBit()
	fmt.Println("Index: ", i)
	playBitBoard.PrintBitBoard()
	fmt.Println("\nCount: ", playBitBoard.CountBits())
	i = (&playBitBoard).PopBit()
	fmt.Println("Index: ", i)
	playBitBoard.PrintBitBoard()
	fmt.Println("\nCount: ", playBitBoard.CountBits())

	fmt.Println("Setting and clearing of bits on the chess board:")
	playBitBoard = 0
	(&playBitBoard).SetBit(61)
	playBitBoard.PrintBitBoard()
	(&playBitBoard).ClrBit(61)
	playBitBoard.PrintBitBoard()
}
