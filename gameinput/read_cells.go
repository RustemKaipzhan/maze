package gameinput

import "fmt"

func ReadCellsLine() (string, string) {
	line := ""

	_, err := fmt.Scanf("%s", &line)

	if err != nil || clearBuffer() {
		return errorMessage, ""
	}

	return validateCellsLine(line)
}
func validateCellsLine(line string) (string, string) {
	for _, symbol := range line {
		if symbol < '0' || symbol > '3' {
			return warningMessage + "only characters 0, 1, 2, 4 are allowed!", ""
		}
	}

	return successMessage, line
}
