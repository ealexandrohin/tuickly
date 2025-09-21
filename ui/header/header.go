package header

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/eAlexandrohin/tuickly/ctx"
)

type Model struct {
	Ctx    *ctx.Ctx
	Height int
}

func New(ctx *ctx.Ctx) Model {
	return Model{
		Ctx:    ctx,
		Height: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	row := lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.Ctx.Styles.Tab.Active.Style.Render("Live"),
		m.Ctx.Styles.Tab.Style.Render("Follows"),
		m.Ctx.Styles.Tab.Style.Render("ealexandrohin"),
		m.Ctx.Styles.Tab.Style.Render("Settings"),
		m.Ctx.Styles.Tab.Style.Render("About"),
	)

	gap := m.Ctx.Styles.Tab.Gap.Style.Render(strings.Repeat(" ", max(0, m.Ctx.Window.Width-lipgloss.Width(row))))
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap, "\n")

	return row
}
