package main

// Constants for the entire game of chess
const (
	BrdSqNum         = 120
	MaxGameMoves     = 2048
	MaxPositionMoves = 256
	StartFEN         = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

// Possible values for various fields
const (
	PceChar  = ".PNBRQKpnbrqk" // . for empty
	SideChar = "wb-"           // - for both
	RankChar = "12345678"
	FileChar = "abcdefgh"
)

// PieceBig shows whether respective piece of chess is big or not
var PieceBig = [13]int{False, False, True, True, True, True, True, False, True, True, True, True, True}

// PieceMaj shows whether respective piece of chess is major or not
var PieceMaj = [13]int{False, False, False, False, True, True, True, False, False, False, True, True, True}

// PieceMin shows whether respective piece of chess is minor or not
var PieceMin = [13]int{False, False, True, True, False, False, False, False, True, True, False, False, False}

// PieceVal shows values of respective pieces of chess
var PieceVal = [13]int{0, 100, 325, 325, 550, 1000, 50000, 100, 325, 325, 550, 1000, 50000}

// PieceCol shows colour of respective pieces of chess
var PieceCol = [13]int{Both, White, White, White, White, White, White, Black, Black, Black, Black, Black, Black}

// PieceSlides shows if the piece is sliding piece or not
var PieceSlides = [13]int{False, False, False, True, True, True, False, False, False, True, True, True, False}

// LoopSlidePce is defined to loop through white or black sliding pieces
var LoopSlidePce = [8]int{Wb, Wr, Wq, 0, Bb, Br, Bq, 0}

// LoopNonSlidePce is defined to loop through white or black sliding pieces
var LoopNonSlidePce = [6]int{Wn, Wk, 0, Bn, Bk, 0}

// LoopSlideIndex defines index from where black and white sliding pieces start in LoopSlidePce array
var LoopSlideIndex = [2]int{0, 4}

// LoopNonSlideIndex defines index from where black and white non-sliding pieces start in LoopNonSlidePce array
var LoopNonSlideIndex = [2]int{0, 3}

// Sq120ToSq64 is used to convert a square in 120 sq board to a square in 64 sq board
var Sq120ToSq64 [BrdSqNum]int

// Sq64ToSq120 is used to convert a square in 64 sq board to a square in 120 sq board
var Sq64ToSq120 [64]int

// KnDir represents possible directions in which knight can move
var KnDir = [8]int{-8, -19, -21, -12, 8, 19, 21, 12}

// RkDir represents possible directions in which rook can move
var RkDir = [4]int{-1, -10, 1, 10}

// BiDir represents possible directions in which bishop can move
var BiDir = [4]int{-9, -11, 11, 9}

// KiDir represents possible directions in which king can move
var KiDir = [8]int{-1, -10, 1, 10, -9, -11, 11, 9}

// PieceKnight is used to find out if a piece is knight or not
var PieceKnight = [13]int{False, False, True, False, False, False, False, False, True, False, False, False, False}

// PieceKing is used to find out if a piece is king or not
var PieceKing = [13]int{False, False, False, False, False, False, True, False, False, False, False, False, True}

// PieceRookQueen is used to find out if a piece is rook or queen or not
var PieceRookQueen = [13]int{False, False, False, False, True, True, False, False, False, False, True, True, False}

// PieceBishopQueen is used to find out if a piece is bishop or queen or not
var PieceBishopQueen = [13]int{False, False, False, True, False, True, False, False, False, True, False, True, False}

// SetMask sets the bit in 64 sq board
var SetMask [64]uint64

// ClearMask clears the bit in 64 sq board
var ClearMask [64]uint64

// PieceKeys demonstrates a random number of the position of each piece in all squares of the board
var PieceKeys [13][120]uint64

// SideKey is a random number for which side is to play
var SideKey uint64

// CastleKeys represents random number for wK,wQ,bK,bQ side castling
var CastleKeys [16]uint64

// FilesBrd represents the file to which a square belongs to
var FilesBrd [BrdSqNum]int

// RanksBrd represents the rank to which a square belongs to
var RanksBrd [BrdSqNum]int

// Possible flags for representing a move
const (
	MFlagEP   int = 0x40000   // EnPas bit
	MFlagPS   int = 0x80000   // PawnStart move
	MFlagCA   int = 0x1000000 // Castling move
	MFlagCAP  int = 0x7C000   // Captured piece move
	MFlagPROM int = 0xF00000  // Promoted piece move
)

// Possible values for pieces
const (
	Empty int = iota
	Wp
	Wn
	Wb
	Wr
	Wq
	Wk
	Bp
	Bn
	Bb
	Br
	Bq
	Bk
)

// Columns A-H
const (
	FileA int = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
	FileNone
)

// Rows 1-8
const (
	Rank1 int = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
	RankNone
)

// Colors of pieces
const (
	White int = iota
	Black
	Both
)

// Integer values of true, false
const (
	False int = iota
	True
)

// Integer representation of squares of chess
// Rank 1
const (
	A1 int = iota + 21
	B1
	C1
	D1
	E1
	F1
	G1
	H1
)

// Rank 2
const (
	A2 int = iota + 31
	B2
	C2
	D2
	E2
	F2
	G2
	H2
)

// Rank 3
const (
	A3 int = iota + 41
	B3
	C3
	D3
	E3
	F3
	G3
	H3
)

// Rank 4
const (
	A4 int = iota + 51
	B4
	C4
	D4
	E4
	F4
	G4
	H4
)

// Rank 5
const (
	A5 int = iota + 61
	B5
	C5
	D5
	E5
	F5
	G5
	H5
)

// Rank 6
const (
	A6 int = iota + 71
	B6
	C6
	D6
	E6
	F6
	G6
	H6
)

// Rank 7
const (
	A7 int = iota + 81
	B7
	C7
	D7
	E7
	F7
	G7
	H7
)

// Rank 8
const (
	A8 int = iota + 91
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	NoSq
	Offboard
)

// Wkca represents White king castling, rest have similar connotations
const (
	Wkca = 1
	Wqca = 2
	Bkca = 4
	Bqca = 8
)
