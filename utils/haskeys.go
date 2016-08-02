package utils

import (
	"go-chess/constants"
	"go-chess/models"
	"log"
)

// GeneratePosKey generates posKey of the board
func GeneratePosKey(pos *models.SBoard) uint64 {
	piece := constants.Empty
	var finalKey uint64
	for sq := 0; sq < constants.BrdSqNum; sq++ {
		piece = pos.Pieces[sq]
		if piece != constants.NoSq && piece != constants.Empty {
			if !(piece >= constants.Wp && piece <= constants.Bk) {
				log.Fatalln("Piece taken is not proper")
			}
			finalKey ^= constants.PieceKeys[piece][sq]
		}
	}
	if pos.Side == constants.White {
		finalKey ^= constants.SideKey
	}
	if pos.EnPas != constants.NoSq {
		if !(pos.EnPas >= 0 && pos.EnPas < constants.BrdSqNum) {
			log.Fatalln("Enpas of board structure is not proper")
		}
		finalKey ^= constants.PieceKeys[constants.Empty][pos.EnPas]
	}
	if !(pos.CastlePerm >= 0 && pos.CastlePerm <= 15) {
		log.Fatalln("Castle permissions given to board structure is improper")
	}
	finalKey ^= constants.CastleKeys[pos.CastlePerm]
	return finalKey
}
