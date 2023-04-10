package text

type TextComponent[T any] struct {
	Text string

	// TODO wrap function???
}

func (t TextComponent[T]) View(width int, height int, _ T) string {
	return t.Text
}
