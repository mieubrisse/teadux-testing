package div

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teadux/view/displayable"
)

type Alignment lipgloss.Position

const (
	Left   = Alignment(lipgloss.Left)
	Center = Alignment(lipgloss.Center)
	Right  = Alignment(lipgloss.Right)
)

type Border lipgloss.Border

type Div[T any] struct {
	// TODO margin, padding

	// TODO maybe just use the lipgloss values????
	Alignment Alignment

	Items []displayable.Displayable[T]

	// Setting these will cause the div to be forced to this
	ExtrinsicWidth  int
	ExtrinsicHeight int

	// TODO extract this into some sort of generic Element class??
	Border lipgloss.Border
}

func (d Div[T]) View(width int, height int, state T) string {
	itemViews := make([]string, len(d.Items))
	for i, item := range d.Items {
		itemViews[i] = item.View(width, height, state)
	}

	verticallyJoined := lipgloss.JoinVertical(lipgloss.Position(d.Alignment), itemViews...)

	// Necessary to truncate + do the math because lipgloss applies border *before* truncating
	truncatingStyle := lipgloss.NewStyle()
	if d.ExtrinsicWidth > 0 {
		contentWidth := d.ExtrinsicWidth - d.Border.GetLeftSize() - d.Border.GetRightSize()
		truncatingStyle = truncatingStyle.MaxWidth(contentWidth)
	}
	if d.ExtrinsicHeight > 0 {
		contentHeight := d.ExtrinsicHeight - d.Border.GetTopSize() - d.Border.GetBottomSize()
		truncatingStyle = truncatingStyle.MaxHeight(contentHeight)
	}
	truncated := truncatingStyle.Render(verticallyJoined)

	borderStyle := lipgloss.NewStyle().BorderStyle(d.Border)
	return borderStyle.Render(truncated)
}
