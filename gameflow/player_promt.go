package gameflow

import (
	"fmt"
	"maze/gameinput"
	"os"
	"os/exec"
	"runtime"
)

const (
	enter            = "\033[1;36m→ Enter: "
	guideGridSize    = enter + "grid size (e.g., 3 3):\033[0m "
	guideCustomIcons = "- Would you like to make custom icons ?\n"
	guideResponse    = "\033[1;34m→ Select an option (y/n):\033[0m "
	guideSkip        = "\033[0m (type '/' to skip): "
)

func clearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // For Windows
	} else {
		cmd = exec.Command("clear") // For Unix-based systems
	}

	cmd.Stdout = os.Stdout // Set the output to the terminal
	cmd.Run()              // Execute the command
}

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

		fmt.Println(*message)

		if isValid(*message) {
			break
		}
	}
}

func promptIcon(message, icon *string, instruction string) {
	for {
		fmt.Println(instruction)
		*message, *icon = gameinput.ReadIcon(gameData.GetIcons())
		fmt.Println(*message)

		if isValid(*message) {
			break
		}
	}
}

func promptGrid(message *string) {
	for {

		if isValid(*message) {
			break
		}
	}
}
