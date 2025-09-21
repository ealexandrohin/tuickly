package auth

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) AuthDialog() string {
	URI := m.Ctx.Styles.Auth.URI.Style.
		Width(len(m.URIMsg.URI) + 4).
		Render(m.URIMsg.URI)

	return lipgloss.Place(
		m.Ctx.Window.Width,
		m.Ctx.Window.Height,
		lipgloss.Center,
		lipgloss.Center,
		m.Ctx.Styles.Auth.Dialog.Style.Render(URI),
	)
}
