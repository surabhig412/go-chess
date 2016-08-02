package utils

import (
	"fmt"
	"go-chess/constants"
)

// PrintBitBoard is to represent presence of pieces(particularly pawns) using bitwise operations
func PrintBitBoard(bit uint64) {
	var shiftMe uint64
	shiftMe = 1
	fmt.Println()
	for rank := constants.Rank8; rank >= constants.Rank1; rank-- {
		for file := constants.FileA; file <= constants.FileH; file++ {
			sq := FR2SQ(file, rank) //120 based
			sq64 := SQ64(sq)        //64 based
			expr := (shiftMe << uint64(sq64)) & bit
			// fmt.Println(expr)
			if expr != 0 {
				fmt.Printf("X")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
