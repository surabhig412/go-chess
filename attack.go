package main

import "errors"

// SqAttacked evaluates if a particular square(sq) is attacked by opposing side(side) or not
func SqAttacked(sq, side int, pos *SBoard) (bool, error) {
	// check if square, side and position of board is valid or not
	if !SqOnBoard(sq) {
		return false, errors.New("Square is not on board")
	}
	if !SideValid(side) {
		return false, errors.New("Side is invalid")
	}
	// err := pos.Check()
	// if err != nil {
	// 	return false, err
	// }
	// pawns are attacking or not
	if side == White {
		if (pos.Pieces[sq-9] == Wp) || (pos.Pieces[sq-11] == Wp) {
			return true, nil
		}
	} else {
		if (pos.Pieces[sq+9] == Bp) || (pos.Pieces[sq+11] == Bp) {
			return true, nil
		}
	}

	// knights are attacking or not
	for index := 0; index < 8; index++ {
		piece := pos.Pieces[sq+KnDir[index]]
		if (piece != Offboard) && IsKn(piece) == True && PieceCol[piece] == side {
			return true, nil
		}
	}

	// rooks or queens are attacking or not
	for index := 0; index < 4; index++ {
		direction := RkDir[index]
		checkSq := sq + direction
		piece := pos.Pieces[checkSq]
		for piece != Offboard {
			if piece != Empty {
				if IsRQ(piece) == True && PieceCol[piece] == side {
					return true, nil
				}
				break
			}
			checkSq += direction
			piece = pos.Pieces[checkSq]
		}
	}

	// bishops or queens are attacking or not
	for index := 0; index < 4; index++ {
		direction := BiDir[index]
		checkSq := sq + direction
		piece := pos.Pieces[checkSq]
		for piece != Offboard {
			if piece != Empty {
				if IsBQ(piece) == True && PieceCol[piece] == side {
					return true, nil
				}
				break
			}
			checkSq += direction
			piece = pos.Pieces[checkSq]
		}
	}

	// kings are attacking or not
	for index := 0; index < 8; index++ {
		piece := pos.Pieces[sq+KiDir[index]]
		if (piece != Offboard) && IsKi(piece) == True && PieceCol[piece] == side {
			return true, nil
		}
	}
	return false, nil
}
