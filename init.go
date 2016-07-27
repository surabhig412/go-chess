package main

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

// AllInit is used to initialize 120 and 64 square arrays
func AllInit() {
	initSq120To64()
}
