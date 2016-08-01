package main

import (
	"fmt"
	"go-chess/constants"
	"go-chess/utils"
)

func main() {
	fmt.Println()
	AllInit()
	var playBitBoard constants.U64
	fmt.Println("Random number: ", utils.Rand64())

	fmt.Println("Board structures:")
	for index := 0; index < constants.BrdSqNum; index++ {
		if index%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", constants.Sq120ToSq64[index])
	}
	fmt.Println()
	fmt.Println()
	for index := 0; index < 64; index++ {
		if index%8 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", constants.Sq64ToSq120[index])
	}

	fmt.Println("\nCounting and popping of pawns:")
	playBitBoard = 0
	playBitBoard |= (constants.U64(1) << constants.U64(utils.SQ64(constants.D2)))
	playBitBoard |= (constants.U64(1) << constants.U64(utils.SQ64(constants.D3)))
	playBitBoard |= (constants.U64(1) << constants.U64(utils.SQ64(constants.D4)))
	utils.PrintBitBoard(playBitBoard)
	fmt.Println("\nCount: ", utils.CountBits(playBitBoard))
	i := utils.PopBit(&playBitBoard)
	fmt.Println("Index: ", i)
	utils.PrintBitBoard(playBitBoard)
	fmt.Println("\nCount: ", utils.CountBits(playBitBoard))
	i = utils.PopBit(&playBitBoard)
	fmt.Println("Index: ", i)
	utils.PrintBitBoard(playBitBoard)
	fmt.Println("\nCount: ", utils.CountBits(playBitBoard))
	i = utils.PopBit(&playBitBoard)
	fmt.Println("Index: ", i)
	utils.PrintBitBoard(playBitBoard)
	fmt.Println("\nCount: ", utils.CountBits(playBitBoard))

	fmt.Println("Setting and clearing of bits on the chess board:")
	playBitBoard = 0
	playBitBoard = utils.SetBit(playBitBoard, 61)
	utils.PrintBitBoard(playBitBoard)
	playBitBoard = utils.ClrBit(playBitBoard, 61)
	utils.PrintBitBoard(playBitBoard)
}
