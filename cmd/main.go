package main

import (
	"bufio"
	"fmt"
	"go-chess/adapters/cli"
	"go-chess/internal/game/domain"
	"go-chess/internal/game/service"
	"os"
)

func initBoard() domain.Board {
	var board domain.Board
	backRank := []string{"R", "N", "B", "Q", "K", "B", "N", "R"}

	for i := 0; i < 8; i++ {
		board.Cells[1][i] = &domain.Piece{Color: "w", Type: "P"}
		board.Cells[6][i] = &domain.Piece{Color: "b", Type: "P"}
	}

	for i, t := range backRank {
		board.Cells[0][i] = &domain.Piece{Color: "w", Type: t}
		board.Cells[7][i] = &domain.Piece{Color: "b", Type: t}
	}

	return board
}

func main() {
	board := initBoard()
	currentPlayer := "w"
	scanner := bufio.NewScanner(os.Stdin)

	for {
		cli.PrintBoard(board)
		fmt.Printf("%s's move: ", currentPlayer)
		scanner.Scan()
		move := scanner.Text()

		fromRow, fromCol, toRow, toCol, err := cli.ParseMove(move)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		piece := board.Cells[fromRow][fromCol]
		if piece == nil || piece.Color != currentPlayer {
			fmt.Println("Invalid move: no piece or wrong color")
			continue
		}

		err = service.MovePiece(&board, fromRow, fromCol, toRow, toCol)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		currentPlayer = service.TogglePlayer(currentPlayer)
	}
}
