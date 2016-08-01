package main

// initSq120To64 initializes 64 and 120 sq arrays
func initSq120To64() {
	sq := A1
	sq64 := 0
	for index := 0; index < BrdSqNum; index++ {
		Sq120ToSq64[index] = 65
	}

	for index := 0; index < 64; index++ {
		Sq64ToSq120[index] = 120
	}

	for rank := Rank1; rank <= Rank8; rank++ {
		for file := FileA; file <= FileH; file++ {
			sq = FR2SQ(file, rank)
			Sq64ToSq120[sq64] = sq
			Sq120ToSq64[sq] = sq64
			sq64++
		}
	}
}

// initBitMasks initializes SetMask and ClearMask arrays
func initBitMasks() {
	for index := 0; index < 64; index++ {
		SetMask[index] = U64(0)
		ClearMask[index] = U64(0)
	}
	for index := 0; index < 64; index++ {
		SetMask[index] |= (U64(1) << U64(index))
		ClearMask[index] = ^(SetMask[index])
	}
}

// initHashKeys sets all possible positions of each piece, side to play and castling keys with random number
func initHashKeys() {
	for index := 0; index < 13; index++ {
		for index2 := 0; index2 < 120; index2++ {
			PieceKeys[index][index2] = Rand64()
		}
	}
	SideKey = Rand64()
	for index := 0; index < 16; index++ {
		CastleKeys[index] = Rand64()
	}
}

// AllInit is used to initialize arrays, masks and keys of the board
func AllInit() {
	initSq120To64()
	initBitMasks()
	initHashKeys()
}
