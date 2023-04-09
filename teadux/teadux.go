package teadux

import tea "github.com/charmbracelet/bubbletea"

type Reducer func(msg tea.Msg)

type Teadux[T any] struct {
	state   T
	reducer Reducer
}

func New[T any](
	initialState T,
	reducer Reducer,
) Teadux[T] {
	return Teadux[T]{
		state:   initialState,
		reducer: reducer,
	}
}

func (td Teadux[T]) GetState() T {
	return td.state
}

func Reduce() {

}
