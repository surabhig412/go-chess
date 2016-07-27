package main

import (
	"fmt"
)

// U64 is declared datatype
type U64 uint64

func main() {
	fmt.Println()
	AllInit()
	// for index := 0; index < BrdSqNum; index++ {
	// 	if index%10 == 0 {
	// 		fmt.Println()
	// 	}
	// 	fmt.Printf("%5d", Sq120ToSq64[index])
	// }
	// fmt.Println()
	// fmt.Println()
	// for index := 0; index < 64; index++ {
	// 	if index%8 == 0 {
	// 		fmt.Println()
	// 	}
	// 	fmt.Printf("%5d", Sq64ToSq120[index])
	// }

	var playBitBoard U64
	playBitBoard = 0
	// PrintBitBoard(playBitBoard)

	playBitBoard |= (U64(1) << U64(SQ64(D2)))
	playBitBoard |= (U64(1) << U64(SQ64(D3)))
	playBitBoard |= (U64(1) << U64(SQ64(D4)))
	PrintBitBoard(playBitBoard)
	fmt.Println("\nCount: ", CountBits(playBitBoard))
	_, playBitBoard = PopBit(playBitBoard)
	PrintBitBoard(playBitBoard)
	fmt.Println("\nCount: ", CountBits(playBitBoard))
	_, playBitBoard = PopBit(playBitBoard)
	PrintBitBoard(playBitBoard)
	fmt.Println("\nCount: ", CountBits(playBitBoard))
	_, playBitBoard = PopBit(playBitBoard)
	PrintBitBoard(playBitBoard)
	fmt.Println("\nCount: ", CountBits(playBitBoard))
}
