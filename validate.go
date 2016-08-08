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
