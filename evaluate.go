package main

import "errors"

// EvalPosition evaluates material score of the board in its current state
func EvalPosition(pos *Board) (int, error) {
	score := pos.Material[White] - pos.Material[Black]

	piece := Wp
	for i := 0; i < pos.PceNum[piece]; i++ {
		sq := pos.PList[piece][i]
		if !SqOnBoard(sq) {
			return 0, errors.New("Square whose score is to be evaluated is not on board")
		}
		score += PawnTable[SQ64(sq)]
	}

	piece = Bp
	for i := 0; i < pos.PceNum[piece]; i++ {
		sq := pos.PList[piece][i]
		if !SqOnBoard(sq) {
			return 0, errors.New("Square whose score is to be evaluated is not on board")
		}
		score -= PawnTable[Mirror64[SQ64(sq)]]
	}

	piece = Wn
	for i := 0; i < pos.PceNum[piece]; i++ {
		sq := pos.PList[piece][i]
		if !SqOnBoard(sq) {
			return 0, errors.New("Square whose score is to be evaluated is not on board")
		}
		score += KnightTable[SQ64(sq)]
	}

	piece = Bn
	for i := 0; i < pos.PceNum[piece]; i++ {
		sq := pos.PList[piece][i]
		if !SqOnBoard(sq) {
			return 0, errors.New("Square whose score is to be evaluated is not on board")
		}
		score -= KnightTable[Mirror64[SQ64(sq)]]
	}

	piece = Wb
	for i := 0; i < pos.PceNum[piece]; i++ {
		sq := pos.PList[piece][i]
		if !SqOnBoard(sq) {
			return 0, errors.New("Square whose score is to be evaluated is not on board")
		}
		score += BishopTable[SQ64(sq)]
	}

	piece = Bb
	for i := 0; i < pos.PceNum[piece]; i++ {
		sq := pos.PList[piece][i]
		if !SqOnBoard(sq) {
			return 0, errors.New("Square whose score is to be evaluated is not on board")
		}
		score -= BishopTable[Mirror64[SQ64(sq)]]
	}

	piece = Wr
	for i := 0; i < pos.PceNum[piece]; i++ {
		sq := pos.PList[piece][i]
		if !SqOnBoard(sq) {
			return 0, errors.New("Square whose score is to be evaluated is not on board")
		}
		score += RookTable[SQ64(sq)]
	}

	piece = Br
	for i := 0; i < pos.PceNum[piece]; i++ {
		sq := pos.PList[piece][i]
		if !SqOnBoard(sq) {
			return 0, errors.New("Square whose score is to be evaluated is not on board")
		}
		score -= RookTable[Mirror64[SQ64(sq)]]
	}

	if pos.Side == White {
		return score, nil
	} else {
		return -score, nil
	}
}
