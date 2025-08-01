package cli

import "fmt"

func ParseMove(move string) (fromRow, fromCol, toRow, toCol int, err error) {
	if len(move) != 5 {
		return 0, 0, 0, 0, fmt.Errorf("invalid move format")
	}
	fromCol = int(move[0] - 'a')
	fromRow = 8 - int(move[1]-'0')
	toCol = int(move[3] - 'a')
	toRow = 8 - int(move[4]-'0')
	return
}
