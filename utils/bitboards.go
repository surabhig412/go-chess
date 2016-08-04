package utils

import (
	"fmt"
	"go-chess/constants"
	"strconv"
	"strings"
)

// Uint64Utils is the number whose 64 bit positions is to be printed
type Uint64Utils uint64

// PrintBitBoard is to represent presence of each piece using bitwise operations
func (bit Uint64Utils) PrintBitBoard() {
	var shiftMe uint64
	shiftMe = 1
	fmt.Println()
	for rank := constants.Rank8; rank >= constants.Rank1; rank-- {
		for file := constants.FileA; file <= constants.FileH; file++ {
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

// PopBit returns position of first piece from MSB side of a particular pieces' bitboard
func (bit *Uint64Utils) PopBit() int {
	bs := strconv.FormatUint(uint64(*bit), 2)
	rev := reverse(bs)
	i := strings.Index(rev, "1")
	mask := ^(uint64(1) << uint64(i))
	*bit &= Uint64Utils(mask)
	return i
}

// CountBits counts number of pieces of a particular piece
func (bit Uint64Utils) CountBits() int {
	bs := strconv.FormatUint(uint64(bit), 2)
	return strings.Count(bs, "1")
}

// ClrBit clears the bit of a sq in a particular pieces' bitboard
func (bit *Uint64Utils) ClrBit(sq int) {
	*bit &= Uint64Utils(constants.ClearMask[sq])
}

// SetBit sets the bit of a sq in a particular pieces' bitboard
func (bit *Uint64Utils) SetBit(sq int) {
	*bit |= Uint64Utils(constants.SetMask[sq])
}
