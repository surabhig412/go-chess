package main

import "log"

// GeneratePosKey generates posKey of the board
func GeneratePosKey(pos *Board) uint64 {
	piece := Empty
	var finalKey uint64
	// Include hash key of each piece on the board
	for sq := 0; sq < BrdSqNum; sq++ {
		piece = pos.Pieces[sq]
		if piece < Wp || piece > Bk {
			continue
		}
		finalKey ^= PieceKeys[piece][sq]
	}
	// Include hash key of the side to play
	if pos.Side == White {
		finalKey ^= SideKey
	}
	// Include hash key of enPas square
	if pos.EnPas != NoSq {
		if !(pos.EnPas >= 0 && pos.EnPas < BrdSqNum) {
			log.Fatalln("Enpas of board structure is not proper")
		}
		finalKey ^= PieceKeys[Empty][pos.EnPas]
	}
	// Include hash key of castle permission
	if !(pos.CastlePerm >= 0 && pos.CastlePerm <= 15) {
		log.Fatalln("Castle permissions given to board structure is improper")
	}
	finalKey ^= CastleKeys[pos.CastlePerm]
	return finalKey
}
