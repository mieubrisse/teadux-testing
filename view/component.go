package view

type Component[T any] interface {
	View(state T) string
	Resize(width int, height int) Component[T]
}
