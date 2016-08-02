package engine

import (
	"go-chess/constants"
	"go-chess/models"
	"go-chess/utils"
)

// ResetBoard resets the chess board
func ResetBoard(pos *models.SBoard) {
	for index := 0; index < constants.BrdSqNum; index++ {
		pos.Pieces[index] = constants.Offboard
	}
	for index := 0; index < 64; index++ {
		pos.Pieces[utils.SQ120(index)] = constants.Empty
	}
	for index := 0; index < 3; index++ {
		pos.BigPce[index] = 0
		pos.MajPce[index] = 0
		pos.MinPce[index] = 0
		pos.Pawns[index] = constants.U64(0)
	}
	for index := 0; index < 13; index++ {
		pos.PceNum[index] = 0
	}
	pos.KingSq[constants.White] = constants.NoSq
	pos.KingSq[constants.Black] = constants.NoSq
	pos.Side = constants.Both
	pos.EnPas = constants.NoSq
	pos.FiftyMove = 0
	pos.Ply = 0
	pos.HisPly = 0
	pos.CastlePerm = 0
	pos.PosKey = constants.U64(0)
}
