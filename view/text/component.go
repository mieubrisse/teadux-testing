package text

type TextComponent struct {
	text []rune

	// TODO wrap function???
}

// TODO maybe take in a string and split it based on newlines??
func New(text []rune) TextComponent {
	return TextComponent{text: text}
}

func (t TextComponent) GetRunes(width int, height int, _ interface{}) [][]rune {
	return [][]rune{
		t.text,
	}
}
