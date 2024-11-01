package gameinput

import (
	"fmt"
	"strconv"
)

func ReadSize() (string, int, int) {
	rowS, columnS := "", "" // row string, column string

	_, err := fmt.Scanf("%s %s", &rowS, &columnS)

	rowIsDigit, row := atoi(rowS)
	colIsDigit, column := atoi(columnS)

	// If there error or extra characters or not number
	if err != nil || !clearBuffer() || !rowIsDigit || !colIsDigit {
		return errorMessage, -1, -1
	}

	return validateSize(row, column), row, column
}

func validateSize(row, column int) string {
	warnings := []string{}
	if row <= 0 {
		warnings = append(warnings, messageSize1)
	}

	if column <= 0 {
		warnings = append(warnings, messageSize2)
	}

	if row > 99 {
		warnings = append(warnings, messageSize3)
	}

	if column > 26 {
		warnings = append(warnings, messageSize4)
	}

	if len(warnings) > 0 {
		result := ""
		for idx, warning := range warnings {
			result += "\n" + "  " + strconv.Itoa(idx+1) + ". " + warning
		}

		return warningMessage + result
	}

	return successMessage
}
