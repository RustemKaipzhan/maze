package gameflow

import (
	"fmt"
	"maze/gameinput"
	"os"
	"os/exec"
	"runtime"
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
		fmt.Print("Enter grid size (row column e.g., 4 4): ")
		*message, *row, *column = gameinput.ReadSize()
		fmt.Println(*message)
		if isValid(*message) {
			break
		}
	}
}

func promptResponse(message *string, row, column *int, allowCustomIcons *bool) {
	for {
		fmt.Printf("Grid size: %dx%d\n", *row, *column)
		fmt.Println("Would you like to make custom icons (y/n) ?")
		*message, *allowCustomIcons = gameinput.ReadResponse()

		fmt.Println(*message)

		if isValid(*message) {
			break
		}
	}
}

func promtIcon(message, icon *string, icons *map[int]string, instruction string) {
	for {
		fmt.Println(instruction)
		*message, *icon = gameinput.ReadIcon(icons)
		fmt.Println(*message)

		if isValid(*message) {
			break
		}
	}
}
