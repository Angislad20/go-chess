package service

import "go-chess/internal/game/domain"

func IsValidPawnMove(b *domain.Board, fromRow, frowCol, toRow, toCol int, color string) bool {
	dir := 1
	startRow := 1

	if color == "b" {
		dir = -1
		startRow = 6
	}

	// Moving forward
	if toCol == frowCol && toRow == fromRow+dir && b.Cells[toRow][toCol] == nil {
		return true
	}

	// Moving two squares forward from the starting position
	if fromRow == startRow && toCol == frowCol && toRow == fromRow+2*dir && b.Cells[toRow][toCol] == nil && b.Cells[fromRow+dir][frowCol] == nil {
		return true
	}

	// Capturing diagonally
	if (toCol == frowCol-1 || toCol == frowCol+1) && toRow == fromRow+dir {
		target := b.Cells[toRow][toCol]
		return target != nil && target.Color != color
	}

	return false
}
