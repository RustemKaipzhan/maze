package gameflow

import (
	"fmt"
	"maze/data"
	"strconv"
)

var gameData *data.Data

func isValid(message string) bool {
	return message == "\033[1;32m✔\033[0m"
}

func StartGame() {
	clearTerminal()
	message, row, column, allowCustomIcons := "message", 0, 0, false
	promptGridSize(&message, &row, &column)

	gameData = data.NewData(row, column)

	addGridSizeDetail()

	clearTerminal()

	promptCustomIcons(&message, &allowCustomIcons)

	if allowCustomIcons {
		setIcons(&message)
	}

	fmt.Println("gameData: ", gameData)
}

func setIcons(message *string) {
	iconNames := map[int]string{0: "wall", 2: "player", 3: "award"}
	icon := ""
	indexes := getSortedIconIndexes()

	for _, idx := range indexes {
		if idx == 1 {
			continue
		}

		clearTerminal()
		updateIconDetail()
		promptIcon(message, &icon,
			gameDetails+"\n"+enter+iconNames[idx]+" icon"+guideSkip)
		gameData.SetIcon(strconv.Itoa(idx), icon)
	}

	clearTerminal()
	updateIconDetail()
	fmt.Println(gameDetails)
}
