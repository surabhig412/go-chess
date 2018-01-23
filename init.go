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
		SetMask[index] = uint64(0)
		ClearMask[index] = uint64(0)
	}
	for index := 0; index < 64; index++ {
		SetMask[index] |= (uint64(1) << uint64(index))
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

// initFilesRankBrd initializes files and ranks arrays with their respective file and ranks
func initFilesRankBrd() {
	for index := 0; index < BrdSqNum; index++ {
		FilesBrd[index] = Offboard
		RanksBrd[index] = Offboard
	}

	for rank := Rank1; rank <= Rank8; rank++ {
		for file := FileA; file <= FileH; file++ {
			sq := FR2SQ(file, rank)
			FilesBrd[sq] = file
			RanksBrd[sq] = rank
		}
	}
}

func initEvalMasks() {
	for rank := Rank8; rank >= Rank1; rank-- {
		for file := FileA; file <= FileH; file++ {
			sq := rank*8 + file
			FileBBMask[file] |= (uint64(1) << uint64(sq))
			RankBBMask[rank] |= (uint64(1) << uint64(sq))
		}
	}
}

// initCastlePermission initializes CastlePerm array with castling permission values
func initCastlePermission() {
	for index := 0; index < BrdSqNum; index++ {
		CastlePerm[index] = 15
	}
	CastlePerm[A1] = 13 // disables Wqca
	CastlePerm[E1] = 12 // disables Wqca and Wkca
	CastlePerm[H1] = 14 // disables Wkca
	CastlePerm[A8] = 7  // disables Bqca
	CastlePerm[E8] = 3  // disables Bqca and Bkca
	CastlePerm[H8] = 11 // disables Bkca
}

// initMvvLva initializes the score of each piece according to most valuable victim and least valuable attacker
func initMvvLva() {
	for attacker := Wp; attacker <= Bk; attacker++ {
		for victim := Wp; victim <= Bk; victim++ {
			MvvLvaScores[victim][attacker] = VictimScore[victim] + 6 - (VictimScore[attacker] / 100)
		}
	}
}

// AllInit is used to initialize arrays, masks and keys of the board
func AllInit() {
	initSq120To64()
	initBitMasks()
	initHashKeys()
	initFilesRankBrd()
	initEvalMasks()
	initCastlePermission()
	initMvvLva()
}
