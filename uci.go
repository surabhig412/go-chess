package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// UciLoop implements UCI(Universal Chess Interface) protocol
func UciLoop() {
	fmt.Printf("id name Surabhi\n")
	fmt.Printf("id author go-chess\n")
	fmt.Printf("uciok\n")
	var board Board
	var info SearchInfo
	(&board).PvTable.Init()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "isready") {
			fmt.Printf("readyok\n")
		} else if strings.HasPrefix(line, "position") {
			parsePosition(line, &board)
		} else if strings.HasPrefix(line, "ucinewgame") {
			parsePosition("position startpos\n", &board)
		} else if strings.HasPrefix(line, "go") {
			parseGo(line, &info, &board)
		} else if strings.HasPrefix(line, "quit") {
			info.quit = true
			break
		} else if strings.HasPrefix(line, "uci") {
			fmt.Printf("id name Surabhi\n")
			fmt.Printf("id author go-chess\n")
			fmt.Printf("uciok\n")
		}
		if info.quit {
			break
		}
	}
}

/* parsePosition parses position command. Types of position command:
position startpos
position fen fenstr
... moves e2e4 e7e5 b7b8q(optional)
*/
func parsePosition(lineIn string, pos *Board) {
	parseArr := strings.Split(lineIn, "position ")
	if strings.HasPrefix(parseArr[1], "fen") {
		parseArr = strings.Split(parseArr[1], "fen ")
		ParseFEN(parseArr[1], pos)
	} else {
		ParseFEN(StartFEN, pos)
	}
	if strings.Contains(lineIn, "moves") {
		parseArr = strings.Split(lineIn, "moves ")
		movesArr := strings.Split(parseArr[1], " ")
		for i := 0; i < len(movesArr); i++ {
			fmt.Println(movesArr[i])
			move, _ := ParseMove(movesArr[i], pos)
			if move == NoMove {
				break
			}
			MakeMove(move, pos)
			pos.Ply = 0
		}
	}
	pos.Print()
}

/* parseGo parses go command. Type of go command:
go depth 6 wtime 180000 btime 100000 binc 1000 winc 1000 movetime 1000 movestogo 40
*/
func parseGo(line string, info *SearchInfo, pos *Board) {
	if info.stopped {
		info.stopped = false
	}
	depth := -1
	movestogo := 30
	var movetime time.Duration
	var duration time.Duration
	var inc time.Duration
	info.timeSet = false
	if strings.Contains(line, "binc") && pos.Side == Black {
		parseArr := strings.Split(line, "binc ")
		parseArr = strings.Split(parseArr[1], " ")
		inc, _ = time.ParseDuration(fmt.Sprintf("%sms", parseArr[0]))
	}
	if strings.Contains(line, "winc") && pos.Side == White {
		parseArr := strings.Split(line, "winc ")
		parseArr = strings.Split(parseArr[1], " ")
		inc, _ = time.ParseDuration(fmt.Sprintf("%sms", parseArr[0]))
	}
	if strings.Contains(line, "wtime") && pos.Side == White {
		parseArr := strings.Split(line, "wtime ")
		parseArr = strings.Split(parseArr[1], " ")
		duration, _ = time.ParseDuration(fmt.Sprintf("%sms", parseArr[0]))
	}
	if strings.Contains(line, "btime") && pos.Side == Black {
		parseArr := strings.Split(line, "btime ")
		parseArr = strings.Split(parseArr[1], " ")
		duration, _ = time.ParseDuration(fmt.Sprintf("%sms", parseArr[0]))
	}
	if strings.Contains(line, "movestogo") {
		parseArr := strings.Split(line, "movestogo ")
		parseArr = strings.Split(parseArr[1], " ")
		movestogo, _ = strconv.Atoi(parseArr[0])
	}
	if strings.Contains(line, "movetime") {
		parseArr := strings.Split(line, "movetime ")
		parseArr = strings.Split(parseArr[1], " ")
		movetime, _ = time.ParseDuration(fmt.Sprintf("%sms", parseArr[0]))
	}
	if strings.Contains(line, "depth") {
		parseArr := strings.Split(line, "depth ")
		parseArr = strings.Split(parseArr[1], " ")
		depth, _ = strconv.Atoi(parseArr[0])
	}

	if movetime.String() != "0s" {
		duration = movetime
		movestogo = 1
	}
	info.startTime = time.Now()
	if duration.String() != "0s" {
		info.timeSet = true
		duration, _ = time.ParseDuration(fmt.Sprintf("%vms", duration.Seconds()/float64(movestogo)*1000-50))
		info.stopTime = info.startTime.Add(duration).Add(inc)
	}
	if depth == -1 {
		info.depth = MaxDepth
	} else {
		info.depth = depth
	}

	fmt.Printf("time: %s start: %s stop: %s depth: %d timeset: %v\n", duration.String(), info.startTime.String(), info.stopTime.String(), info.depth, info.timeSet)
	SearchPosition(pos, info)
}
