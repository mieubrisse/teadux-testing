package teact_tedux

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teadux/view"
)

type Reducer[T any] func(currentState T, msg tea.KeyMsg) T

type TeactTedux[T any] struct {
	// Could timetravel if we want
	state T

	reducer Reducer[T]

	parentComponent view.Component[T]
}

func New[T any](
	initialState T,
	reducer Reducer[T],
	parentComponent view.Component[T],
) TeactTedux[T] {
	return TeactTedux[T]{
		state:           initialState,
		reducer:         reducer,
		parentComponent: parentComponent,
	}

}

func (model TeactTedux[T]) Init() tea.Cmd {
	return nil
}

func (model TeactTedux[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return model, tea.Quit
		}
		model.state = model.reducer(model.state, msg)
	case tea.WindowSizeMsg:
		model.parentComponent.Resize(msg.Width, msg.Height)
	}

	return model, nil
}

func (model TeactTedux[T]) View() string {
	return model.parentComponent.View(model.state)
}

/*
func (model teadux) Resize(width int, height int) teadux {
	model.width = width
	model.height = height

	model.list = model.list.Resize(width, height)

	return model
}

func (model teadux) GetHeight() int {
	return model.height
}

func (model teadux) GetWidth() int {
	return model.width
}
*/
