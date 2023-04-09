package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teadux/controller"
	"github.com/mieubrisse/teadux/model"
	"github.com/mieubrisse/teadux/teact_tedux"
	"github.com/mieubrisse/teadux/view/app_component"
	"github.com/mieubrisse/teadux/view/scrollable_list"
	"os"
	"regexp"
)

var acceptableFormFieldRegex = regexp.MustCompile("^[a-zA-Z0-9.-]+$")

func main() {
	state := model.AppState{
		ListState: model.ListState{
			Elems: []string{
				"foo",
				"bar",
				"bang",
			},
			HighlightedItemIdx: 0,
		},
	}

	parentComponent := app_component.AppComponent{List: scrollable_list.ScrollableList{}}

	topLevelModel := teact_tedux.New(state, controller.Reduce, parentComponent)

	p := tea.NewProgram(topLevelModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
