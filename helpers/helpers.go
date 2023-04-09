package helpers

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

func TruncateBlock(block [][]rune, width int, height int) [][]rune {
	result := make([][]rune, height)
	for idx, line := range block[:height] {
		result[idx] = line[:width]
	}
	return result
}

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
