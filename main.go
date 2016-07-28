package main

import (
	"fmt"
)

// U64 is declared datatype
type U64 uint64

func main() {
	fmt.Println()
	AllInit()
	var playBitBoard U64
	playBitBoard = 0
	// for index := 0; index < 64; index++ {
	// 	fmt.Println("Index:", index)
	// 	PrintBitBoard(ClearMask[index])
	// 	fmt.Println()
	// }
	playBitBoard = SetBit(playBitBoard, 61)
	PrintBitBoard(playBitBoard)
	playBitBoard = ClrBit(playBitBoard, 61)
	PrintBitBoard(playBitBoard)
}
