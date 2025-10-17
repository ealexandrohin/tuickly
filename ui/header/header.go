package header

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ealexandrohin/tuickly/ctx"
)

type Model struct {
	Ctx    *ctx.Ctx
	Height int
}

func New(ctx *ctx.Ctx) Model {
	return Model{
		Ctx:    ctx,
		Height: 1,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	var tabs []string

	tabs = append(tabs,
		m.Ctx.Styles.Tab.Normal.Render(""),
		m.Ctx.Styles.Tab.Active.Render("Following"),
		m.Ctx.Styles.Tab.Normal.Render("Browse"),
		m.Ctx.Styles.Tab.Normal.Render(""),
	)

	left := m.Ctx.Styles.Tab.Left.Render("twitch")
	right := m.Ctx.Styles.Tab.Right.Render("ealexandrohin")
	middle := m.Ctx.Styles.Tab.Style.
		Width(m.Ctx.Window.Width - lipgloss.Width(left) - lipgloss.Width(right)).
		Render(strings.Join(tabs, m.Ctx.Styles.Tab.Separator.Render("|")))

	row := lipgloss.JoinHorizontal(
		lipgloss.Left,
		left,
		middle,
		right,
	)

	return row
}
