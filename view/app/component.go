package app

import (
	"github.com/mieubrisse/teadux/model"
	"github.com/mieubrisse/teadux/view/flexbox"
)

type AppComponent struct {
	// List *scrollable_list.ScrollableList

	// width  int
	// height int

	Flexbox flexbox.Flexbox[model.AppState]
}

func (a AppComponent) View(width int, height int, state model.AppState) string {
	return a.Flexbox.View(width, height, state)
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
