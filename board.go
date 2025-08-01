package main

type Piece struct {
    Color string // "w" or "b" for white or black
    Type  string // "P", "R", "N", "B", "Q", "K" for Pawn, Rook, Knight, Bishop, Queen, King
}

type Board struct {
    Cells [8][8]*Piece // 8x8 chess board
}