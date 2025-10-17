package footer

import (
	"time"

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
	left := m.Ctx.Styles.Status.Left.Render("tuickly")
	// right := m.Ctx.Styles.Status.Right.Render(time.Now().Format(time.TimeOnly))
	right := m.Ctx.Styles.Status.Right.Render(time.Now().Format(time.TimeOnly))
	bar := m.Ctx.Styles.Status.Style.
		Width(m.Ctx.Window.Width - lipgloss.Width(left) - lipgloss.Width(right)).
		Render("@ealexandrohin")

	status := lipgloss.JoinHorizontal(lipgloss.Top, left, bar, right)

	return status
}
