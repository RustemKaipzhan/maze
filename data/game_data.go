package data

import "strconv"

// add only ascii to icons!
type Data struct {
	cellTypes map[string]cell // map for cell types
	// Player position
	playerPosX int
	playerPosY int

	// Derived dimensions based on size
	row, column int

	cells [][]cell // 2D slice representing game field
}

type cell struct {
	name string
	icon string
}

//🤠
var (
	wall   = cell{"wall", "▓"}
	player = cell{"player", "🏃"}
	award  = cell{"award", "🏆"}
	free   = cell{"free", " "}
)

func NewData(row, column int) *Data {
	// Initialize cells
	cells := make([][]cell, row)
	for idx := range cells {
		cells[idx] = make([]cell, column)
	}

	return &Data{
		cellTypes:  map[string]cell{"0": wall, "1": free, "2": player, "3": award},
		playerPosX: 0,
		playerPosY: 0,
		row:        row,
		column:     column,
		cells:      cells,
	}
}

func (data *Data) GetSize() (int, int) {
	return data.row, data.column
}
func (data *Data) GetCells() [][]cell {
	return data.cells
}

func (data *Data) GetCellTypes() map[string]cell {
	return data.cellTypes
}
func (data *Data) SetIcon(idx string, icon string) {
	if icon == "" {
		return
	}

	cell := data.cellTypes[idx]
	cell.icon = icon
	data.cellTypes[idx] = cell
}

func (data *Data) GetIcons() map[int]string {
	newMap := make(map[int]string)
	for idxS, cell := range data.cellTypes {
		if idxS == "1" {
			continue
		}

		idx, error := strconv.Atoi(idxS)
		if error != nil {
		}

		newMap[idx] = cell.icon
	}

	return newMap
}
