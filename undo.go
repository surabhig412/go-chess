package main

// SUndo is the structure of Undo when a step is undone
type SUndo struct {
	move       int
	castlePerm int
	enPas      int
	fiftyMove  int
	posKey     U64
}
