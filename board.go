package main

import (
	"errors"
	"fmt"
)

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
		pos.Material[index] = 0
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

	for index := 0; index < 3; index++ {
		pos.Pawns[index] = uint64(0)
	}
}

// UpdateListsMaterial updates various arrays based on the piece on board
func (pos *SBoard) UpdateListsMaterial() {
	for index := 0; index < BrdSqNum; index++ {
		piece := pos.Pieces[index]
		if piece != Offboard && piece != Empty {
			colour := PieceCol[piece]

			// Updating piece types
			if PieceBig[piece] == True {
				pos.BigPce[colour]++
			}
			if PieceMaj[piece] == True {
				pos.MajPce[colour]++
			}
			if PieceMin[piece] == True {
				pos.MinPce[colour]++
			}

			// Updating material of each side
			pos.Material[colour] += PieceVal[piece]

			// Updating piece list
			pos.PList[piece][pos.PceNum[piece]] = index
			pos.PceNum[piece]++

			// Updating King squares of each side
			if piece == Wk || piece == Bk {
				pos.KingSq[colour] = index
			}

			// Updating pawns arrays as per the piece
			if piece == Wp || piece == Bp {
				pawnStructure := Bitboard(pos.Pawns[colour])
				(&pawnStructure).Set(SQ64(index))
				pos.Pawns[colour] = uint64(pawnStructure)
				pawnStructure = Bitboard(pos.Pawns[Both])
				(&pawnStructure).Set(SQ64(index))
				pos.Pawns[Both] = uint64(pawnStructure)
			}
		}
	}
}

// Check cross-checks if all pieces are placed properly
func (pos *SBoard) Check() error {
	var tempPceNumArr = [13]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var tempBigPceArr = [2]int{0, 0}
	var tempMajPceArr = [2]int{0, 0}
	var tempMinPceArr = [2]int{0, 0}
	var tempMaterialArr = [2]int{0, 0}
	var tempPawns = [3]uint64{pos.Pawns[White], pos.Pawns[Black], pos.Pawns[Both]}

	// check piece lists
	for tempPiece := Wp; tempPiece <= Bk; tempPiece++ {
		for tempPieceNum := 0; tempPieceNum < pos.PceNum[tempPiece]; tempPieceNum++ {
			sq120 := pos.PList[tempPiece][tempPieceNum]
			if pos.Pieces[sq120] != tempPiece {
				return errors.New("Piece List mismatch")
			}
		}
	}

	// check piece count and minor, major, big and material counters
	for sq64 := 0; sq64 < 64; sq64++ {
		sq120 := SQ120(sq64)
		tempPiece := pos.Pieces[sq120]
		tempPceNumArr[tempPiece]++
		colour := PieceCol[tempPiece]
		if colour == Black || colour == White {
			if PieceBig[tempPiece] == True {
				tempBigPceArr[colour]++
			}
			if PieceMaj[tempPiece] == True {
				tempMajPceArr[colour]++
			}
			if PieceMin[tempPiece] == True {
				tempMinPceArr[colour]++
			}
			tempMaterialArr[colour] += PieceVal[tempPiece]
		}
	}

	for tempPiece := Wp; tempPiece <= Bk; tempPiece++ {
		if tempPceNumArr[tempPiece] != pos.PceNum[tempPiece] {
			return errors.New("Piece count mismatch")
		}
	}
	for pieceColor := White; pieceColor <= Black; pieceColor++ {
		if tempBigPceArr[pieceColor] != pos.BigPce[pieceColor] {
			return errors.New("BigPce mismatch")
		}
		if tempMajPceArr[pieceColor] != pos.MajPce[pieceColor] {
			return errors.New("MajPce mismatch")
		}
		if tempMinPceArr[pieceColor] != pos.MinPce[pieceColor] {
			return errors.New("MinPce mismatch")
		}
		if tempMaterialArr[pieceColor] != pos.Material[pieceColor] {
			return errors.New("Material mismatch")
		}
	}

	// check bitboards count
	pCount := (Bitboard(tempPawns[White])).Count()
	if pCount != pos.PceNum[Wp] {
		return errors.New("White pawn bitboard mismatch")
	}
	pCount = (Bitboard(tempPawns[Black])).Count()
	if pCount != pos.PceNum[Bp] {
		return errors.New("Black pawn bitboard mismatch")
	}
	pCount = (Bitboard(tempPawns[Both])).Count()
	if pCount != (pos.PceNum[Wp] + pos.PceNum[Bp]) {
		return errors.New("Both pawns bitboard mismatch")
	}

	// check bitboards squares
	for tempPawns[White] != uint64(0) {
		whiteBitboard := Bitboard(tempPawns[White])
		sq64 := (&whiteBitboard).Pop()
		tempPawns[White] = uint64(whiteBitboard)
		if pos.Pieces[SQ120(sq64)] != Wp {
			return errors.New("White pawn bitboard mapping to square mismatch")
		}
	}
	for tempPawns[Black] != uint64(0) {
		blackBitboard := Bitboard(tempPawns[Black])
		sq64 := (&blackBitboard).Pop()
		tempPawns[Black] = uint64(blackBitboard)
		if pos.Pieces[SQ120(sq64)] != Bp {
			return errors.New("Black pawn bitboard mapping to square mismatch")
		}
	}
	for tempPawns[Both] != uint64(0) {
		bothColourBitboard := Bitboard(tempPawns[Both])
		sq64 := (&bothColourBitboard).Pop()
		tempPawns[Both] = uint64(bothColourBitboard)
		if !((pos.Pieces[SQ120(sq64)] == Wp) || (pos.Pieces[SQ120(sq64)] == Bp)) {
			return errors.New("Both pawns bitboard mapping to square mismatch")
		}
	}

	// check side, PosKey, enPas, king squares and castle permissions
	if !((pos.Side == White) || (pos.Side == Black)) {
		return errors.New("Side mismatch")
	}
	if GeneratePosKey(pos) != pos.PosKey {
		return errors.New("PosKey mismatch")
	}
	if !((pos.EnPas == NoSq) || (RanksBrd[pos.EnPas] == Rank6 && pos.Side == White) || (RanksBrd[pos.EnPas] == Rank3 && pos.Side == Black)) {
		return errors.New("EnPas rule mismatch")
	}
	if pos.Pieces[pos.KingSq[White]] != Wk {
		return errors.New("White king square mismatch")
	}
	if pos.Pieces[pos.KingSq[Black]] != Bk {
		return errors.New("Black king square mismatch")
	}

	return nil
}
