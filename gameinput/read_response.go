package gameinput

import "fmt"

// Read player y/n
func ReadResponse() (string, bool) {
	response := ""
	_, err := fmt.Scanf("%s", &response)

	if err != nil || !clearBuffer() {
		return errorMessage, false
	}

	return validateResponse(response)
}

func validateResponse(response string) (string, bool) {
	if len(response) != 1 {
		return warningMessage + messageResponse1, false
	}

	if response[0] != 'n' && response[0] != 'y' {
		return warningMessage + messageResponse2, false
	}

	return successMessage, true
}
