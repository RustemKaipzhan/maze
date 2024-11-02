package gameflow

import (
	"fmt"
	"maze/data"
	"strconv"
)

var gameData *data.Data

func isValid(message string) bool {
	return message == "\033[1;32m✔ Success\033[0m"
}

func StartGame() {
	message, row, column, allowCustomIcons := "message", 0, 0, false
	promptGridSize(&message, &row, &column)

	gameData = data.NewData(row, column)

	promptResponse(&message, &row, &column, &allowCustomIcons)

	if allowCustomIcons {
		setIcons(&message)
	}

	fmt.Println("gameData: ", gameData)
}

func setIcons(message *string) {
	iconNames := map[int]string{0: "wall", 2: "player", 3: "award"}
	icon := ""
	icons := gameData.GetIcons()
	indexes := getSortedIconIndexes()

	for _, idx := range indexes {
		if idx == 1 {
			continue
		}

		clearTerminal()

		icons = gameData.GetIcons()
		promptIcon(message, &icon, icons,
			currentIcons(icons)+"\n\033[1;34m→ Enter: "+iconNames[idx]+" icon \033[0m(type '/' to skip): ")
		gameData.SetIcon(strconv.Itoa(idx), icon)
	}

	clearTerminal()
	fmt.Println("Final", currentIcons(gameData.GetIcons()))
}
