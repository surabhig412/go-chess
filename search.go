package main

import (
	"fmt"
	"time"
)

type SearchInfo struct {
	startTime time.Time
	stopTime  time.Time
	depth     int
	depthSet  int
	timeSet   time.Time
	movesToGo int
	infinite  bool
	nodes     uint64
	quit      bool
	stopped   bool
}

// checkUp checks if time is up or there is an interrupt from GUI
func checkUp() {

}

// isRepetition checks if current board structure is repetition of some previous board structure
func isRepetition(pos *Board) bool {
	for index := pos.HisPly - pos.FiftyMove; index < pos.HisPly-1; index++ {
		if index < 0 || index >= MaxGameMoves {
			fmt.Println("Error in finding repetition")
			return false
		}
		if pos.PosKey == pos.History[index].PosKey {
			return true
		}
	}
	return false
}

// ClearForSearch clears the board before searching
func (si *SearchInfo) ClearForSearch(pos *Board) {
	for i := 0; i < 13; i++ {
		for j := 0; j < BrdSqNum; j++ {
			pos.SearchHistory[i][j] = 0
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < MaxDepth; j++ {
			pos.SearchKillers[i][j] = 0
		}
	}
	pos.PvTable.Clear()
	pos.Ply = 0
	si.startTime = time.Now()
	si.stopTime = time.Time{}
	si.nodes = 0
}

// Quiescence search does alpha beta considering all the quiet moves. It covers horizon effects on the chess board
func (si *SearchInfo) Quiescence(alpha, beta int, pos *Board) int {
	return 0
}

func (si *SearchInfo) AlphaBeta(alpha, beta, depth int, pos *Board, doNull bool) int {
	return 0
}

func (si *SearchInfo) SearchPosition(pos *Board) {

}
