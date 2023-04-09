package displayable

type Displayable[T any] interface {
	// Block of printable runes for displaying this component
	// The return value is an array of lines
	// Note that the returned value does NOT need to conform to the
	//   width & height; it is provided only so the displayable can
	//   nicely format its own text
	// It is expected that the parent will truncate to the desired length
	GetRunes(width int, height int, state T) [][]rune
}
