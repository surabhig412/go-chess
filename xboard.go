package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// threeFoldRepetition find out how many times current position is repeated in the game
func threeFoldRepetition(pos *Board) (r int) {
	for i := 0; i < pos.HisPly; i++ {
		if pos.History[i].PosKey == pos.PosKey {
			r++
		}
	}
	return
}

// drawMaterial suggests if the game is in draw in the current state of the game
func drawMaterial(pos *Board) bool {
	if pos.PceNum[Wp] > 0 || pos.PceNum[Bp] > 0 {
		return false
	}
	if pos.PceNum[Wq] > 0 || pos.PceNum[Bq] > 0 || pos.PceNum[Wr] > 0 || pos.PceNum[Br] > 0 {
		return false
	}
	if pos.PceNum[Wb] > 1 || pos.PceNum[Bb] > 1 {
		return false
	}
	if pos.PceNum[Wn] > 1 || pos.PceNum[Bn] > 1 {
		return false
	}
	if pos.PceNum[Wn] > 0 || pos.PceNum[Wb] > 0 {
		return false
	}
	if pos.PceNum[Bn] > 0 || pos.PceNum[Bb] > 0 {
		return false
	}
	return true
}

// checkResult checks whether game is in draw, stale mate or mate position after each move in the game
func checkResult(pos *Board) bool {
	if pos.FiftyMove > 100 {
		fmt.Printf("1/2-1/2 {fifty move rule (claimed by go-chess)}\n")
		return true
	}
	if threeFoldRepetition(pos) >= 2 {
		fmt.Printf("1/2-1/2 {3-fold repetition (claimed by go-chess)}\n")
		return true
	}
	if drawMaterial(pos) {
		fmt.Printf("1/2-1/2 {insufficient material (claimed by go-chess)}\n")
		return true
	}
	var list MoveList
	found := false
	(&list).GenerateAllMoves(pos)
	for i := 0; i < list.count; i++ {
		if moveMade, _ := MakeMove(list.moves[i].move, pos); !moveMade {
			continue
		}
		found = true
		TakeMove(pos)
		break
	}
	if found {
		return false
	}
	if attacked, _ := SqAttacked(pos.KingSq[pos.Side], pos.Side^1, pos); attacked {
		if pos.Side == White {
			fmt.Printf("0-1 {black mates (claimed by go-chess)}\n")
		} else {
			fmt.Printf("0-1 {white mates (claimed by go-chess)}\n")
		}
		return true
	} else {
		fmt.Printf("\n1/2-1/2 {stalemate (claimed by go-chess)}\n")
		return true
	}
	return false
}

func XboardLoop(pos *Board, info *SearchInfo) {
	depth := -1
	mps := 0
	movestogo := []int{30, 30}
	engineSide := Both
	var movetime time.Duration
	var duration time.Duration
	var inc time.Duration

	for true {
		if pos.Side == engineSide && !checkResult(pos) {
			info.startTime = time.Now()
			info.depth = depth
			if duration.String() != "0s" {
				info.timeSet = true
				duration, _ = time.ParseDuration(fmt.Sprintf("%vms", duration.Seconds()/float64(movestogo[pos.Side])*1000-50))
				info.stopTime = info.startTime.Add(duration).Add(inc)
			}
			if depth == -1 || depth > MaxDepth {
				info.depth = MaxDepth
			}
			fmt.Printf("time: %s start: %s stop: %s depth: %d timeset: %v movestogo: %d movetime: %s\n", duration.String(), info.startTime.String(), info.stopTime.String(), info.depth, info.timeSet, movestogo[pos.Side], movetime.String())
			SearchPosition(pos, info)

			if mps != 0 {
				movestogo[pos.Side^1]--
				if movestogo[pos.Side^1] < 1 {
					movestogo[pos.Side^1] = mps
				}
			}
		}
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.Trim(command, " ")
		if command == "quit" {
			break
		}
		if command == "force" {
			engineSide = Both
			continue
		}
		if command == "protover" {
			fmt.Println("feature ping=1 setboard=1 colors=0 usermove=1")
			fmt.Println("feature done=1")
			continue
		}
		if strings.Contains(command, "sd") {
			parseArr := strings.Split(command, "sd ")
			parseArr = strings.Split(parseArr[1], " ")
			depth, _ = strconv.Atoi(parseArr[0])
			continue
		}
		if strings.Contains(command, "st") {
			parseArr := strings.Split(command, "st ")
			parseArr = strings.Split(parseArr[1], " ")
			movetime, _ = time.ParseDuration(fmt.Sprintf("%sms", parseArr[0]))
			continue
		}
		if strings.Contains(command, "ping") {
			parseArr := strings.Split(command, "ping ")
			parseArr = strings.Split(parseArr[1], " ")
			fmt.Printf("pong%s\n", parseArr[0])
			continue
		}
		if strings.Contains(command, "new") {
			engineSide = Black
			ParseFEN(StartFEN, pos)
			depth = -1
			continue
		}
		if strings.Contains(command, "setboard") {
			engineSide = Both
			parseArr := strings.Split(command, "setboard ")
			ParseFEN(parseArr[1], pos)
			continue
		}
		if strings.Contains(command, "go") {
			engineSide = pos.Side
			continue
		}
		if strings.Contains(command, "usermove") {
			movestogo[pos.Side]--
			parseArr := strings.Split(command, "setboard ")
			move, _ := ParseMove(strings.Trim(parseArr[1], " "), pos)
			if move != NoMove {
				MakeMove(move, pos)
				pos.Ply = 0
			} else {
				continue
			}
		}
	}

}
