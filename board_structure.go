package main

// SBoard is board structure
type SBoard struct {
	pieces     [BrdSqNum]int
	pawns      [3]U64
	kingSq     [2]int
	side       int
	enPas      int
	fiftyMove  int
	castlePerm int
	ply        int
	hisPly     int
	posKey     U64
	pceNum     [13]int
	bigPce     [3]int
	majPce     [3]int
	minPce     [3]int
	history    [MaxGameMoves]SUndo
	pList      [13][10]int //piece list
}
