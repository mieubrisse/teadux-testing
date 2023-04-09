package scrollable_list

import (
	"github.com/mieubrisse/teadux/model"
	"strings"
)

type ScrollableList struct {
	// TODO changing the selected-item char
	width  int
	height int
}

func (list ScrollableList) View(state model.ListState) string {
	allLineStrs := []string{}
	for idx, itemString := range state.Elems {
		prefix := "  "
		if idx == state.HighlightedItemIdx {
			prefix = "x "
		}
		lineRunes := []rune(prefix + itemString)
		lineStr := string(lineRunes[:list.width])
		allLineStrs = append(allLineStrs, lineStr)

		if idx >= list.height {
			break
		}
	}
	return strings.Join(allLineStrs, "\n")
}

func (list ScrollableList) Resize(width int, height int) ScrollableList {
	list.width = width
	list.height = height

	return list
}
