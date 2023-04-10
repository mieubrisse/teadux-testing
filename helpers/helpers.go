package helpers

const (
	padChar = ' '
)

func GetMaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetMinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// Convert the given block into a new block of the given height and width, padding
// with space
func Blockify(block [][]rune, width int, height int) [][]rune {
	result := make([][]rune, height)

	for y := 0; y < height; y++ {
		line := make([]rune, 0)
		if y < len(block) {
			line = block[y]
		}

		padLength := GetMaxInt(
			0,
			width-len(line),
		)
		pad := make([]rune, padLength)
		for x := 0; x < padLength; x++ {
			pad[x] = padChar
		}
		linePlusPad := append(line, pad...)

		result[y] = linePlusPad[:width]
	}

	return result
}

// TODO handle uneven blocks
func JoinBlocksVertically(blocks ...[][]rune) [][]rune {
	totalHeight := 0
	for _, block := range blocks {
		totalHeight += len(block)
	}

	result := make([][]rune, totalHeight)
	writeIdx := 0
	for _, block := range blocks {
		for _, line := range block {
			result[writeIdx] = line
			writeIdx++
		}
	}
	return result
}
