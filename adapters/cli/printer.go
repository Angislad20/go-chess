package cli

import (
	"fmt"
	"go-chess/internal/game/domain"
)

func PrintBoard(b domain.Board) {
	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", 8-i)
		for j := 0; j < 8; j++ {
			p := b.Cells[i][j]
			if p == nil {
				fmt.Print(". ")
			} else {
				fmt.Printf("%s%s ", p.Color, p.Type)
			}
		}
		fmt.Println()
	}
	fmt.Println("  a b c d e f g h")
}
