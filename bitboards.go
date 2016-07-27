package main

import "fmt"

// PrintBitBoard is to represent presence of pieces(particularly pawns) using bitwise operations
func PrintBitBoard(bit U64) {
	var shiftMe U64
	shiftMe = 1
	fmt.Println()
	for rank := Rank8; rank >= Rank1; rank-- {
		for file := FileA; file <= FileH; file++ {
			sq := FR2SQ(file, rank) //120 based
			sq64 := SQ64(sq)        //64 based
			expr := (shiftMe << U64(sq64)) & bit
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
