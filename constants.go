package main

// Constants for the entire game of chess
const (
	BrdSqNum         = 120
	MaxGameMoves     = 2048
	MaxPositionMoves = 256
	StartFEN         = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	NoMove           = 0
	PvSize           = 0x100000 * 2
	MaxDepth         = 64
)

// Possible values for various fields
const (
	PceChar  = ".PNBRQKpnbrqk" // . for empty
	SideChar = "wb-"           // - for both
	RankChar = "12345678"
	FileChar = "abcdefgh"
)

// PieceBig shows whether respective piece of chess is big or not(except pawns all big)
var PieceBig = [13]int{False, False, True, True, True, True, True, False, True, True, True, True, True}

// PieceMaj shows whether respective piece of chess is major or not(rook, quuen and king are major pieces)
var PieceMaj = [13]int{False, False, False, False, True, True, True, False, False, False, True, True, True}

// PieceMin shows whether respective piece of chess is minor or not(knight and bishop are minor pieces)
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

// PieceDir is direction of each piece
var PieceDir = [13][8]int{
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{-8, -19, -21, -12, 8, 19, 21, 12},
	{-9, -11, 11, 9, 0, 0, 0, 0},
	{-1, -10, 1, 10, 0, 0, 0, 0},
	{-1, -10, 1, 10, -9, -11, 11, 9},
	{-1, -10, 1, 10, -9, -11, 11, 9},
	{0, 0, 0, 0, 0, 0, 0},
	{-8, -19, -21, -12, 8, 19, 21, 12},
	{-9, -11, 11, 9, 0, 0, 0, 0},
	{-1, -10, 1, 10, 0, 0, 0, 0},
	{-1, -10, 1, 10, -9, -11, 11, 9},
	{-1, -10, 1, 10, -9, -11, 11, 9},
}

// PawnTable gives the material score of pawn on respective squares on the board
var PawnTable = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	10, 10, 0, -10, -10, 0, 10, 10,
	5, 0, 0, 5, 5, 0, 0, 5,
	0, 0, 10, 20, 20, 10, 0, 0,
	5, 5, 5, 10, 10, 5, 5, 5,
	10, 10, 10, 20, 20, 10, 10, 10,
	20, 20, 20, 30, 30, 20, 20, 20,
	0, 0, 0, 0, 0, 0, 0, 0}

// KnightTable gives the material score of knight on respective squares on the board
var KnightTable = [64]int{
	0, -10, 0, 0, 0, 0, -10, 0,
	0, 0, 0, 5, 5, 0, 0, 0,
	0, 0, 10, 10, 10, 10, 0, 0,
	0, 0, 10, 20, 20, 10, 5, 0,
	5, 10, 15, 20, 20, 15, 10, 5,
	5, 10, 10, 20, 20, 10, 10, 5,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0}

// BishopTable gives the material score of bishop on respective squares on the board
var BishopTable = [64]int{
	0, 0, -10, 0, 0, -10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0}

// RookTable gives the material score of rook on respective squares on the board
var RookTable = [64]int{
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	25, 25, 25, 25, 25, 25, 25, 25,
	0, 0, 5, 10, 10, 5, 0, 0}

// Mirror64 gives repective square for the black side
var Mirror64 = [64]int{
	56, 57, 58, 59, 60, 61, 62, 63,
	48, 49, 50, 51, 52, 53, 54, 55,
	40, 41, 42, 43, 44, 45, 46, 47,
	32, 33, 34, 35, 36, 37, 38, 39,
	24, 25, 26, 27, 28, 29, 30, 31,
	16, 17, 18, 19, 20, 21, 22, 23,
	8, 9, 10, 11, 12, 13, 14, 15,
	0, 1, 2, 3, 4, 5, 6, 7}

// NumDir specifies number of directions in which a piece can move
var NumDir = [13]int{0, 0, 8, 4, 4, 8, 8, 0, 8, 4, 4, 8, 8}

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

// PiecePawn is used to find out if a piece is pawn or not
var PiecePawn = [13]int{False, True, False, False, False, False, False, True, False, False, False, False, False}

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

// CastlePerm gives castling permission to each square on the board
var CastlePerm [BrdSqNum]int

// Possible flags for representing a move
const (
	MFlagEP   int = 0x40000   // EnPas bit
	MFlagPS   int = 0x80000   // PawnStart move
	MFlagCA   int = 0x1000000 // Castling move
	MFlagCAP  int = 0x7C000   // Captured piece move(including ep capture)
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
