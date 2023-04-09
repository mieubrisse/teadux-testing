package view

type Component[T any] interface {
	View(state T) string

	// This has to be a pointer-receiver method because I haven't figured out how to do by-value passing of
	// genericized interfaces
	Resize(width int, height int)
}
