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

func atoi(str string) (bool, int) {
	// If str is empty
	if str == "" {
		return false, -1
	}

	sign := 1
	startIdx := 0
	if len(str) > 1 && str[0] == '-' {
		sign = -1
		startIdx = 1
	}

	number := 0

	for i := startIdx; i < len(str); i++ {
		char := rune(str[i])

		// If is not digit
		if char < '0' || char > '9' {
			return false, -1
		}

		digit := int(char - '0')
		number = 10*number + digit
	}

	return true, number * sign
}
