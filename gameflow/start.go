package gameflow

// "fmt"
import (
	"fmt"
	"maze/data"
	"maze/gameinput"
)

func StartGame() {

	message, row, column := gameinput.ReadSize()
	if message != "success" {
		fmt.Println(message)
		return
	}
	gameData := data.NewData(row, column)
	fmt.Println("gameData: ", gameData)
}

// func runGame() {

// }

// func endGame() {

// }
