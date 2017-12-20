package main

import "errors"

// ClearPiece clears the piece when making a move
func ClearPiece(sq int, pos *SBoard) error {
	targetPceNum := -1
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

	// Hash in hash key of piece into PosKey
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

func MakeMove(move int, pos *SBoard) (bool, error) {
	err := pos.Check()
	if err != nil {
		return false, err
	}

	fromSq := FromSq(move)
	toSq := ToSq(move)
	side := pos.Side
	capturedPiece := Captured(move)
	promotedPiece := Promoted(move)

	if !SqOnBoard(fromSq) {
		return false, errors.New("Square from where piece is to be moved is not on board")
	}
	if !SqOnBoard(toSq) {
		return false, errors.New("Square where piece is to be moved is not on board")
	}
	if !SideValid(side) {
		return false, errors.New("Side is invalid")
	}
	if !PieceValid(pos.Pieces[fromSq]) {
		return false, errors.New("Piece to be added is invalid")
	}

	pos.History[pos.HisPly].PosKey = pos.PosKey
	if (move & MFlagEP) == 1 {
		if side == White {
			ClearPiece(toSq-10, pos)
		} else {
			ClearPiece(toSq+10, pos)
		}
	} else if (move & MFlagCA) == 1 {
		switch toSq {
		case C1:
			MovePiece(A1, D1, pos)
			break
		case C8:
			MovePiece(A8, D8, pos)
			break
		case G1:
			MovePiece(H1, F1, pos)
			break
		case G8:
			MovePiece(H8, F8, pos)
			break
		default:
			return false, errors.New("Error in move")
		}
	}

	if pos.EnPas != NoSq {
		// Hash out hash key of enPas piece from PosKey
		pos.PosKey ^= PieceKeys[Empty][pos.EnPas]
	}
	// Hash out hash key of castling from PosKey
	pos.PosKey ^= CastleKeys[pos.CastlePerm]

	pos.History[pos.HisPly].Move = move
	pos.History[pos.HisPly].FiftyMove = pos.FiftyMove
	pos.History[pos.HisPly].EnPas = pos.EnPas
	pos.History[pos.HisPly].CastlePerm = pos.CastlePerm
	pos.CastlePerm &= CastlePerm[fromSq]
	pos.CastlePerm &= CastlePerm[toSq]
	pos.EnPas = NoSq
	pos.FiftyMove++
	pos.HisPly++
	pos.Ply++
	// Hash in hash key of castling into PosKey
	pos.PosKey ^= CastleKeys[pos.CastlePerm]

	if capturedPiece != Empty {
		if !PieceValid(capturedPiece) {
			return false, errors.New("Piece to be captured is invalid")
		}
		ClearPiece(toSq, pos)
		pos.FiftyMove = 0
	}

	if PiecePawn[pos.Pieces[fromSq]] == True {
		pos.FiftyMove = 0
		if (move & MFlagPS) == 1 {
			if side == White {
				pos.EnPas = fromSq + 10
				if RanksBrd[pos.EnPas] != Rank3 {
					return false, errors.New("Rank of enPas should be rank 3")
				}
			} else {
				pos.EnPas = fromSq - 10
				if RanksBrd[pos.EnPas] != Rank6 {
					return false, errors.New("Rank of enPas should be rank 6")
				}
			}
			// Hash in hash key of enPas piece into PosKey
			pos.PosKey ^= PieceKeys[Empty][pos.EnPas]
		}
	}

	MovePiece(fromSq, toSq, pos)

	if promotedPiece != Empty {
		if !PieceValid(promotedPiece) || PiecePawn[promotedPiece] == True {
			return false, errors.New("Piece to be promoted is invalid")
		}
		ClearPiece(toSq, pos)
		AddPiece(toSq, promotedPiece, pos)
	}

	if PieceKing[pos.Pieces[toSq]] == True {
		pos.KingSq[pos.Side] = toSq
	}

	pos.Side ^= 1
	pos.PosKey ^= SideKey

	err = pos.Check()
	if err != nil {
		return false, err
	}

	if attacked, _ := SqAttacked(pos.KingSq[side], pos.Side, pos); attacked {
		TakeMove(pos)
		return false, errors.New("Move taken back as King will be attacked by the move")
	}

	return true, nil
}

func TakeMove(pos *SBoard) error {
	err := pos.Check()
	if err != nil {
		return err
	}

	pos.HisPly--
	pos.Ply--
	move := pos.History[pos.HisPly].Move
	fromSq := FromSq(move)
	toSq := ToSq(move)
	capturedPiece := Captured(move)
	promotedPiece := Promoted(move)
	if !SqOnBoard(fromSq) {
		return errors.New("Square from where piece is to be moved is not on board")
	}
	if !SqOnBoard(toSq) {
		return errors.New("Square where piece is to be moved is not on board")
	}
	if pos.EnPas != NoSq {
		// Hash out hash key of enPas piece from PosKey
		pos.PosKey ^= PieceKeys[Empty][pos.EnPas]
	}
	// Hash out hash key of castling from PosKey
	pos.PosKey ^= CastleKeys[pos.CastlePerm]

	pos.CastlePerm = pos.History[pos.HisPly].CastlePerm
	pos.FiftyMove = pos.History[pos.HisPly].FiftyMove
	pos.EnPas = pos.History[pos.HisPly].EnPas
	if pos.EnPas != NoSq {
		// Hash in hash key of enPas piece into PosKey
		pos.PosKey ^= PieceKeys[Empty][pos.EnPas]
	}
	// Hash in hash key of castling into PosKey
	pos.PosKey ^= CastleKeys[pos.CastlePerm]
	pos.Side ^= 1
	pos.PosKey ^= SideKey

	if (move & MFlagEP) == 1 {
		if pos.Side == White {
			AddPiece(toSq-10, Bp, pos)
		} else {
			AddPiece(toSq+10, Wp, pos)
		}
	} else if (move & MFlagCA) == 1 {
		switch toSq {
		case C1:
			MovePiece(D1, A1, pos)
			break
		case C8:
			MovePiece(D8, A8, pos)
			break
		case G1:
			MovePiece(F1, H1, pos)
			break
		case G8:
			MovePiece(F8, H8, pos)
			break
		default:
			return errors.New("Error in move")
		}
	}

	MovePiece(toSq, fromSq, pos)

	if PieceKing[pos.Pieces[fromSq]] == True {
		pos.KingSq[pos.Side] = fromSq
	}
	if capturedPiece != Empty {
		if !PieceValid(capturedPiece) {
			return errors.New("Piece to be captured is invalid")
		}
		AddPiece(toSq, capturedPiece, pos)
	}

	if promotedPiece != Empty {
		if !PieceValid(promotedPiece) || PiecePawn[promotedPiece] == True {
			return errors.New("Piece to be promoted is invalid")
		}
		ClearPiece(fromSq, pos)
		if PieceCol[promotedPiece] == White {
			AddPiece(fromSq, Wp, pos)
		} else {
			AddPiece(fromSq, Bp, pos)
		}
	}
	err = pos.Check()
	if err != nil {
		return err
	}
	return nil
}
