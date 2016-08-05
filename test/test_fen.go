package test

import (
	"fmt"
	"go-chess/board"
	"go-chess/engine"
	"go-chess/models"
)

// ParseFENTesting tests parsing of FEN string
func ParseFENTesting() {
	var b models.SBoard

	// Position of pieces
	fmt.Println("\nInitial board structure: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	err := engine.ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	board.PrintBoard(&b)

	fmt.Println("\n\nNumbers represent number of blank squares: 4kbnr/5p2/8/8/8/8/PPPP4/RNB2BNR w KQkq a4 0 1")
	err = engine.ParseFEN("4kbnr/5p2/8/8/8/8/PPPP4/RNB2BNR w KQkq a4 0 1", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	board.PrintBoard(&b)

	fmt.Println("\n\nNumber should be 1-8: rnbqkbnr/pppppppp/0/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a1 0 1")
	err = engine.ParseFEN("rnbqkbnr/pppppppp/0/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a1 0 1", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("\n\nInitial board structure with a missing /: rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	err = engine.ParseFEN("rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	board.PrintBoard(&b)

	// Side
	fmt.Println("\n\nInitial board structure with missing side to play: rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR ")
	err = engine.ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR ", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	board.PrintBoard(&b)

	fmt.Println("\n\nError when given invalid side: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR s KQkq a4 0 1")
	err = engine.ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR s KQkq a4 0 1", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Castling
	fmt.Println("\n\nInitial board structure with all castling permissions: rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	err = engine.ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Castling permissions: ", b.CastlePerm)

	fmt.Println("\n\nInitial board structure with few castling permissions: rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQ - 0 1")
	err = engine.ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQ - 0 1", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Castling permissions: ", b.CastlePerm)

	fmt.Println("\n\nMissing castling condition: rnbqkbnr/pppppppp/0/8/8/8/PPPPPPPP/RNBQKBNR b ")
	err = engine.ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b ", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Castling permissions: ", b.CastlePerm)

	// EnPas
	fmt.Println("\n\nChecking enPas when nil: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")
	err = engine.ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("EnPas Condition: ", b.EnPas)

	fmt.Println("\n\nChecking enPas: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a1 0 1")
	err = engine.ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a1 0 1", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("EnPas Condition: ", b.EnPas)

	fmt.Println("\n\nMissing enPas condition: rnbqkbnr/pppppppp/0/8/8/8/PPPPPPPP/RNBQKBNR b KQkq ")
	err = engine.ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq ", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("EnPas Condition: ", b.EnPas)

	fmt.Println("\nAfter first move: rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1")
	err = engine.ParseFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	board.PrintBoard(&b)

	fmt.Println("\nAfter second move: rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2")
	err = engine.ParseFEN("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	board.PrintBoard(&b)

	fmt.Println("\nAfter third move: rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")
	err = engine.ParseFEN("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2", &b)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	board.PrintBoard(&b)

}
