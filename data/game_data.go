package data

type Data struct {
	wall   rune
	player rune
	award  rune

	// Player position
	playerPosX int
	playerPosY int

	// Derived dimensions based on size
	row, column int

	cells [][]cell // 2D slice representing game field
	size  int
}

type cell struct {
	icon rune
}

func NewData(size int) *Data {
	cells := make([][]cell, size)
	for idx := range cells {
		cells[idx] = make([]cell, size)
	}

	return &Data{
		wall:       'X',
		player:     '>',
		award:      '@',
		playerPosX: 0,
		playerPosY: 0,
		row:        size * 3,
		column:     size * 8,
		cells:      cells,
		size:       size,
	}
}
