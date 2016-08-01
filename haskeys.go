package main

import "log"

// GeneratePosKey generates posKey of the board
func GeneratePosKey(pos *SBoard) U64 {
	piece := Empty
	var finalKey U64
	for sq := 0; sq < BrdSqNum; sq++ {
		piece = pos.pieces[sq]
		if piece != NoSq && piece != Empty {
			if !(piece >= wP && piece <= bK) {
				log.Fatalln("Piece taken is not proper")
			}
			finalKey ^= PieceKeys[piece][sq]
		}
	}
	if pos.side == White {
		finalKey ^= SideKey
	}
	if pos.enPas != NoSq {
		if !(pos.enPas >= 0 && pos.enPas < BrdSqNum) {
			log.Fatalln("Enpas of board structure is not proper")
		}
		finalKey ^= PieceKeys[Empty][pos.enPas]
	}
	if !(pos.castlePerm >= 0 && pos.castlePerm <= 15) {
		log.Fatalln("Castle permissions given to board structure is improper")
	}
	finalKey ^= CastleKeys[pos.castlePerm]
	return finalKey
}
