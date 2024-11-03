package gameflow

import (
	"os"
	"os/exec"
	"runtime"
	"sort"
	"time"
)

// Return sorted indexes of icons
func getSortedIconIndexes() []int {
	icons := gameData.GetIcons()
	indexes := make([]int, 0, len(icons))
	for idx := range icons {
		// if it is free cell
		if idx == 1 {
			continue
		}

		indexes = append(indexes, idx)
	}

	sort.Ints(indexes)
	return indexes
}

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

func wait(sec float64, clearAfter bool) {
	time.Sleep(time.Duration(sec * float64(time.Second)))
	if clearAfter {
		clearTerminal()
	}
}
