package board

import (
	"fmt"
	"go-chess/constants"
	"go-chess/models"
	"go-chess/utils"
)

// PrintBoard prints the entire chess board
func PrintBoard(pos *models.SBoard) {

	fmt.Println("\nGame Board:")

	for rank := constants.Rank8; rank >= constants.Rank1; rank-- {
		fmt.Printf("%d  ", rank+1)
		for file := constants.FileA; file <= constants.FileH; file++ {
			sq := utils.FR2SQ(file, rank)
			piece := pos.Pieces[sq]
			fmt.Printf("%3c", constants.PceChar[piece])
		}
		fmt.Println()
	}

	fmt.Printf("\n   ")
	for file := constants.FileA; file <= constants.FileH; file++ {
		fmt.Printf("%3c", rune(97+file))
	}
	fmt.Println()
	fmt.Printf("side:%c\n", constants.SideChar[pos.Side])
	fmt.Printf("enPas:%d\n", pos.EnPas)
	castling := ""
	if (pos.CastlePerm & constants.Wkca) == constants.Wkca {
		castling += "K"
	} else {
		castling += "-"
	}
	if (pos.CastlePerm & constants.Wqca) == constants.Wqca {
		castling += "Q"
	} else {
		castling += "-"
	}
	if (pos.CastlePerm & constants.Bkca) == constants.Bkca {
		castling += "k"
	} else {
		castling += "-"
	}
	if (pos.CastlePerm & constants.Bqca) == constants.Bqca {
		castling += "q"
	} else {
		castling += "-"
	}
	fmt.Printf("castle:%s\n", castling)
	fmt.Printf("PosKey:%x\n", pos.PosKey)
}
