package main

import (
	"bufio"
	"fmt"
	"os"
)

type Piece struct {
	Color string // "w" or "b" for white or black
	Type  string // "P", "R", "N", "B", "Q", "K" for Pawn, Rook, Knight, Bishop, Queen, King
}

type Board struct {
	Cells [8][8]*Piece // 8x8 chess board
}

func initBoard() Board {
	var board Board

	// set up pawns
	for i := 0; i < 8; i++ {
		board.Cells[1][i] = &Piece{Color: "w", Type: "P"} // White pawns
		board.Cells[6][i] = &Piece{Color: "b", Type: "P"} // Black pawns
	}

	// Set up major pieces
	backRank := []string{"R", "N", "B", "Q", "K", "B", "N", "R"}
	for i, piece := range backRank {
		board.Cells[0][i] = &Piece{Color: "w", Type: piece} // White pieces
		board.Cells[7][i] = &Piece{Color: "b", Type: piece} // Black pieces
	}

	return board
}

func printBoard(b Board) {
	for i := range 8 {
		fmt.Printf("%d ", 8-i) // Print row numbers
		for j := range 8 {
			if b.Cells[i][j] == nil {
				fmt.Print(". ") // Empty cell
			} else {
				fmt.Printf("%s%s ", b.Cells[i][j].Color, b.Cells[i][j].Type) // Print piece
			}
		}
		fmt.Println()
	}
	fmt.Println("  a b c d e f g h ") // Print column labels
}

func parseMove(move string) (fromRow, fromCol, toRow, toCol int, err error) {
	if len(move) != 5 {
		return 0, 0, 0, 0, fmt.Errorf("invalid move format")
	}

	fromCol = int(move[0] - 'a')
	fromRow = 8 - int(move[1]-'0') // Convert to 0-indexed row
	toCol = int(move[3] - 'a')
	toRow = 8 - int(move[4]-'0') // Convert to 0-indexed row
	return

}

func movePiece(b *Board, fromRow, fromCol, toRow, toCol int) error {
	if fromRow < 0 || fromRow > 7 || fromCol < 0 || fromCol > 7 || toRow < 0 || toRow > 7 || toCol < 0 || toCol > 7 {
		return fmt.Errorf("move out of bounds")
	}

	piece := b.Cells[fromRow][fromCol]
	if piece == nil {
		return fmt.Errorf("no piece at source")
	}

	// Simple rule: can't capture your own piece
	if b.Cells[toRow][toCol] != nil && b.Cells[toRow][toCol].Color == piece.Color {
		return fmt.Errorf("cannot capture your own piece")
	}

	b.Cells[toRow][toCol] = piece
	b.Cells[fromRow][fromCol] = nil
	return nil
}

func main() {
	board := initBoard()
	scanner := bufio.NewScanner(os.Stdin)
	currentPlayer := "w"

	for {
		printBoard(board)
		fmt.Printf("%s's move: ", currentPlayer)
		scanner.Scan()
		move := scanner.Text()

		fromRow, fromCol, toRow, toCol, err := parseMove(move)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		piece := board.Cells[fromRow][fromCol]
		if piece == nil || piece.Color != currentPlayer {
			fmt.Println("Invalid move: no piece or wrong color")
			continue
		}

		err = movePiece(&board, fromRow, fromCol, toRow, toCol)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		currentPlayer = togglePlayer(currentPlayer)
	}

}

func togglePlayer(p string) string {
	if p == "w" {
		return "b"
	}
	return "w"
}
