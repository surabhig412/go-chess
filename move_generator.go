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

// addQuietMove to move list
func (list *SMoveList) addQuietMove(pos *SBoard, move int) {
	list.moves[list.count].move = move
	list.moves[list.count].score = 0
	list.count++
}

// addCaptureMove to move list
func (list *SMoveList) addCaptureMove(pos *SBoard, move int) {
	list.moves[list.count].move = move
	list.moves[list.count].score = 0
	list.count++
}

// addEnPassantMove to move list
func (list *SMoveList) addEnPassantMove(pos *SBoard, move int) {
	list.moves[list.count].move = move
	list.moves[list.count].score = 0
	list.count++
}

// addPawnCaptureMove are possible capture moves of pawn
func (list *SMoveList) addPawnCaptureMove(pos *SBoard, from, to, capture, side int) error {
	if !PieceValidEmpty(capture) {
		return errors.New("Capture piece is not valid")
	}
	if !SqOnBoard(from) {
		return errors.New("From square is not on board")
	}
	if !SqOnBoard(to) {
		return errors.New("To square is not on board")
	}
	if side == White {
		if RanksBrd[from] == Rank7 {
			list.addCaptureMove(pos, Move(from, to, capture, Wq, 0))
			list.addCaptureMove(pos, Move(from, to, capture, Wr, 0))
			list.addCaptureMove(pos, Move(from, to, capture, Wb, 0))
			list.addCaptureMove(pos, Move(from, to, capture, Wn, 0))
		} else {
			list.addCaptureMove(pos, Move(from, to, capture, Empty, 0))
		}
	} else {
		if RanksBrd[from] == Rank2 {
			list.addCaptureMove(pos, Move(from, to, capture, Bq, 0))
			list.addCaptureMove(pos, Move(from, to, capture, Br, 0))
			list.addCaptureMove(pos, Move(from, to, capture, Bb, 0))
			list.addCaptureMove(pos, Move(from, to, capture, Bn, 0))
		} else {
			list.addCaptureMove(pos, Move(from, to, capture, Empty, 0))
		}
	}
	return nil
}

// addPawnMove are possible quiet moves of pawn
func (list *SMoveList) addPawnMove(pos *SBoard, from, to, side int) error {
	if !SqOnBoard(from) {
		return errors.New("From square is not on board")
	}
	if !SqOnBoard(to) {
		return errors.New("To square is not on board")
	}
	if side == White {
		if RanksBrd[from] == Rank7 {
			list.addQuietMove(pos, Move(from, to, Empty, Wq, 0))
			list.addQuietMove(pos, Move(from, to, Empty, Wr, 0))
			list.addQuietMove(pos, Move(from, to, Empty, Wb, 0))
			list.addQuietMove(pos, Move(from, to, Empty, Wn, 0))
		} else {
			list.addQuietMove(pos, Move(from, to, Empty, Empty, 0))
		}
	} else {
		if RanksBrd[from] == Rank2 {
			list.addQuietMove(pos, Move(from, to, Empty, Bq, 0))
			list.addQuietMove(pos, Move(from, to, Empty, Br, 0))
			list.addQuietMove(pos, Move(from, to, Empty, Bb, 0))
			list.addQuietMove(pos, Move(from, to, Empty, Bn, 0))
		} else {
			list.addQuietMove(pos, Move(from, to, Empty, Empty, 0))
		}
	}
	return nil
}

