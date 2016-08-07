package main

import (
	"errors"
	"log"
	"strconv"
	"unicode/utf8"
)

// ParseFEN parses FEN notation
func ParseFEN(fen string, pos *SBoard) error {
	if len(fen) <= 0 || pos == nil {
		log.Fatalln("FEN or board structure is invalid")
	}
	rank := Rank8
	file := FileA
	piece := 0
	pos.Reset()

	// Parsing position of pieces in FEN notation
	for (rank >= Rank1) && (len(fen) > 0) {
		count := 1
		switch string(fen[0]) {
		case "p":
			piece = Bp
			break
		case "r":
			piece = Br
			break
		case "n":
			piece = Bn
			break
		case "b":
			piece = Bb
			break
		case "k":
			piece = Bk
			break
		case "q":
			piece = Bq
			break
		case "P":
			piece = Wp
			break
		case "R":
			piece = Wr
			break
		case "N":
			piece = Wn
			break
		case "B":
			piece = Wb
			break
		case "K":
			piece = Wk
			break
		case "Q":
			piece = Wq
			break
		case "1":
			fallthrough
		case "2":
			fallthrough
		case "3":
			fallthrough
		case "4":
			fallthrough
		case "5":
			fallthrough
		case "6":
			fallthrough
		case "7":
			fallthrough
		case "8":
			piece = Empty
			count, _ = strconv.Atoi(string(fen[0]))
			break
		case "/":
			fallthrough
		case " ":
			rank--
			file = FileA
			fen = fen[1:len(fen)]
			continue

		default:
			log.Println("Error in parsing FEN")
			pos.Reset()
			return errors.New("Error in parsing FEN")
		}
		for i := 0; i < count; i++ {
			sq64 := (rank * 8) + file
			var sq120 int
			if sq64 >= 0 && sq64 < 64 {
				sq120 = SQ120(sq64)
			} else {
				log.Println("Error in parsing FEN")
				pos.Reset()
				return errors.New("Error in parsing FEN")
			}
			if piece != Empty {
				pos.Pieces[sq120] = piece
			}
			file++
		}
		fen = fen[1:len(fen)]
	}

	// Parsing side to play in FEN notation
	if len(fen) <= 0 {
		log.Println("Missing side in FEN notation")
		pos.Reset()
		return errors.New("Missing side in FEN notation")
	}
	if !((string(fen[0]) == "w") || (string(fen[0]) == "b")) {
		log.Println("Error in parsing FEN, side should be w or b in FEN notation")
		return errors.New("Error in parsing FEN, side should be w or b in FEN notation")
	}
	if string(fen[0]) == "w" {
		pos.Side = White
	} else {
		pos.Side = Black
	}
	fen = fen[2:len(fen)]

	// Parsing castling in FEN notation
	if len(fen) <= 0 {
		log.Println("Missing castling permissions in FEN notation")
		pos.Reset()
		return errors.New("Missing castling permissions in FEN notation")
	}
	for i := 0; i < 4; i++ {
		if string(fen[0]) == " " {
			break
		}
		switch string(fen[0]) {
		case "K":
			pos.CastlePerm |= Wkca
			break
		case "Q":
			pos.CastlePerm |= Wqca
			break
		case "k":
			pos.CastlePerm |= Bkca
			break
		case "q":
			pos.CastlePerm |= Bqca
			break
		default:
			break
		}
		fen = fen[1:len(fen)]
	}
	fen = fen[1:len(fen)]

	// Parsing enPas in FEN notation
	if len(fen) <= 0 {
		log.Println("Missing enPas in FEN notation")
		pos.Reset()
		return errors.New("Missing enPas in FEN notation")
	}
	if string(fen[0]) != "-" {
		fileRune, _ := utf8.DecodeRuneInString(string(fen[0]))
		file = int(fileRune) - 97
		rankConvert, _ := strconv.Atoi(string(fen[1]))
		rank = rankConvert - 1
		if file < FileA || file > FileH {
			log.Println("Error in parsing enPas in FEN notation")
			return errors.New("Error in parsing enPas in FEN notation")
		}
		if rank < Rank1 || rank > Rank8 {
			log.Println("Error in parsing enPas in FEN notation")
			return errors.New("Error in parsing enPas in FEN notation")
		}
		pos.EnPas = FR2SQ(file, rank)
	}

	// Generating PosKey of board structure
	pos.PosKey = GeneratePosKey(pos)
	return nil
}
