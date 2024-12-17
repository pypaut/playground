package d04

func CountXmas(lines []string) (nbXmas int) {
	for lineNb, line := range lines {
		for charNb, char := range line {
			if char == 'X' {
				if CheckTopTop(lines, lineNb, charNb) {
					nbXmas++
				}

				if CheckTopRight(lines, lineNb, charNb) {
					nbXmas++
				}

				if CheckTopLeft(lines, lineNb, charNb) {
					nbXmas++
				}

				if CheckBottomBottom(lines, lineNb, charNb) {
					nbXmas++
				}

				if CheckBottomRight(lines, lineNb, charNb) {
					nbXmas++
				}

				if CheckBottomLeft(lines, lineNb, charNb) {
					nbXmas++
				}

				if CheckRightRight(lines, lineNb, charNb) {
					nbXmas++
				}

				if CheckLeftLeft(lines, lineNb, charNb) {
					nbXmas++
				}
			}
		}
	}

	return
}

func CanFitLeft(charNb int) bool {
	return charNb-3 >= 0
}

func CanFitRight(charNb, lineLen int) bool {
	return charNb+3 < lineLen
}

func CanFitTop(lineNb int) bool {
	return lineNb-3 >= 0
}

func CanFitBottom(lineNb, nbLines int) bool {
	return lineNb+3 < nbLines
}

func CheckTopTop(lines []string, lineNb, charNb int) bool {
	return CanFitTop(lineNb) &&
		lines[lineNb-1][charNb] == 'M' &&
		lines[lineNb-2][charNb] == 'A' &&
		lines[lineNb-3][charNb] == 'S'
}

func CheckTopRight(lines []string, lineNb, charNb int) bool {
	return CanFitTop(lineNb) &&
		CanFitRight(charNb, len(lines[lineNb])) &&
		lines[lineNb-1][charNb+1] == 'M' &&
		lines[lineNb-2][charNb+2] == 'A' &&
		lines[lineNb-3][charNb+3] == 'S'
}

func CheckTopLeft(lines []string, lineNb, charNb int) bool {
	return CanFitTop(lineNb) &&
		CanFitLeft(charNb) &&
		lines[lineNb-1][charNb-1] == 'M' &&
		lines[lineNb-2][charNb-2] == 'A' &&
		lines[lineNb-3][charNb-3] == 'S'
}

func CheckBottomBottom(lines []string, lineNb, charNb int) bool {
	return CanFitBottom(lineNb, len(lines)) &&
		lines[lineNb+1][charNb] == 'M' &&
		lines[lineNb+2][charNb] == 'A' &&
		lines[lineNb+3][charNb] == 'S'
}

func CheckBottomRight(lines []string, lineNb, charNb int) bool {
	return CanFitBottom(lineNb, len(lines)) &&
		CanFitRight(charNb, len(lines[lineNb])) &&
		lines[lineNb+1][charNb+1] == 'M' &&
		lines[lineNb+2][charNb+2] == 'A' &&
		lines[lineNb+3][charNb+3] == 'S'
}

func CheckBottomLeft(lines []string, lineNb, charNb int) bool {
	return CanFitBottom(lineNb, len(lines)) &&
		CanFitLeft(charNb) &&
		lines[lineNb+1][charNb-1] == 'M' &&
		lines[lineNb+2][charNb-2] == 'A' &&
		lines[lineNb+3][charNb-3] == 'S'
}

func CheckRightRight(lines []string, lineNb, charNb int) bool {
	return CanFitRight(charNb, len(lines[lineNb])) &&
		lines[lineNb][charNb+1] == 'M' &&
		lines[lineNb][charNb+2] == 'A' &&
		lines[lineNb][charNb+3] == 'S'
}

func CheckLeftLeft(lines []string, lineNb, charNb int) bool {
	return CanFitLeft(charNb) &&
		lines[lineNb][charNb-1] == 'M' &&
		lines[lineNb][charNb-2] == 'A' &&
		lines[lineNb][charNb-3] == 'S'
}
