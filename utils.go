package main

import (
	"math/rand"
	"time"
)

// FR2SQ is used to represent file rank to a particular square in 120-sq board
func FR2SQ(f, r int) int {
	return (21 + f + (r * 10))
}

// SQ64 returns 64-square equivalent of 120-square board
func SQ64(sq120 int) int {
	return Sq120ToSq64[sq120]
}

// SQ120 returns 120-square equivalent of 64-square board
func SQ120(sq64 int) int {
	return Sq64ToSq120[sq64]
}

// Rand64 creates a random 64 bit uint value
func Rand64() uint64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return (uint64(r.Int63()) | uint64((0|1)<<63))
}

// IsBQ returns if the piece is bishop or queen or not
func IsBQ(piece int) int {
	return PieceBishopQueen[piece]
}

// IsRQ returns if the piece is rook or queen or not
func IsRQ(piece int) int {
	return PieceRookQueen[piece]
}

// IsKn returns if the piece is knight or not
func IsKn(piece int) int {
	return PieceKnight[piece]
}

// IsKi returns if the piece is king or not
func IsKi(piece int) int {
	return PieceKing[piece]
}

// FromSq returns value of from-square where move has been done
func FromSq(num int) int {
	return (num & 0x7F)
}

// ToSq returns value of to-square where move has been done
func ToSq(num int) int {
	return ((num >> 7) & 0x7F)
}

// Captured returns value of pieces captured in the move
func Captured(num int) int {
	return ((num >> 14) & 0xF)
}

// Promoted returns value of pieces promoted in the move
func Promoted(num int) int {
	return ((num >> 20) & 0xF)
}

// Move creates move with all possible parameters
func Move(from, to, capture, promoted, flag int) int {
	return (from | (to << 7) | (capture << 14) | (promoted << 20) | flag)
}
