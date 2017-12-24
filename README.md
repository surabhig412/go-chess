# go-chess

A chess engine implemented in go language.

## Installation and running

```
go get -u github.com/surabhig412/go-chess
go build
./go-chess
```

## Play chess

You can play chess with the following rules:
* Press `q` to quit the game
* Press `t` to take back last move
* Press `p` to find perft testing solutions at depth 3
* Give move in algebraic notation. Example: b7a8q - This move will move the piece from square b7 to a8 and will promote the piece to queen(q). Promoting piece is optional.
* Board will be printed after each move. It prints all the necessary information about the current state of the game such as the following:
  * `side` - w(white) or b(black) depending on the side to play
  * `enPas` - If enPas rule applies, it prints the square.
  * `castle` - It prints the castling permissions still available.
