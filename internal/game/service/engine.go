package service

import (
	"fmt"
	"go-chess/internal/game/domain"
)

func MovePiece(b *domain.Board, fromRow, fromCol, toRow, toCol int) error {
	if fromRow < 0 || fromRow > 7 || fromCol < 0 || fromCol > 7 || toRow < 0 || toRow > 7 || toCol < 0 || toCol > 7 {
		return fmt.Errorf("move out of bounds")
	}

	piece := b.Cells[fromRow][fromCol]
	if piece == nil {
		return fmt.Errorf("no piece at the source position")
	}

	if b.Cells[toRow][toCol] != nil && b.Cells[toRow][toCol].Color == piece.Color {
		return fmt.Errorf("cannot capture own piece")
	}

	if piece.Type == "P" {
		if !IsValidPawnMove(b, fromRow, fromCol, toRow, toCol, piece.Color) {
			return fmt.Errorf("invalid pawn move")
		}
	}

	b.Cells[toRow][toCol] = piece
	b.Cells[fromRow][fromCol] = nil
	return nil
}
