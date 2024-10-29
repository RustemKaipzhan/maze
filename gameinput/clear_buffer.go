package gameinput

import "fmt"

// Read line and check if it is clean
func clearBuffer() bool {
	var ch byte
	clean := true

	// Read and discard characters until newLine is encountered
	for {
		_, err := fmt.Scanf("%c", &ch)
		if ch == '\n' {
			break
		}

		// If it is any symbol except whitespace
		if err != nil || ch != ' ' {
			clean = false
		}
	}

	return clean
}
