package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Bitboard is the number whose 64 bit positions is to be printed
type Bitboard uint64

// Print is to represent presence of each piece using bitwise operations
func (bit Bitboard) Print() {
	var shiftMe uint64
	shiftMe = 1
	fmt.Println()
	for rank := Rank8; rank >= Rank1; rank-- {
		for file := FileA; file <= FileH; file++ {
			sq := FR2SQ(file, rank) //120 based
			sq64 := SQ64(sq)        //64 based
			expr := (shiftMe << uint64(sq64)) & uint64(bit)
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

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// Pop returns position of first piece from MSB side of a particular pieces' bitboard
func (bit *Bitboard) Pop() (i int) {
	bs := strconv.FormatUint(uint64(*bit), 2)
	rev := reverse(bs)
	i = strings.Index(rev, "1")
	mask := ^(uint64(1) << uint64(i))
	*bit &= Bitboard(mask)
	return
}

// Count counts number of pieces of a particular piece
func (bit Bitboard) Count() int {
	bs := strconv.FormatUint(uint64(bit), 2)
	return strings.Count(bs, "1")
}

// Clear clears the bit of a sq in a particular pieces' bitboard
func (bit *Bitboard) Clear(sq int) {
	*bit &= Bitboard(ClearMask[sq])
}

// Set sets the bit of a sq in a particular pieces' bitboard
func (bit *Bitboard) Set(sq int) {
	*bit |= Bitboard(SetMask[sq])
}
