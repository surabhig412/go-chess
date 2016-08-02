package utils

import (
	"go-chess/constants"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// FR2SQ is used to represent file rank to a particular square in 120-sq board
func FR2SQ(f, r int) int {
	return (21 + f + (r * 10))
}

// SQ64 returns 64-square equivalent of 120-square board
func SQ64(sq120 int) int {
	return constants.Sq120ToSq64[sq120]
}

// SQ120 returns 120-square equivalent of 64-square board
func SQ120(sq64 int) int {
	return constants.Sq64ToSq120[sq64]
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// PopBit returns position of first pawn from MSB side
func PopBit(n *uint64) int {
	bs := strconv.FormatUint(uint64(*n), 2)
	rev := reverse(bs)
	i := strings.Index(rev, "1")
	mask := ^(uint64(1) << uint64(i))
	*n &= mask
	return i
}

// CountBits counts number of pawns on the board
func CountBits(n uint64) int {
	bs := strconv.FormatUint(uint64(n), 2)
	return strings.Count(bs, "1")
}

// ClrBit clears the bit of a sq
func ClrBit(bit uint64, sq int) uint64 {
	bit &= constants.ClearMask[sq]
	return bit
}

// SetBit sets the bit of a sq
func SetBit(bit uint64, sq int) uint64 {
	bit |= constants.SetMask[sq]
	return bit
}

// Rand64 creates a random 64 bit uint value
func Rand64() uint64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return (uint64(r.Int63()) + uint64((0|1)<<63))
}
