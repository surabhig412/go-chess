package main

import "errors"

// ClearPiece clears the piece when making a move
func ClearPiece(sq int, pos *SBoard) error {
	if !SqOnBoard(sq) {
		return errors.New("Square to be cleared is not on board")
	}
	err := pos.Check()
	if err != nil {
		return err
	}

	piece := pos.Pieces[sq]
	if !PieceValid(piece) {
		return errors.New("Piece to be cleared is invalid")
	}
	colour := PieceCol[piece]
	if !SideValid(colour) {
		return errors.New("Side is invalid")
	}
	targetPceNum := -1

	// Hash out hash key of piece from PosKey
	pos.PosKey ^= PieceKeys[piece][sq]

	// Clear values for the piece from arrays representing the board
	pos.Pieces[sq] = Empty
	pos.Material[colour] -= PieceVal[piece]
	if PieceBig[piece] == True {
		pos.BigPce[colour]--
		if PieceMaj[piece] == True {
			pos.MajPce[colour]--
		} else {
			pos.MinPce[colour]--
		}
	} else {
		pawnCol := Bitboard(pos.Pawns[colour])
		(&pawnCol).Clear(SQ64(sq))
		pos.Pawns[colour] = uint64(pawnCol)
		pawnBoth := Bitboard(pos.Pawns[Both])
		(&pawnBoth).Clear(SQ64(sq))
		pos.Pawns[Both] = uint64(pawnBoth)
	}

	// Remove the piece from PceNum and PList of the board
	for index := 0; index < pos.PceNum[piece]; index++ {
		if pos.PList[piece][index] == sq {
			targetPceNum = index
			break
		}
	}

	if targetPceNum == -1 {
		return errors.New("Piece not found in piece list, invalid board structure")
	}
	if targetPceNum < 0 || targetPceNum >= 10 {
		return errors.New("Index of piece invalid, invalid board structure")
	}

	pos.PceNum[piece]--
	pos.PList[piece][targetPceNum] = pos.PList[piece][pos.PceNum[piece]]
	return nil
}
