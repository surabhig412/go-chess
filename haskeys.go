package main

import "log"

// GeneratePosKey generates posKey of the board
func GeneratePosKey(pos *SBoard) uint64 {
	piece := Empty
	var finalKey uint64
	for sq := 0; sq < BrdSqNum; sq++ {
		piece = pos.Pieces[sq]
		if piece < Wp || piece > Bk {
			continue
		}
		finalKey ^= PieceKeys[piece][sq]
	}
	if pos.Side == White {
		finalKey ^= SideKey
	}
	if pos.EnPas != NoSq {
		if !(pos.EnPas >= 0 && pos.EnPas < BrdSqNum) {
			log.Fatalln("Enpas of board structure is not proper")
		}
		finalKey ^= PieceKeys[Empty][pos.EnPas]
	}
	if !(pos.CastlePerm >= 0 && pos.CastlePerm <= 15) {
		log.Fatalln("Castle permissions given to board structure is improper")
	}
	finalKey ^= CastleKeys[pos.CastlePerm]
	return finalKey
}
