package gameflow

import (
	"fmt"
	"maze/data"
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
		promptIcons(&message)
	}

	promptCells(&message)

	fmt.Println("gameData: ", gameData)
}

func updateIcons(guide string) {
	clearTerminal()
	updateIconDetail()
	fmt.Println(gameDetails + guide)
}
