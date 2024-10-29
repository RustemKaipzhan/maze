package gameinput

func atoi(str string) int {
	if str == "" {
		return -1
	}

	number := 0
	for _, char := range str {
		if char < '0' || char > '9' {
			return -1
		}

		digit := int(char - '0')
		number = 10*number + digit
	}

	return number
}
