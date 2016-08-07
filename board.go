package main

import "fmt"

// SBoard is board structure
type SBoard struct {
	Pieces     [BrdSqNum]int //all pieces on 120 board
	Pawns      [3]uint64     // 64 bit structure of white, black and both pawns
	KingSq     [2]int        // Positions of white, black and both kings
	Side       int           // White or black side playing
	EnPas      int           // enPas rule
	FiftyMove  int           // Number of move according to the fifty rule
	CastlePerm int           // castling Permission
	Ply        int           // Number of half moves played
	HisPly     int           // History of ply- highest number of moves played in game
	PosKey     uint64        // value of hashing piece, side, enPas and castle values
	PceNum     [13]int       // Total number of respective pieces on the board
	BigPce     [2]int        // index 0 represents number of big pieces on white side and index 1 of black side
	MajPce     [2]int        // index 0 represents number of major pieces on white side and index 1 of black side
	MinPce     [2]int        // index 0 represents number of minor pieces on white side and index 1 of black side
	Material   [2]int        // index 0 represents total value of all pieces on white side and index 1 of black side
	History    [MaxGameMoves]SUndo
	PList      [13][10]int //piece list
}

// Print prints the entire chess board
func (pos *SBoard) Print() {

	fmt.Println("\nGame Board:")

	for rank := Rank8; rank >= Rank1; rank-- {
		fmt.Printf("%d  ", rank+1)
		for file := FileA; file <= FileH; file++ {
			sq := FR2SQ(file, rank)
			piece := pos.Pieces[sq]
			fmt.Printf("%3c", PceChar[piece])
		}
		fmt.Println()
	}

	fmt.Printf("\n   ")
	for file := FileA; file <= FileH; file++ {
		fmt.Printf("%3c", rune(97+file))
	}
	fmt.Println()
	fmt.Printf("side:%c\n", SideChar[pos.Side])
	fmt.Printf("enPas:%d\n", pos.EnPas)
	castling := ""
	if (pos.CastlePerm & Wkca) == Wkca {
		castling += "K"
	} else {
		castling += "-"
	}
	if (pos.CastlePerm & Wqca) == Wqca {
		castling += "Q"
	} else {
		castling += "-"
	}
	if (pos.CastlePerm & Bkca) == Bkca {
		castling += "k"
	} else {
		castling += "-"
	}
	if (pos.CastlePerm & Bqca) == Bqca {
		castling += "q"
	} else {
		castling += "-"
	}
	fmt.Printf("castle:%s\n", castling)
	fmt.Printf("PosKey:%x\n", pos.PosKey)
}

// Reset resets the chess board
func (pos *SBoard) Reset() {
	for index := 0; index < BrdSqNum; index++ {
		pos.Pieces[index] = Offboard
	}
	for index := 0; index < 64; index++ {
		pos.Pieces[SQ120(index)] = Empty
	}
	for index := 0; index < 2; index++ {
		pos.BigPce[index] = 0
		pos.MajPce[index] = 0
		pos.MinPce[index] = 0
		pos.Pawns[index] = uint64(0)
	}
	for index := 0; index < 13; index++ {
		pos.PceNum[index] = 0
	}
	pos.KingSq[White] = NoSq
	pos.KingSq[Black] = NoSq
	pos.Side = Both
	pos.EnPas = NoSq
	pos.FiftyMove = 0
	pos.Ply = 0
	pos.HisPly = 0
	pos.CastlePerm = 0
	pos.PosKey = uint64(0)
}

// UpdateListMaterial updates various arrays based on the piece on board
func (pos *SBoard) UpdateListMaterial() {
	for index := 0; index < BrdSqNum; index++ {
		piece := pos.Pieces[index]
		if piece != Offboard && piece != Empty {
			colour := PieceCol[piece]

			// Setting piece types
			if PieceBig[piece] == True {
				pos.BigPce[colour]++
			}
			if PieceMaj[piece] == True {
				pos.MajPce[colour]++
			}
			if PieceMin[piece] == True {
				pos.MinPce[colour]++
			}

			// Setting material of each side
			pos.Material[colour] += PieceVal[piece]

			// Setting piece list
			pos.PList[piece][pos.PceNum[piece]] = index
			pos.PceNum[piece]++

			// Setting King squares of each side
			if piece == Wk || piece == Bk {
				pos.KingSq[colour] = index
			}
		}
	}
}
