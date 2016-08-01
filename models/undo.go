package models

import "go-chess/constants"

// SUndo is the structure of Undo when a step is undone
type SUndo struct {
	Move       int
	CastlePerm int
	EnPas      int
	FiftyMove  int
	PosKey     constants.U64
}
