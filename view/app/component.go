package app

import (
	"github.com/mieubrisse/teadux/model"
	"github.com/mieubrisse/teadux/view/displayable_array"
)

type AppComponent struct {
	// List *scrollable_list.ScrollableList

	// width  int
	// height int

	TextArray displayable_array.DisplayableArray[model.AppState]
}

func (a AppComponent) GetRunes(width int, height int, state model.AppState) [][]rune {
	return a.TextArray.GetRunes(width, height, state)
}

/*
func (a AppComponent) View(state model.AppState) string {
	return a.List.View(state.ListState)
}

func (a *AppComponent) Resize(width int, height int) {
	a.width = width
	a.height = height

	a.List.Resize(width, height)
}


*/
