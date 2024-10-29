package gameinput

import "fmt"

// Read player y/n
func ReadResponse() (string, bool) {
	response := ""

	_, err := fmt.Scanf("%s", &response)

	if err != nil {
		return "Invalid input", false
	}

	if len(response) < 1 || len(response) > 1 || !clearBuffer() {
		return "Invalid input", false
	}

	return validateResponse(rune(response[0]))

}

func validateResponse(response rune) (string, bool) {
	if response != 'n' && response != 'y' {
		return "Invalid input", false
	}

	return "success", true
}