// GenerateAllMoves will generate all possible moves of board
func (list *SMoveList) GenerateAllMoves(pos *SBoard) error {
	err := pos.Check()
	if err != nil {
		return err
	}
	list.count = 0

	if pos.Side == White {
		// Generating possible moves for white pawn
		for pieceNum := 0; pieceNum < pos.PceNum[Wp]; pieceNum++ {
			sq := pos.PList[Wp][pieceNum]
			if !SqOnBoard(sq) {
				return errors.New("Square is not on board")
			}
			if pos.Pieces[sq+10] == Empty {
				err = list.addPawnMove(pos, sq, sq+10, White)
				if err != nil {
					return err
				}
				if RanksBrd[sq] == Rank2 && pos.Pieces[sq+20] == Empty {
					list.addQuietMove(pos, Move(sq, sq+20, Empty, Empty, MFlagPS))
				}
			}
			if SqOnBoard(sq+9) && PieceCol[pos.Pieces[sq+9]] == Black {
				err = list.addPawnCaptureMove(pos, sq, sq+9, pos.Pieces[sq+9], White)
				if err != nil {
					return err
				}
			}
			if SqOnBoard(sq+11) && PieceCol[pos.Pieces[sq+11]] == Black {
				err = list.addPawnCaptureMove(pos, sq, sq+11, pos.Pieces[sq+11], White)
				if err != nil {
					return err
				}
			}
			if sq+9 == pos.EnPas {
				list.addCaptureMove(pos, Move(sq, sq+9, Empty, Empty, MFlagEP))
			}
			if sq+11 == pos.EnPas {
				list.addCaptureMove(pos, Move(sq, sq+11, Empty, Empty, MFlagEP))
			}
		}

		// Castling

		if (pos.CastlePerm & Wkca) > 0 {
			if pos.Pieces[F1] == Empty && pos.Pieces[G1] == Empty {
				resE1, _ := SqAttacked(E1, Black, pos)
				resF1, _ := SqAttacked(F1, Black, pos)
				if !resE1 && !resF1 {
					list.addQuietMove(pos, Move(E1, G1, Empty, Empty, MFlagCA))
				}
			}
		}

		if (pos.CastlePerm & Wqca) > 0 {
			if pos.Pieces[D1] == Empty && pos.Pieces[C1] == Empty && pos.Pieces[B1] == Empty {
				resE1, _ := SqAttacked(E1, Black, pos)
				resD1, _ := SqAttacked(D1, Black, pos)
				if !resE1 && !resD1 {
					list.addQuietMove(pos, Move(E1, C1, Empty, Empty, MFlagCA))
				}
			}
		}

	} else {
		// Generating possible moves for black pawns
		for pieceNum := 0; pieceNum < pos.PceNum[Bp]; pieceNum++ {
			sq := pos.PList[Bp][pieceNum]
			if !SqOnBoard(sq) {
				return errors.New("Square is not on board")
			}
			if pos.Pieces[sq-10] == Empty {
				err = list.addPawnMove(pos, sq, sq-10, Black)
				if err != nil {
					return err
				}
				if RanksBrd[sq] == Rank7 && pos.Pieces[sq-20] == Empty {
					list.addQuietMove(pos, Move(sq, sq-20, Empty, Empty, MFlagPS))
				}
			}
			if SqOnBoard(sq-9) && PieceCol[pos.Pieces[sq-9]] == White {
				err = list.addPawnCaptureMove(pos, sq, sq-9, pos.Pieces[sq-9], Black)
				if err != nil {
					return err
				}
			}
			if SqOnBoard(sq-11) && PieceCol[pos.Pieces[sq-11]] == White {
				err = list.addPawnCaptureMove(pos, sq, sq-11, pos.Pieces[sq-11], Black)
				if err != nil {
					return err
				}
			}
			if sq-9 == pos.EnPas {
				list.addCaptureMove(pos, Move(sq, sq-9, Empty, Empty, MFlagEP))
			}
			if sq-11 == pos.EnPas {
				list.addCaptureMove(pos, Move(sq, sq-11, Empty, Empty, MFlagEP))
			}
		}

		// Castling
		if (pos.CastlePerm & Bkca) > 0 {
			if pos.Pieces[F8] == Empty && pos.Pieces[G8] == Empty {
				resE8, _ := SqAttacked(E8, White, pos)
				resF8, _ := SqAttacked(F8, White, pos)
				if !resE8 && !resF8 {
					list.addQuietMove(pos, Move(E8, G8, Empty, Empty, MFlagCA))
				}
			}
		}

		if (pos.CastlePerm & Bqca) > 0 {
			if pos.Pieces[D8] == Empty && pos.Pieces[C8] == Empty && pos.Pieces[B8] == Empty {
				resE8, _ := SqAttacked(E8, White, pos)
				resD8, _ := SqAttacked(D8, White, pos)
				if !resE8 && !resD8 {
					list.addQuietMove(pos, Move(E8, C8, Empty, Empty, MFlagCA))
				}
			}
		}

	}

	// Loop for sliding pieces
	pieceIndex := LoopSlideIndex[pos.Side]
	piece := LoopSlidePce[pieceIndex]
	for piece != 0 {
		if !PieceValid(piece) {
			return errors.New("Piece is not valid")
		}
		fmt.Printf("sliders piece index: %d piece:%d \n", pieceIndex, piece)
		for pieceNum := 0; pieceNum < pos.PceNum[piece]; pieceNum++ {
			sq := pos.PList[piece][pieceNum]
			if !SqOnBoard(sq) {
				return errors.New("Square is not on board")
			}
			fmt.Printf("Piece: %c on %s \n", PceChar[piece], PrintSq(sq))
			for index := 0; index < NumDir[piece]; index++ {
				direction := PieceDir[piece][index]
				checkSq := sq + direction
				for SqOnBoard(checkSq) {
					if pos.Pieces[checkSq] != Empty {
						if PieceCol[pos.Pieces[checkSq]] == (pos.Side ^ 1) {
							list.addCaptureMove(pos, Move(sq, checkSq, pos.Pieces[checkSq], Empty, 0))
						}
						break
					}
					list.addQuietMove(pos, Move(sq, checkSq, Empty, Empty, 0))
					checkSq += direction
				}
			}
		}
		pieceIndex++
		piece = LoopSlidePce[pieceIndex]
	}

	// Loop for non-sliding pieces
	pieceIndex = LoopNonSlideIndex[pos.Side]
	piece = LoopNonSlidePce[pieceIndex]
	for piece != 0 {
		if !PieceValid(piece) {
			return errors.New("Piece is not valid")
		}
		fmt.Printf("non sliders piece index: %d piece: %d \n", pieceIndex, piece)
		for pieceNum := 0; pieceNum < pos.PceNum[piece]; pieceNum++ {
			sq := pos.PList[piece][pieceNum]
			if !SqOnBoard(sq) {
				return errors.New("Square is not on board")
			}
			fmt.Printf("Piece: %c on %s \n", PceChar[piece], PrintSq(sq))
			for index := 0; index < NumDir[piece]; index++ {
				direction := PieceDir[piece][index]
				checkSq := sq + direction
				if !SqOnBoard(checkSq) {
					continue
				}
				if pos.Pieces[checkSq] != Empty {
					if PieceCol[pos.Pieces[checkSq]] == (pos.Side ^ 1) {
						list.addCaptureMove(pos, Move(sq, checkSq, pos.Pieces[checkSq], Empty, 0))
					}
					continue
				}
				list.addQuietMove(pos, Move(sq, checkSq, Empty, Empty, 0))
			}
		}
		pieceIndex++
		piece = LoopNonSlidePce[pieceIndex]
	}
	return nil
}
