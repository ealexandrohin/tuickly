package footer

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/eAlexandrohin/tuickly/ui/styles"
	"github.com/eAlexandrohin/tuickly/ui/window"
)

type Model struct {
	Height int
	Styles *styles.Styles
	Window *window.Window
}

func New(w *window.Window, s *styles.Styles) *Model {
	return &Model{
		Styles: s,
		Window: w,
		Height: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return &m, nil
}

func (m *Model) View() string {
	left := m.Styles.Status.Left.Render("tuickly")
	right := m.Styles.Status.Right.Render(time.Now().Format(time.TimeOnly))
	bar := m.Styles.Status.Style.Width(m.Window.Width - lipgloss.Width(left) - lipgloss.Width(right)).Render("@ealexandrohin")

	status := lipgloss.JoinHorizontal(lipgloss.Top, left, bar, right, "\n")
	m.Height = lipgloss.Height(status)

	return status
}
