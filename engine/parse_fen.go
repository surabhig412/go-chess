package engine

import (
	"errors"
	"go-chess/constants"
	"go-chess/models"
	"go-chess/utils"
	"log"
	"strconv"
	"unicode/utf8"
)

// ParseFEN parses FEN notation
func ParseFEN(fen string, pos *models.SBoard) error {
	if len(fen) <= 0 || pos == nil {
		log.Fatalln("FEN or board structure is invalid")
	}
	rank := constants.Rank8
	file := constants.FileA
	piece := 0
	ResetBoard(pos)

	// Parsing position of pieces in FEN notation
	for (rank >= constants.Rank1) && (len(fen) > 0) {
		count := 1
		switch string(fen[0]) {
		case "p":
			piece = constants.Bp
			break
		case "r":
			piece = constants.Br
			break
		case "n":
			piece = constants.Bn
			break
		case "b":
			piece = constants.Bb
			break
		case "k":
			piece = constants.Bk
			break
		case "q":
			piece = constants.Bq
			break
		case "P":
			piece = constants.Wp
			break
		case "R":
			piece = constants.Wr
			break
		case "N":
			piece = constants.Wn
			break
		case "B":
			piece = constants.Wb
			break
		case "K":
			piece = constants.Wk
			break
		case "Q":
			piece = constants.Wq
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
			piece = constants.Empty
			count, _ = strconv.Atoi(string(fen[0]))
			break
		case "/":
			fallthrough
		case " ":
			rank--
			file = constants.FileA
			fen = fen[1:len(fen)]
			continue

		default:
			log.Println("Error in parsing FEN")
			ResetBoard(pos)
			return errors.New("Error in parsing FEN")
		}
		for i := 0; i < count; i++ {
			sq64 := (rank * 8) + file
			var sq120 int
			if sq64 >= 0 && sq64 < 64 {
				sq120 = utils.SQ120(sq64)
			} else {
				log.Println("Error in parsing FEN")
				ResetBoard(pos)
				return errors.New("Error in parsing FEN")
			}
			if piece != constants.Empty {
				pos.Pieces[sq120] = piece
			}
			file++
		}
		fen = fen[1:len(fen)]
	}

	// Parsing side to play in FEN notation
	if len(fen) <= 0 {
		log.Println("Missing side in FEN notation")
		ResetBoard(pos)
		return errors.New("Missing side in FEN notation")
	}
	if !((string(fen[0]) == "w") || (string(fen[0]) == "b")) {
		log.Println("Error in parsing FEN, side should be w or b in FEN notation")
		return errors.New("Error in parsing FEN, side should be w or b in FEN notation")
	}
	if string(fen[0]) == "w" {
		pos.Side = constants.White
	} else {
		pos.Side = constants.Black
	}
	fen = fen[2:len(fen)]

	// Parsing castling in FEN notation
	if len(fen) <= 0 {
		log.Println("Missing castling permissions in FEN notation")
		ResetBoard(pos)
		return errors.New("Missing castling permissions in FEN notation")
	}
	for i := 0; i < 4; i++ {
		if string(fen[0]) == " " {
			break
		}
		switch string(fen[0]) {
		case "K":
			pos.CastlePerm |= constants.Wkca
			break
		case "Q":
			pos.CastlePerm |= constants.Wqca
			break
		case "k":
			pos.CastlePerm |= constants.Bkca
			break
		case "q":
			pos.CastlePerm |= constants.Bqca
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
		ResetBoard(pos)
		return errors.New("Missing enPas in FEN notation")
	}
	if string(fen[0]) != "-" {
		fileRune, _ := utf8.DecodeRuneInString(string(fen[0]))
		file = int(fileRune) - 97
		rankConvert, _ := strconv.Atoi(string(fen[1]))
		rank = rankConvert - 1
		if file < constants.FileA || file > constants.FileH {
			log.Println("Error in parsing enPas in FEN notation")
			return errors.New("Error in parsing enPas in FEN notation")
		}
		if rank < constants.Rank1 || rank > constants.Rank8 {
			log.Println("Error in parsing enPas in FEN notation")
			return errors.New("Error in parsing enPas in FEN notation")
		}
		pos.EnPas = utils.FR2SQ(file, rank)
	}

	// Generating PosKey of board structure
	pos.PosKey = utils.GeneratePosKey(pos)
	return nil
}
