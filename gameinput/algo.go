package gameinput

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

func contains(icon rune) bool {
	for _, char := range "X@v" {
		if icon == char {
			return true
		}
	}

	return false
}
