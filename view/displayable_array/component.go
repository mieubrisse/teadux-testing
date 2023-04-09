package displayable_array

import (
	"github.com/mieubrisse/teadux/helpers"
	"github.com/mieubrisse/teadux/view/displayable"
	"math"
)

type LayoutDirection int

const (
	// Elements will be one after the other vertically
	Vertical LayoutDirection = iota
)

// Represents an array of Displayables, with options for dividing the space between the items
// Each item can maximally compress to one cell (though the item will disappear if there aren't enough terminal
//
//	display cells available)
type DisplayableArray[T any] struct {
	layoutDirection LayoutDirection

	displayables []displayable.Displayable[T]

	// TODO margin, padding
	// TODO different types of layouts that aren't just "even"
}

func (d DisplayableArray[T]) GetRunes(width int, height int, state T) [][]rune {
	numDisplayables := len(d.displayables)

	// TODO smarter algorithm
	heightPerElem := int(math.Round(float64(height) / float64(numDisplayables)))
	if heightPerElem == 0 {
		heightPerElem = 1
	}

	displayableBlocks := make([][][]rune, numDisplayables)
	for idx, toDisplay := range d.displayables {
		blockRunes := toDisplay.GetRunes(width, heightPerElem, state)
		// TODO optimization to skip getting runes for blocks that aren't displayed
		displayableBlocks[idx] = helpers.TruncateBlock(blockRunes, width, heightPerElem)
	}

	untruncatedResultBlock := helpers.JoinBlocksVertically(displayableBlocks...)
	result := helpers.TruncateBlock(untruncatedResultBlock, width, height)
	return result
}
