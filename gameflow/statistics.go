package gameflow

func currentIcons(icons map[int]string) string {
	return "Game icons:" +
		"\n  1. wall - " + icons[0] +
		"\n  2. player - " + icons[2] +
		"\n  3. award - " + icons[3]
}
