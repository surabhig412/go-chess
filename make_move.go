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

// AddPiece adds the piece when making a move
func AddPiece(sq, piece int, pos *SBoard) error {
	if !SqOnBoard(sq) {
		return errors.New("Square where piece is to be added is not on board")
	}
	err := pos.Check()
	if err != nil {
		return err
	}

	if !PieceValid(piece) {
		return errors.New("Piece to be added is invalid")
	}
	colour := PieceCol[piece]
	if !SideValid(colour) {
		return errors.New("Side is invalid")
	}

	// Hash out hash key of piece from PosKey
	pos.PosKey ^= PieceKeys[piece][sq]

	// Add values for the piece to arrays representing the board
	pos.Pieces[sq] = piece
	pos.Material[colour] += PieceVal[piece]
	if PieceBig[piece] == True {
		pos.BigPce[colour]++
		if PieceMaj[piece] == True {
			pos.MajPce[colour]++
		} else {
			pos.MinPce[colour]++
		}
	} else {
		pawnCol := Bitboard(pos.Pawns[colour])
		(&pawnCol).Set(SQ64(sq))
		pos.Pawns[colour] = uint64(pawnCol)
		pawnBoth := Bitboard(pos.Pawns[Both])
		(&pawnBoth).Set(SQ64(sq))
		pos.Pawns[Both] = uint64(pawnBoth)
	}

	pos.PList[piece][pos.PceNum[piece]] = sq
	pos.PceNum[piece]++
	return nil
}

// MovePiece moves the piece from "from" sq to "to" sq on the board
func MovePiece(from, to int, pos *SBoard) error {
	if !SqOnBoard(from) {
		return errors.New("Square from where piece is to be moved is not on board")
	}
	if !SqOnBoard(to) {
		return errors.New("Square where piece is to be moved is not on board")
	}
	err := pos.Check()
	if err != nil {
		return err
	}

	piece := pos.Pieces[from]
	if !PieceValid(piece) {
		return errors.New("Piece to be added is invalid")
	}
	colour := PieceCol[piece]
	if !SideValid(colour) {
		return errors.New("Side is invalid")
	}
	targetPceNum := false

	// Hash out hash key of piece at "from" square from PosKey
	pos.PosKey ^= PieceKeys[piece][from]
	pos.Pieces[from] = Empty

	// Hash in hash key of piece at "to" square into PosKey
	pos.PosKey ^= PieceKeys[piece][to]
	pos.Pieces[to] = piece

	// Material and big pieces remain the same
	// Set and clear the pawn bitboard for "to" and "from" squares respectively
	if PieceBig[piece] == False {
		pawnCol := Bitboard(pos.Pawns[colour])
		(&pawnCol).Clear(SQ64(from))
		(&pawnCol).Set(SQ64(to))
		pos.Pawns[colour] = uint64(pawnCol)
		pawnBoth := Bitboard(pos.Pawns[Both])
		(&pawnBoth).Clear(SQ64(from))
		(&pawnBoth).Set(SQ64(to))
		pos.Pawns[Both] = uint64(pawnBoth)
	}

	// Move the piece in Plist
	for index := 0; index < pos.PceNum[piece]; index++ {
		if pos.PList[piece][index] == from {
			pos.PList[piece][index] = to
			targetPceNum = true
			break
		}
	}

	if targetPceNum == false {
		return errors.New("Piece not found in piece list, invalid board structure")
	}

	return nil
}
