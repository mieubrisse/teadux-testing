package list

import "github.com/mieubrisse/teadux-testing/view/displayable"

type List[T any] struct {
	contentDeriver func(T) []displayable.Displayable[T]
}

func (l List[T]) GetRunes(width int, height int, state T) [][]rune {
	//TODO implement me
	panic("implement me")
}
