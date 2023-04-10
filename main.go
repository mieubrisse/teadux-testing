package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teadux/controller"
	"github.com/mieubrisse/teadux/model"
	"github.com/mieubrisse/teadux/teact_tedux"
	"github.com/mieubrisse/teadux/view/app"
	"github.com/mieubrisse/teadux/view/displayable"
	"github.com/mieubrisse/teadux/view/displayable_array"
	"github.com/mieubrisse/teadux/view/text"
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

	/*
		var parentComponent view.Component[model.AppState] = &app.AppComponent{
			List: &scrollable_list.ScrollableList{},
		}

	*/

	var topLevelComponent displayable.Displayable[model.AppState] = app.AppComponent{
		TextArray: displayable_array.DisplayableArray[model.AppState]{
			LayoutDirection: displayable_array.Vertical,
			Displayables: []displayable.Displayable[model.AppState]{
				text.TextComponent[model.AppState]{Text: []rune("Hello world")},
				text.TextComponent[model.AppState]{Text: []rune("I'm going")},
				text.TextComponent[model.AppState]{Text: []rune("to dinner soon")},
			},
		},
	}

	topLevelModel := teact_tedux.New(state, controller.Reduce, topLevelComponent)

	p := tea.NewProgram(topLevelModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
