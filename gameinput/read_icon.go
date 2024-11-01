package gameinput

import (
	"fmt"
)

func ReadIcon(icons *map[int]string) (string, string) {
	icon := ""
	_, err := fmt.Scanf("%s", &icon)

	if err != nil || !clearBuffer() {
		return errorMessage, "0"
	}

	return validateIcon(icon, icons)
}

func validateIcon(icon string, icons *map[int]string) (string, string) {
	if len(icon) > 1 {
		return warningMessage + messageIcon1, "0"
	}

	if (*icons)[0] == icon || (*icons)[2] == icon || (*icons)[3] == icon {
		return warningMessage + messageIcon2, "0"
	}

	if icon == "/" {
		return successMessage, ""
	}

	return successMessage, icon
}
