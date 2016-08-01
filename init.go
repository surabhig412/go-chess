package main

import (
	"go-chess/constants"
	"go-chess/utils"
)

// initSq120To64 initializes 64 and 120 sq arrays
func initSq120To64() {
	sq := constants.A1
	sq64 := 0
	for index := 0; index < constants.BrdSqNum; index++ {
		constants.Sq120ToSq64[index] = 65
	}

	for index := 0; index < 64; index++ {
		constants.Sq64ToSq120[index] = 120
	}

	for rank := constants.Rank1; rank <= constants.Rank8; rank++ {
		for file := constants.FileA; file <= constants.FileH; file++ {
			sq = utils.FR2SQ(file, rank)
			constants.Sq64ToSq120[sq64] = sq
			constants.Sq120ToSq64[sq] = sq64
			sq64++
		}
	}
}

// initBitMasks initializes SetMask and ClearMask arrays
func initBitMasks() {
	for index := 0; index < 64; index++ {
		constants.SetMask[index] = constants.U64(0)
		constants.ClearMask[index] = constants.U64(0)
	}
	for index := 0; index < 64; index++ {
		constants.SetMask[index] |= (constants.U64(1) << constants.U64(index))
		constants.ClearMask[index] = ^(constants.SetMask[index])
	}
}

// initHashKeys sets all possible positions of each piece, side to play and castling keys with random number
func initHashKeys() {
	for index := 0; index < 13; index++ {
		for index2 := 0; index2 < 120; index2++ {
			constants.PieceKeys[index][index2] = utils.Rand64()
		}
	}
	constants.SideKey = utils.Rand64()
	for index := 0; index < 16; index++ {
		constants.CastleKeys[index] = utils.Rand64()
	}
}

// AllInit is used to initialize arrays, masks and keys of the board
func AllInit() {
	initSq120To64()
	initBitMasks()
	initHashKeys()
}
