package main

import "strconv"

// PrintSq prints the algebraic notation of particular square
func PrintSq(sq int) string {
	file := FilesBrd[sq]
	rank := RanksBrd[sq]

	algebraicSq := string(rune(97+file)) + strconv.Itoa(rank+1)

	return algebraicSq
}

// PrintMove prints full move with promoted square
func PrintMove(move int) string {
	fileFrom := FilesBrd[FromSq(move)]
	rankFrom := RanksBrd[FromSq(move)]
	fileTo := FilesBrd[ToSq(move)]
	rankTo := RanksBrd[ToSq(move)]

	promoted := Promoted(move)
	if promoted > 0 {
		promotedPiece := "q"
		if IsKn(promoted) == True {
			promotedPiece = "n"
		} else if IsRQ(promoted) == True && IsBQ(promoted) == False {
			promotedPiece = "r"
		} else if IsRQ(promoted) == False && IsBQ(promoted) == True {
			promotedPiece = "b"
		}
		return (string(rune(97+fileFrom)) + strconv.Itoa(rankFrom+1) + string(rune(97+fileTo)) + strconv.Itoa(rankTo+1) + promotedPiece)
	}
	return (string(rune(97+fileFrom)) + strconv.Itoa(rankFrom+1) + string(rune(97+fileTo)) + strconv.Itoa(rankTo+1))
}
