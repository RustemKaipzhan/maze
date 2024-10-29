package gameinput

import "fmt"

func ReadSize() (string, int, int) {
	rowS, columnS := "", "" // row string, column string

	_, err := fmt.Scanf("%s %s", &rowS, &columnS)
	if err != nil || !clearBuffer() {
		return "Invalid input", -1, -1
	}

	row, column := atoi(rowS), atoi(columnS)

	return validateRowColumn(row, column), row, column
}

func validateRowColumn(row, column int) string {
	result := ""
	if row <= 0 {
		result += "\nRow mustn't be negative or 0"
	}

	if column <= 0 {
		result += "\nColumn mustn't be negative or 0"
	}

	if row > 99 {
		result += "\nRow mustn't be greater than 99"
	}
	if column > 26 {
		result += "\nColumn mustn't be greater than 26"
	}

	if len(result) > 0 {
		return "Invalid input:" + result
	}

	return "success"
}