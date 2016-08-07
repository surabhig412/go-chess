package main

import "fmt"

// Test tests functions of go-chess
func Test() {
	fmt.Println("\nTesting Boards: ")

	fmt.Println("Board structures:")
	for index := 0; index < BrdSqNum; index++ {
		if index%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", Sq120ToSq64[index])
	}
	fmt.Println()
	fmt.Println()
	for index := 0; index < 64; index++ {
		if index%8 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", Sq64ToSq120[index])
	}

	fmt.Println("\nTesting Bitboards: ")

	fmt.Println("\nCounting and popping of pawns:")
	var playBitBoard Bitboard
	playBitBoard = 0
	playBitBoard |= (Bitboard(1) << Bitboard(SQ64(D2)))
	playBitBoard |= (Bitboard(1) << Bitboard(SQ64(D3)))
	playBitBoard |= (Bitboard(1) << Bitboard(SQ64(D4)))
	playBitBoard.Print()
	fmt.Println("\nCount: ", playBitBoard.Count())
	i := (&playBitBoard).Pop()
	fmt.Println("Index: ", i)
	playBitBoard.Print()
	fmt.Println("\nCount: ", playBitBoard.Count())
	i = (&playBitBoard).Pop()
	fmt.Println("Index: ", i)
	playBitBoard.Print()
	fmt.Println("\nCount: ", playBitBoard.Count())
	i = (&playBitBoard).Pop()
	fmt.Println("Index: ", i)
	playBitBoard.Print()
	fmt.Println("\nCount: ", playBitBoard.Count())

	fmt.Println("Setting and clearing of bits on the chess board:")
	playBitBoard = 0
	(&playBitBoard).Set(61)
	playBitBoard.Print()
	(&playBitBoard).Clear(61)
	playBitBoard.Print()

	fmt.Println("\nTesting Parsing of FEN notation: ")

	var board SBoard

	// Position of pieces
	fmt.Println("\nInitial board structure: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	err := ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	fmt.Println("\n\nNumbers represent number of blank squares: 4kbnr/5p2/8/8/8/8/PPPP4/RNB2BNR w KQkq a4 0 1")
	err = ParseFEN("4kbnr/5p2/8/8/8/8/PPPP4/RNB2BNR w KQkq a4 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	fmt.Println("\n\nNumber should be 1-8: rnbqkbnr/pppppppp/0/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a1 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/0/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a1 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("\n\nInitial board structure with a missing /: rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	err = ParseFEN("rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	// Side
	fmt.Println("\n\nInitial board structure with missing side to play: rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR ")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR ", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	fmt.Println("\n\nError when given invalid side: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR s KQkq a4 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR s KQkq a4 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Castling
	fmt.Println("\n\nInitial board structure with all castling permissions: rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Castling permissions: ", board.CastlePerm)

	fmt.Println("\n\nInitial board structure with few castling permissions: rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQ - 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQ - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Castling permissions: ", board.CastlePerm)

	fmt.Println("\n\nMissing castling condition: rnbqkbnr/pppppppp/0/8/8/8/PPPPPPPP/RNBQKBNR b ")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b ", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Castling permissions: ", board.CastlePerm)

	// EnPas
	fmt.Println("\n\nChecking enPas when nil: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("EnPas Condition: ", board.EnPas)

	fmt.Println("\n\nChecking enPas: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a1 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a1 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("EnPas Condition: ", board.EnPas)

	fmt.Println("\n\nMissing enPas condition: rnbqkbnr/pppppppp/0/8/8/8/PPPPPPPP/RNBQKBNR b KQkq ")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq ", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("EnPas Condition: ", board.EnPas)

	fmt.Println("\nAfter first move: rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	fmt.Println("\nAfter second move: rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2")
	err = ParseFEN("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	fmt.Println("\nAfter third move: rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")
	err = ParseFEN("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	fmt.Println("\nFiles Board:")
	for index := 0; index < BrdSqNum; index++ {
		if index%10 == 0 && index != 0 {
			fmt.Println()
		}
		fmt.Printf("%4d", FilesBrd[index])
	}

	fmt.Println("\nRanks Board:")
	for index := 0; index < BrdSqNum; index++ {
		if index%10 == 0 && index != 0 {
			fmt.Println()
		}
		fmt.Printf("%4d", RanksBrd[index])
	}

	fmt.Println("\nPrinting bitboards of pawns:")
	ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", &board)
	board.Print()
	fmt.Println("\nWhite pawn: ")
	pawnStructure := Bitboard(board.Pawns[White])
	pawnStructure.Print()
	fmt.Println("\nBlack pawn: ")
	pawnStructure = Bitboard(board.Pawns[Black])
	pawnStructure.Print()
	fmt.Println("\nBoth pawns: ")
	pawnStructure = Bitboard(board.Pawns[Both])
	pawnStructure.Print()

	fmt.Println("\nChecking if board structure is correct:")
	ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", &board)
	err = (&board).Check()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Checking successful")
	}

	fmt.Println("\nChecking if board structure is incorrect after forcefully changing piece count:")
	ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", &board)
	board.PceNum[Wp]--
	err = (&board).Check()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Checking successful")
	}

	fmt.Println("\nChecking if board structure is incorrect after forcefully changing posKey:")
	ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", &board)
	board.PosKey ^= SideKey
	err = (&board).Check()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Checking successful")
	}
}
