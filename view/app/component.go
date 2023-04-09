package app

import (
	"github.com/mieubrisse/teadux/model"
	"github.com/mieubrisse/teadux/view/scrollable_list"
)

type AppComponent struct {
	List *scrollable_list.ScrollableList

	width  int
	height int
}

func (a AppComponent) View(state model.AppState) string {
	return a.List.View(state.ListState)
}

func (a *AppComponent) Resize(width int, height int) {
	a.width = width
	a.height = height

	a.List.Resize(width, height)
}
