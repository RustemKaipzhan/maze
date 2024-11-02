package gameflow

import "sort"

// Return sorted indexes of icons
func getSortedIconIndexes() []int {
	icons := gameData.GetIcons()
	indexes := make([]int, 0, len(icons))
	for idx := range icons {
		indexes = append(indexes, idx)
	}

	sort.Ints(indexes)
	return indexes
}
