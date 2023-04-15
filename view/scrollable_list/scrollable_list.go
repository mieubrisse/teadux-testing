package scrollable_list

import (
	"github.com/mieubrisse/teadux-testing/helpers"
	"github.com/mieubrisse/teadux-testing/model"
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
		lastRuneIdxExclusive := helpers.GetMinInt(list.width, len(lineRunes))
		lineStr := string(lineRunes[:lastRuneIdxExclusive])
		allLineStrs = append(allLineStrs, lineStr)

		if idx >= list.height {
			break
		}
	}
	return strings.Join(allLineStrs, "\n")
}

func (list *ScrollableList) Resize(width int, height int) {
	list.width = width
	list.height = height
}
