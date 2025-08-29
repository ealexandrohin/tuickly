package auth

import (
	"github.com/charmbracelet/lipgloss"
)

var dialogStyle = lipgloss.NewStyle().
	Border(lipgloss.ThickBorder()).
	BorderForeground(lipgloss.Color("ffbf00")).
	Padding(1, 0)

func (m Model) AuthDialog() string {
	URI := lipgloss.NewStyle().
		Width(len(m.URIMsg.URI) + 4).
		Align(lipgloss.Center).
		Render(m.URIMsg.URI)

	return lipgloss.Place(
		m.Width,
		m.Height,
		lipgloss.Center,
		lipgloss.Center,
		dialogStyle.Render(URI),
	)
}
