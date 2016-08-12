package main

// SqOnBoard validates if square is on board or not
func SqOnBoard(sq int) bool {
	if FilesBrd[sq] == Offboard {
		return false
	}
	return true
}

// SideValid validates if side is white or black or not
func SideValid(side int) bool {
	if side == White || side == Black {
		return true
	}
	return false
}

// FileRankValid validates if file or rank is valid or not
func FileRankValid(fr int) bool {
	if fr >= 0 && fr <= 7 {
		return true
	}
	return false
}

// PieceValidEmpty validates if piece is valid or empty
func PieceValidEmpty(piece int) bool {
	if piece >= Empty && piece <= Bk {
		return true
	}
	return false
}

// PieceValid validates if piece is valid
func PieceValid(piece int) bool {
	if piece >= Wp && piece <= Bk {
		return true
	}
	return false
}
