package main

type Board struct {
	boxes  [][]Boxe
	height int
	width  int
}

func (b *Board) Content(col, row int) Boxe {
	return b.boxes[row][col]
}

type Bords struct {
	ground  Board
	objects Board
}
