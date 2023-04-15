package controller

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teadux-testing/helpers"
	"github.com/mieubrisse/teadux-testing/model"
)

func Reduce(state model.AppState, msg tea.KeyMsg) model.AppState {
	switch msg.String() {
	case "j":
		newHighlightedIdx := helpers.GetMinInt(
			state.ListState.HighlightedItemIdx+1,
			len(state.ListState.Elems)-1,
		)
		state.ListState.HighlightedItemIdx = newHighlightedIdx
	case "k":
		newHighlightedIdx := helpers.GetMaxInt(
			state.ListState.HighlightedItemIdx-1,
			0,
		)
		state.ListState.HighlightedItemIdx = newHighlightedIdx
	}
	return state
}
