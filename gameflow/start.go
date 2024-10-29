package gameflow

// "fmt"
import (
	"fmt"
	"maze/data"
	"maze/gameinput"
)

func StartGame() {
	message, row, column := "", 0, 0

	for {
		fmt.Print("Enter grid size (row column e.g., 4 7): ")
		message, row, column = gameinput.ReadSize()

		if message != "success" {
			fmt.Println(message)
			continue
		} else {
			break
		}
	}

	gameData := data.NewData(row, column)
	fmt.Println("gameData: ", gameData)
}

// func runGame() {

// }

// func endGame() {

// }
