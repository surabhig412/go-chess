package models

import "go-chess/constants"

// SBoard is board structure
type SBoard struct {
	Pieces     [constants.BrdSqNum]int //all pieces on 120 board
	Pawns      [3]uint64               // 64 bit structure of white, black and both pawns
	KingSq     [2]int                  // Positions of white, black and both kings
	Side       int                     // White or black side playing
	EnPas      int                     // enPas rule
	FiftyMove  int                     // Number of move according to the fifty rule
	CastlePerm int                     // castling Permission
	Ply        int                     // Number of half moves played
	HisPly     int                     // History of ply- highest number of moves played in game
	PosKey     uint64
	PceNum     [13]int
	BigPce     [3]int
	MajPce     [3]int
	MinPce     [3]int
	History    [constants.MaxGameMoves]SUndo
	PList      [13][10]int //piece list
}
