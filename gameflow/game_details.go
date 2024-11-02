package gameflow

import "strconv"

var (
	gameDetailsTop = "\033[1;34mGame Details 📋\033[0m"
	gridSizeDetail = "\n\033[1;37m– Grid Size:\033[0m "
	iconsDetail    = "\n\033[1;37m– Game Icons:\033[0m"
	gameDetails    = ""
)

func addGridSizeDetail() {
	row, column := gameData.GetSize()
	gridSizeDetail += "\033[1;32m" + strconv.Itoa(row) + "x" + strconv.Itoa(column) + "\033[0m"
	gameDetails = gameDetailsTop + gridSizeDetail
}
func updateIconDetail() {
	icons := gameData.GetIcons()
	gameDetails = gameDetailsTop + gridSizeDetail + iconsDetail + "\033[1;33m" +
		"\n  • wall   " + icons[0] +
		"\n  • player " + icons[2] +
		"\n  • award  " + icons[3] + "\033[0m"
}
