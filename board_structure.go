package main

// SBoard is board structure
type SBoard struct {
	pieces     [BrdSqNum]int //all pieces on 120 board
	pawns      [3]U64        // 64 bit structure of white, black and both pawns
	kingSq     [2]int        // Positions of white, black and both kings
	side       int           // White or black side playing
	enPas      int           // enPas rule
	fiftyMove  int           // Number of move according to the fifty rule
	castlePerm int           // castling Permission
	ply        int           // Number of half moves played
	hisPly     int           // History of ply- highest number of moves played in game
	posKey     U64
	pceNum     [13]int
	bigPce     [3]int
	majPce     [3]int
	minPce     [3]int
	history    [MaxGameMoves]SUndo
	pList      [13][10]int //piece list
}
