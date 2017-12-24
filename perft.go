package main

import (
	"fmt"
	"time"
)

var leafNodes int

func perft(depth int, pos *SBoard) error {
	err := pos.Check()
	if err != nil {
		return err
	}
	if depth == 0 {
		leafNodes++
		return nil
	}
	var list SMoveList
	list.GenerateAllMoves(pos)
	for i := 0; i < list.count; i++ {
		moveMade, _ := MakeMove(list.moves[i].move, pos)
		if !moveMade {
			// fmt.Println("movenum, Move, error: ", i, list.moves[i].move, err)
			continue
		}
		perft(depth-1, pos)
		TakeMove(pos)
	}
	return nil
}

func PerftTest(depth int, pos *SBoard) error {
	startTime := time.Now()
	err := pos.Check()
	if err != nil {
		return err
	}
	pos.Print()
	fmt.Printf("Starting test to depth %d \n", depth)

	leafNodes = 0
	var list SMoveList
	list.GenerateAllMoves(pos)
	for i := 0; i < list.count; i++ {
		move := list.moves[i].move
		moveMade, _ := MakeMove(move, pos)
		if !moveMade {
			continue
		}
		cumulativeNodes := leafNodes
		perft(depth-1, pos)
		TakeMove(pos)
		oldNodes := leafNodes - cumulativeNodes
		fmt.Printf("Move %d: %s : %d \n", i+1, PrintMove(move), oldNodes)
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("Test complete: %d nodes visited in time %v\n", leafNodes, elapsedTime)
	return nil
}
