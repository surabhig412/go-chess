package main

import (
	"fmt"
	"strconv"
	"strings"
)

// FR2SQ is used to represent file rank to a particular square in 120-sq board
func FR2SQ(f, r int) int {
	return (21 + f + (r * 10))
}

// SQ64 returns 64-square equivalent of 120-square board
func SQ64(sq120 int) int {
	return Sq120ToSq64[sq120]
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// PopBit returns position of first pawn from MSB side
func PopBit(n U64) (int, U64) {
	bs := strconv.FormatUint(uint64(n), 2)
	rev := reverse(bs)
	i := strings.Index(rev, "1")
	fmt.Println("Index: ", i)
	mask := ^(U64(1) << U64(i))
	n &= mask
	return i, n
}

// CountBits counts number of pawns on the board
func CountBits(n U64) int {
	bs := strconv.FormatUint(uint64(n), 2)
	return strings.Count(bs, "1")
}
