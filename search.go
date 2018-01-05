package main

import (
	"fmt"
	"log"
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
	fh        float32 //fail high
	fhf       float32 // fail high first
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

// ClearForSearch clears the board and searchInfo before searching
func ClearForSearch(pos *Board, info *SearchInfo) {
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
	info.startTime = time.Now()
	info.stopTime = time.Time{}
	info.nodes = 0
	info.fh = 0
	info.fhf = 0
}

// Quiescence search does alpha beta considering all the quiet moves. It covers horizon effects on the chess board
func Quiescence(alpha, beta int, info *SearchInfo, pos *Board) int {
	err := pos.Check()
	if err != nil {
		log.Fatalln("Error occurred in Quiescence")
	}
	info.nodes++
	if isRepetition(pos) || pos.FiftyMove >= 100 {
		return 0
	}
	if pos.Ply > MaxDepth-1 {
		score, _ := EvalPosition(pos)
		return score
	}
	score, _ := EvalPosition(pos)
	if score >= beta {
		return beta
	}
	if score > alpha {
		alpha = score
	}
	var list MoveList
	OnlyCapturedMoves = true
	list.GenerateAllMoves(pos)
	legalMovesCount := 0
	oldAlpha := alpha
	bestMove := NoMove
	score = -Infinite

	for i := 0; i < list.count; i++ {
		list.PickNextMove(i)
		moveMade, _ := MakeMove(list.moves[i].move, pos)
		if !moveMade {
			continue
		}
		legalMovesCount++
		score = -Quiescence(-beta, -alpha, info, pos)
		TakeMove(pos)
		if score > alpha {
			if score >= beta {
				if legalMovesCount == 1 {
					info.fhf++
				}
				info.fh++
				return beta
			}
			alpha = score
			bestMove = list.moves[i].move
		}
	}

	if alpha != oldAlpha {
		StorePvMove(pos, bestMove)
	}
	return alpha
}

func AlphaBeta(alpha, beta, depth int, info *SearchInfo, pos *Board, doNull bool) int {
	err := pos.Check()
	if err != nil {
		log.Fatalln("Invalid board position")
		return 0
	}
	if depth == 0 {
		return Quiescence(alpha, beta, info, pos)
	}
	info.nodes++
	// draw condition
	if isRepetition(pos) || pos.FiftyMove >= 100 {
		return 0
	}
	// when depth has exceeded maxdepth
	if pos.Ply > MaxDepth-1 {
		score, _ := EvalPosition(pos)
		return score
	}
	var list MoveList
	list.GenerateAllMoves(pos)
	legalMovesCount := 0 // if no legal moves then it is condition of check mate or stale mate
	oldAlpha := alpha
	bestMove := NoMove
	score := -Infinite
	pvMove, _ := ProbePvTable(pos)
	if pvMove != NoMove {
		for i := 0; i < list.count; i++ {
			if list.moves[i].move == pvMove {
				list.moves[i].score = 2000000
				break
			}
		}
	}
	for i := 0; i < list.count; i++ {
		list.PickNextMove(i)
		moveMade, _ := MakeMove(list.moves[i].move, pos)
		if !moveMade {
			continue
		}
		legalMovesCount++
		score = -AlphaBeta(-beta, -alpha, depth-1, info, pos, doNull)
		TakeMove(pos)
		if score > alpha {
			if score >= beta {
				if legalMovesCount == 1 {
					info.fhf++
				}
				info.fh++
				if (list.moves[i].move & MFlagCAP) == 0 {
					pos.SearchKillers[1][pos.Ply] = pos.SearchKillers[0][pos.Ply]
					pos.SearchKillers[0][pos.Ply] = list.moves[i].move
				}
				return beta
			}
			alpha = score
			bestMove = list.moves[i].move
			if (list.moves[i].move & MFlagCAP) == 0 {
				pos.SearchHistory[pos.Pieces[FromSq(bestMove)]][ToSq(bestMove)] += depth
			}
		}
	}
	if legalMovesCount == 0 {
		if attacked, _ := SqAttacked(pos.KingSq[pos.Side], pos.Side^1, pos); attacked {
			return -Mate + pos.Ply
		} else {
			return 0
		}
	}
	if alpha != oldAlpha {
		StorePvMove(pos, bestMove)
	}
	return alpha
}

func SearchPosition(pos *Board, info *SearchInfo) {
	bestMove := NoMove
	bestScore := -Infinite
	currentDepth := 0
	ClearForSearch(pos, info)
	// iterative deepening
	for currentDepth = 1; currentDepth <= info.depth; currentDepth++ {
		bestScore = AlphaBeta(-Infinite, Infinite, currentDepth, info, pos, true)
		// out of time check to be included
		pvMoves, _ := GetPvLine(currentDepth, pos)
		bestMove = pos.PvArray[0]
		if bestMove != NoMove {
			fmt.Printf("Depth:%d, score: %d, move: %s, nodes: %v pv", currentDepth, bestScore, PrintMove(bestMove), info.nodes)
			for i := 0; i < pvMoves; i++ {
				fmt.Printf(" %s", PrintMove(pos.PvArray[i]))
			}
			fmt.Println()
			fmt.Println("Ordering: ", info.fhf/info.fh)
		}
	}
}
