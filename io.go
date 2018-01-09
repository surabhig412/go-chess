package main

import (
	"errors"
	"regexp"
	"strconv"
)

// PrintSq prints the algebraic notation of particular square
func PrintSq(sq int) string {
	return string(rune(97+FilesBrd[sq])) + strconv.Itoa(RanksBrd[sq]+1)
}

// PrintMove prints full move with promoted square(convert move in integer to algebraic format)
func PrintMove(move int) string {
	if move == NoMove {
		return "NoMove"
	}
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

// ParseMove converts algebraic move to its corresponding integer format
func ParseMove(algebraicMove string, pos *Board) (int, error) {
	var list MoveList
	re := regexp.MustCompile("[a-h][1-8][a-h][1-8][qnrb]?")
	if !re.MatchString(algebraicMove) {
		return NoMove, errors.New("Invalid move provided")
	}
	fromSq := FR2SQ(int(algebraicMove[0])-97, int(algebraicMove[1])-49)
	toSq := FR2SQ(int(algebraicMove[2])-97, int(algebraicMove[3])-49)

	if !SqOnBoard(fromSq) {
		return NoMove, errors.New("Square from where piece is to be moved is not on board")
	}
	if !SqOnBoard(toSq) {
		return NoMove, errors.New("Square where piece is to be moved is not on board")
	}

	(&list).GenerateAllMoves(pos)
	for i := 0; i < list.count; i++ {
		move := list.moves[i].move
		if FromSq(move) == fromSq && ToSq(move) == toSq {
			promotedPiece := Promoted(move)
			if promotedPiece != Empty && len(algebraicMove) >= 5 {
				if (IsRQ(promotedPiece) == True && IsBQ(promotedPiece) == False && algebraicMove[4] == 'r') ||
					(IsRQ(promotedPiece) == False && IsBQ(promotedPiece) == True && algebraicMove[4] == 'b') ||
					(IsRQ(promotedPiece) == True && IsBQ(promotedPiece) == True && algebraicMove[4] == 'q') ||
					(IsKn(promotedPiece) == True && algebraicMove[4] == 'n') {
					return move, nil
				}
				continue
			}
			return move, nil
		}
	}
	return NoMove, nil
}
