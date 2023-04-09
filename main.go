package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teadux/controller"
	"github.com/mieubrisse/teadux/model"
	"github.com/mieubrisse/teadux/teact_tedux"
	"github.com/mieubrisse/teadux/view"
	"github.com/mieubrisse/teadux/view/app"
	"github.com/mieubrisse/teadux/view/scrollable_list"
	"os"
)

func main() {
	elems := []string{}
	for i := 0; i < 60; i++ {
		elems = append(elems, fmt.Sprintf("foo bar bang blork %d", i))
	}

	state := model.AppState{
		ListState: model.ListState{
			Elems:              elems,
			HighlightedItemIdx: 0,
		},
	}

	var parentComponent view.Component[model.AppState] = &app.AppComponent{
		List: &scrollable_list.ScrollableList{},
	}

	topLevelModel := teact_tedux.New(state, controller.Reduce, parentComponent)

	p := tea.NewProgram(topLevelModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
