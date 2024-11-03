package gameflow

import (
	"fmt"
	"maze/gameinput"
	"strconv"
)

const (
	enter            = "\033[1;34m→ Enter: "
	guideGridSize    = enter + "grid size (e.g., 3 3):\033[0m "
	guideCustomIcons = "- Would you like to make custom icons ?\n"
	guideResponse    = "\033[1;34m→ Select an option (y/n):\033[0m "
	guideSkip        = "\033[0m (type '/' to skip): "
)

func promptGridSize(message *string, row, column *int) {
	for {
		fmt.Print(guideGridSize)
		*message, *row, *column = gameinput.ReadSize()

		if isValid(*message) {
			break
		}

		fmt.Println(*message)
	}
}

func promptCustomIcons(message *string, allowCustomIcons *bool) {
	for {
		fmt.Print(guideCustomIcons + guideResponse)

		*message, *allowCustomIcons = gameinput.ReadResponse()

		if isValid(*message) {
			break
		}

		fmt.Println(*message)
		wait(1.5, true)
	}
}
func promptIcons(message *string) {
	iconNames := map[int]string{0: "wall", 2: "player", 3: "award"}
	icon := ""
	indexes := getSortedIconIndexes()

	for _, idx := range indexes {
		promptIcon(message, &icon, enter+"\033[1;4;34m"+iconNames[idx]+"\033[0;1;34m"+" icon"+guideSkip)
		gameData.SetIcon(strconv.Itoa(idx), icon)
	}

	updateIcons("")
}

func promptIcon(message, icon *string, guide string) {
	for {
		updateIcons(guide)
		*message, *icon = gameinput.ReadIcon(gameData.GetIcons())

		if isValid(*message) {
			break
		}

		fmt.Println(*message)
		wait(1.5, true)
	}
}

func promptCells(message *string) {
	cells := gameData.GetCells()
	cellTypes := gameData.GetCellTypes()
	cellsLine := ""
	for i := range cells {
		promptCellsLine(message, &cellsLine, len(cells[i]))
		for j := range cells[i] {
			cells[i][j] = cellTypes[string(cellsLine[j])]
		}
	}
}
func promptCellsLine(message, cellsLine *string, length int) {
	for {
		fmt.Println(gameDetails + enter + "cells (0, 2, 3):\n" + "wall -> 0\nplayer -> 2\naward -> 3\033[0m")
		*message, *cellsLine = gameinput.ReadCellsLine()
		if len(*cellsLine) < length || len(*cellsLine) > length {
			*message = "\033[33m⚠ Warning:\033[0m must be" + strconv.Itoa(length) + " characters\n"
		}

		if isValid(*message) {
			break
		}
		fmt.Println(*message)
		wait(1.5, true)
	}
}
