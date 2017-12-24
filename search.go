package main

import "fmt"

// isRepetition checks if current board structure is repetition of some previous board structure
func isRepetition(pos *SBoard) bool {
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
