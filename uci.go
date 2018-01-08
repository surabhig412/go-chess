package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

/* parsePosition parse position command. Types of position command:
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

func parseGo(line string, info *SearchInfo, pos *Board) {

}
