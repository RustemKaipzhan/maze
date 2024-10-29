package gameinput

import "fmt"


func contains(icon rune) bool{
	for _, char := range "X@v" {
		 if icon == char {
			return true
		 }
	}

	return false
}
func ReadWallIcon() (string, rune){
	wall := ""

	_, err := fmt.Scanf("%s", &wall)
	
	if err != nil {
		return "Invalid input", '0'
	}

	if len(wall) < 1 || len(wall) > 1 {
		return "Invalid input", '0'
	}

	return validateIcon(rune(wall[0]))
}

func validateIcon(icon rune) (string, rune){
	if icon
}
