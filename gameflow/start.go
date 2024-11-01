package gameflow

import (
	"fmt"
	"maze/data"
	"sort"
	"strconv"
)

func isValid(message string) bool {
	return message == "\033[1;32m✔ Success\033[0m"
}
func StartGame() {
	message, row, column, allowCustomIcons := "message", 0, 0, false

	promtGridSize(&message, &row, &column)

	gameData := data.NewData(row, column)

	promtResponse(&message, &row, &column, &allowCustomIcons)

	if allowCustomIcons {
		setIcons(gameData, &message)
	}

	fmt.Println("gameData: ", gameData)
}

func setIcons(gameData *data.Data, message *string) {
	iconNames := map[int]string{0: "wall", 2: "player", 3: "award"}
	icons := gameData.GetIcons()
	icon := ""
	keys := make([]int, 0)

	for k := range icons {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, idx := range keys {
		if idx == 1 {
			continue
		}
		clearTerminal()

		icons = gameData.GetIcons()
		promtIcon(message, &icon, &icons,
			currentIcons(&icons)+"\nEnter \033[1m"+iconNames[idx]+"\033[0m icon (type '/' to skip): ")
		gameData.SetIcon(strconv.Itoa(idx), icon)
	}

	clearTerminal()
	icons = gameData.GetIcons()
	fmt.Println(currentIcons(&icons))
}

func currentIcons(icons *map[int]string) string {
	return "Game icons:" +
		"\n  1. wall - " + (*icons)[0] +
		"\n  2. player - " + (*icons)[2] +
		"\n  3. award - " + (*icons)[3]
}
