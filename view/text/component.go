package text

import "github.com/mieubrisse/teadux/helpers"

type TextComponent[T any] struct {
	Text []rune

	// TODO wrap function???
}

func (t TextComponent[T]) GetRunes(width int, height int, _ T) [][]rune {
	runeBlock := [][]rune{
		t.Text,
	}
	return helpers.Blockify(runeBlock, width, height)
}
