package input

type InputComponent[T any] struct {
	stateDeriver func(T) string
}

func (i InputComponent[T]) GetRunes(width int, height int, state T) [][]rune {
	text := i.stateDeriver(state)

	return [][]rune{
		// TODO nice thing of truncating the string and replacing with "..."
		[]rune(text),
	}
}
