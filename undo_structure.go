package main

// Undo is the structure of Undo when a step is undone
type Undo struct {
	Move       int
	CastlePerm int
	EnPas      int
	FiftyMove  int
	PosKey     uint64
}
