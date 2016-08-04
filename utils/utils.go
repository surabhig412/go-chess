package utils

import (
	"go-chess/constants"
	"math/rand"
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

// Rand64 creates a random 64 bit uint value
func Rand64() uint64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return (uint64(r.Int63()) + uint64((0|1)<<63))
}
