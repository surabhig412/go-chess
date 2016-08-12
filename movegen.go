package main

import (
	"errors"
	"fmt"
)

// SMove defines the structure of each move
type SMove struct {
	move  int
	score int
}

// SMoveList is the list of all moves
type SMoveList struct {
	moves [MaxPositionMoves]SMove
	count int
}

// Print prints the movelist
func (list *SMoveList) Print() {
	fmt.Println("MoveList: ", list.count)
	for index := 0; index < list.count; index++ {
		move := list.moves[index].move
		score := list.moves[index].score
		fmt.Printf("Move: %d > %s (score: %d)\n", index+1, PrintMove(move), score)
	}
	fmt.Printf("MoveList Total %d moves:\n", list.count)
}

// AddQuietMove to move list
func (list *SMoveList) AddQuietMove(pos *SBoard, move int) {
	list.moves[list.count].move = move
	list.moves[list.count].score = 0
	list.count++
}

// AddCaptureMove to move list
func (list *SMoveList) AddCaptureMove(pos *SBoard, move int) {
	list.moves[list.count].move = move
	list.moves[list.count].score = 0
	list.count++
}

// AddEnPassantMove to move list
func (list *SMoveList) AddEnPassantMove(pos *SBoard, move int) {
	list.moves[list.count].move = move
	list.moves[list.count].score = 0
	list.count++
}

// AddWhitePawnCaptureMove are possible capture moves of white pawn
func (list *SMoveList) AddWhitePawnCaptureMove(pos *SBoard, from, to, capture int) {
	if RanksBrd[from] == Rank7 {
		list.AddCaptureMove(pos, Move(from, to, capture, Wq, 0))
		list.AddCaptureMove(pos, Move(from, to, capture, Wr, 0))
		list.AddCaptureMove(pos, Move(from, to, capture, Wb, 0))
		list.AddCaptureMove(pos, Move(from, to, capture, Wn, 0))
	} else {
		list.AddCaptureMove(pos, Move(from, to, capture, Empty, 0))
	}
}

// AddWhitePawnMove are possible quiet moves of white pawn
func (list *SMoveList) AddWhitePawnMove(pos *SBoard, from, to int) {
	if RanksBrd[from] == Rank7 {
		list.AddQuietMove(pos, Move(from, to, Empty, Wq, 0))
		list.AddQuietMove(pos, Move(from, to, Empty, Wr, 0))
		list.AddQuietMove(pos, Move(from, to, Empty, Wb, 0))
		list.AddQuietMove(pos, Move(from, to, Empty, Wn, 0))
	} else {
		list.AddQuietMove(pos, Move(from, to, Empty, Empty, 0))
	}
}

// AddBlackPawnCaptureMove are possible capture moves of black pawn
func (list *SMoveList) AddBlackPawnCaptureMove(pos *SBoard, from, to, capture int) {
	if RanksBrd[from] == Rank2 {
		list.AddCaptureMove(pos, Move(from, to, capture, Bq, 0))
		list.AddCaptureMove(pos, Move(from, to, capture, Br, 0))
		list.AddCaptureMove(pos, Move(from, to, capture, Bb, 0))
		list.AddCaptureMove(pos, Move(from, to, capture, Bn, 0))
	} else {
		list.AddCaptureMove(pos, Move(from, to, capture, Empty, 0))
	}
}

// AddBlackPawnMove are possible quiet moves of black pawn
func (list *SMoveList) AddBlackPawnMove(pos *SBoard, from, to int) {
	if RanksBrd[from] == Rank2 {
		list.AddQuietMove(pos, Move(from, to, Empty, Bq, 0))
		list.AddQuietMove(pos, Move(from, to, Empty, Br, 0))
		list.AddQuietMove(pos, Move(from, to, Empty, Bb, 0))
		list.AddQuietMove(pos, Move(from, to, Empty, Bn, 0))
	} else {
		list.AddQuietMove(pos, Move(from, to, Empty, Empty, 0))
	}
}

// GenerateAllMoves will generate all possible moves of board
func (list *SMoveList) GenerateAllMoves(pos *SBoard) error {
	err := pos.Check()
	if err != nil {
		return err
	}
	list.count = 0
	if pos.Side == White {
		for pieceNum := 0; pieceNum < pos.PceNum[Wp]; pieceNum++ {
			sq := pos.PList[Wp][pieceNum]
			if !SqOnBoard(sq) {
				return errors.New("Square is not on board")
			}
			if pos.Pieces[sq+10] == Empty {
				list.AddWhitePawnMove(pos, sq, sq+10)
				if RanksBrd[sq] == Rank2 && pos.Pieces[sq+20] == Empty {
					list.AddQuietMove(pos, Move(sq, sq+20, Empty, Empty, MFlagPS))
				}
			}
			if SqOnBoard(sq+9) && PieceCol[pos.Pieces[sq+9]] == Black {
				list.AddWhitePawnCaptureMove(pos, sq, sq+9, pos.Pieces[sq+9])
			}
			if SqOnBoard(sq+11) && PieceCol[pos.Pieces[sq+11]] == Black {
				list.AddWhitePawnCaptureMove(pos, sq, sq+11, pos.Pieces[sq+11])
			}
			if sq+9 == pos.EnPas {
				list.AddCaptureMove(pos, Move(sq, sq+9, Empty, Empty, MFlagEP))
			}
			if sq+11 == pos.EnPas {
				list.AddCaptureMove(pos, Move(sq, sq+11, Empty, Empty, MFlagEP))
			}
		}
	} else {
		for pieceNum := 0; pieceNum < pos.PceNum[Bp]; pieceNum++ {
			sq := pos.PList[Bp][pieceNum]
			if !SqOnBoard(sq) {
				return errors.New("Square is not on board")
			}
			if pos.Pieces[sq-10] == Empty {
				list.AddBlackPawnMove(pos, sq, sq-10)
				if RanksBrd[sq] == Rank7 && pos.Pieces[sq-20] == Empty {
					list.AddQuietMove(pos, Move(sq, sq-20, Empty, Empty, MFlagPS))
				}
			}
			if SqOnBoard(sq-9) && PieceCol[pos.Pieces[sq-9]] == White {
				list.AddBlackPawnCaptureMove(pos, sq, sq-9, pos.Pieces[sq-9])
			}
			if SqOnBoard(sq-11) && PieceCol[pos.Pieces[sq-11]] == White {
				list.AddBlackPawnCaptureMove(pos, sq, sq-11, pos.Pieces[sq-11])
			}
			if sq-9 == pos.EnPas {
				list.AddCaptureMove(pos, Move(sq, sq-9, Empty, Empty, MFlagEP))
			}
			if sq-11 == pos.EnPas {
				list.AddCaptureMove(pos, Move(sq, sq-11, Empty, Empty, MFlagEP))
			}
		}
	}
	return nil
}
