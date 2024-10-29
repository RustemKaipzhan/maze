package data

type Data struct {
	icons []icon // key: name (wall player award), value: symbol

	// Player position
	playerPosX int
	playerPosY int

	// Derived dimensions based on size
	row, column int

	cells [][]cell // 2D slice representing game field
}

type cell struct {
	icon icon
}
type icon struct {
	name  string
	value rune
}

func NewData(row, column int) *Data {
	cells := make([][]cell, row)
	for idx := range cells {
		cells[idx] = make([]cell, column)
	}

	icons := []icon{
		icon{name: "wall", value: 'X'},
		icon{name: "player", value: '>'},
		icon{name: "award", value: '@'},
	}

	return &Data{
		icons:      icons,
		playerPosX: 0,
		playerPosY: 0,
		row:        row,
		column:     column,
		cells:      cells,
	}
}

func (data *Data) SetIcons(icons map[string]Icon) {
	for name, icon := range icons {
		data.icons[name] = Icon{iconValue}
	}

}
func (data *Data) GetIcons() map[string]Icon {
	return data.icons
}
