package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teadux/controller"
	"github.com/mieubrisse/teadux/model"
	"github.com/mieubrisse/teadux/teact_tedux"
	"github.com/mieubrisse/teadux/view/app"
	"github.com/mieubrisse/teadux/view/displayable"
	"github.com/mieubrisse/teadux/view/div"
	"github.com/mieubrisse/teadux/view/flexbox"
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

	/*
		var topLevelComponent displayable.Displayable[model.AppState] = app.AppComponent{
			Flexbox: flexbox.Flexbox[model.AppState]{
				LayoutDirection: flexbox.Vertical,
				Displayables: []displayable.Displayable[model.AppState]{
					text.TextComponent[model.AppState]{Text: []rune("Hello world")},
					text.TextComponent[model.AppState]{Text: []rune("I'm going")},
					text.TextComponent[model.AppState]{Text: []rune("to dinner soon")},
				},
			},
		}

	*/

	/*
		innerFlexbox := flexbox.Flexbox[model.AppState]{
			LayoutDirection: flexbox.Horizontal,
			Items: []displayable.Displayable[model.AppState]{
				text.TextComponent[model.AppState]{Text: "Yes, this is cool!"},
				text.TextComponent[model.AppState]{Text: "This is also cool!"},
			},
		}
	*/

	topLevelFlexbox := flexbox.Flexbox[model.AppState]{
		LayoutDirection: flexbox.Vertical,
		Items: []displayable.Displayable[model.AppState]{
			flexbox.Flexbox[model.AppState]{
				Items: []displayable.Displayable[model.AppState]{
					div.Div[model.AppState]{
						Items: []displayable.Displayable[model.AppState]{
							text.TextComponent[model.AppState]{Text: "Hello"},
						},
						Border: lipgloss.NormalBorder(),
					},
					div.Div[model.AppState]{
						Items: []displayable.Displayable[model.AppState]{
							text.TextComponent[model.AppState]{Text: "       Cool"},
						},
						Border: lipgloss.NormalBorder(),
					},
					div.Div[model.AppState]{
						Items: []displayable.Displayable[model.AppState]{
							text.TextComponent[model.AppState]{Text: "            World"},
						},
						Border: lipgloss.NormalBorder(),
					},
				},
			},
			/*
				text.TextComponent[model.AppState]{Text: "Hello world"},
				text.TextComponent[model.AppState]{Text: "I'm going"},
				text.TextComponent[model.AppState]{Text: "to dinner soon"},
			*/
		},
	}

	var topLevelComponent displayable.Displayable[model.AppState] = app.AppComponent{
		Flexbox: topLevelFlexbox,
	}

	topLevelModel := teact_tedux.New(state, controller.Reduce, topLevelComponent)

	p := tea.NewProgram(topLevelModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
