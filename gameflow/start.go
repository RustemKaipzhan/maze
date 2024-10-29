package gameflow

// "fmt"
import (
	"fmt"
	"maze/data"
	"maze/gameinput"
)

func isValid(message string) bool {
	if message == "success" {
		return true
	}

	return false
}
func StartGame() {
	message, row, column, allowCustIcons := "message", 0, 0, false
	gameIcons := "X>@" // wall player award

	for {
		fmt.Print("Enter grid size (row column e.g., 4 7): ")
		message, row, column = gameinput.ReadSize()

		if isValid(message) {
			break
		}

		fmt.Println(message)
	}

	for {
		fmt.Printf("Grid size %d x %d", row, column)
		fmt.Println("Would you like to make custom icons? (y/n)")
		message, allowCustIcons = gameinput.ReadResponse()

		if !isValid(message) {
			continue
		}

		break
	}

	if allowCustIcons {
		gameIcons = getCustomIcons()
	}

	gameData := data.NewData(row, column, gameIcons)
	fmt.Println("gameData: ", gameData)
}

func getCustomIcons() string {

	return ""
}

// func runGame() {

// }

// func endGame() {

// }
