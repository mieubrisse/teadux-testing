package flexbox

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teadux/view/displayable"
)

type LayoutDirection int
type Justification int

const (
	Horizontal LayoutDirection = iota
	Vertical

	// Orientations around the main and cross axes
	FlexStart Justification = iota
	Center
	FlexEnd
)

// Represents an array of Items, with options for dividing the space between the items
// Each item can maximally compress to one cell (though the item will disappear if there aren't enough terminal
//
//	display cells available)
type Flexbox[T any] struct {
	LayoutDirection LayoutDirection

	// TODO alignment?

	Items []displayable.Displayable[T]

	// TODO margin, padding
	// TODO different types of layouts that aren't just "even"
}

func (d Flexbox[T]) View(width int, height int, state T) string {
	numItems := len(d.Items)

	// TODO consider extrinsic height & width

	// TODO what to do with expanded items????
	itemViews := make([]string, numItems)
	for idx, item := range d.Items {
		// TODO height??
		itemViews[idx] = item.View(width, height/numItems, state)
	}

	var result string
	switch d.LayoutDirection {
	case Horizontal:
		result = lipgloss.JoinHorizontal(lipgloss.Center, itemViews...)
	case Vertical:
		result = lipgloss.JoinVertical(lipgloss.Left, itemViews...)
	default:
		panic("Unknown flexbox direction")
	}

	return result
}
