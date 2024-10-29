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
}

type cell struct {
	icon rune
}

func NewData(row, column int, customIcons string) *Data {
	cells := make([][]cell, row)
	for idx := range cells {
		cells[idx] = make([]cell, column)
	}

	return &Data{
		wall:       rune(customIcons[0]),
		player:     rune(customIcons[1]),
		award:      rune(customIcons[2]),
		playerPosX: 0,
		playerPosY: 0,
		row:        row,
		column:     column,
		cells:      cells,
	}
}
