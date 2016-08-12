package main

// SMove defines the structure of each move
type SMove struct {
	move  int
	score int
}

// SMoveList is the list of all moves
type SMoveList struct {
	moves [MaxPositionMoves]SMove
	count int
}

// AddQuietMove to move list
func (list *SMoveList) AddQuietMove(pos *SBoard, move int) {
	list.moves[list.count].move = move
	list.moves[list.count].score = 0
	list.count++
}

// AddCaptureMove to move list
func (list *SMoveList) AddCaptureMove(pos *SBoard, move int) {
	list.moves[list.count].move = move
	list.moves[list.count].score = 0
	list.count++
}

// AddEnPassantMove to move list
func (list *SMoveList) AddEnPassantMove(pos *SBoard, move int) {
	list.moves[list.count].move = move
	list.moves[list.count].score = 0
	list.count++
}

// GenerateAllMoves will generate all possible moves of board
func (list *SMoveList) GenerateAllMoves(pos *SBoard) {
	list.count = 0
	//TODO
}
