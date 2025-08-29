package header

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/eAlexandrohin/tuickly/ui/styles"
	"github.com/eAlexandrohin/tuickly/ui/window"
)

type Model struct {
	Styles *styles.Styles
	Window *window.Window
	Height int
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
	return m, nil
}

func (m Model) View() string {
	row := lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.Styles.Tab.Active.Style.Render("Live"),
		m.Styles.Tab.Style.Render("Follows"),
		m.Styles.Tab.Style.Render("ealexandrohin"),
		m.Styles.Tab.Style.Render("Settings"),
		m.Styles.Tab.Style.Render("About"),
	)

	gap := m.Styles.Tab.Gap.Style.Render(strings.Repeat(" ", max(0, m.Window.Width-lipgloss.Width(row))))
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap, "\n")

	return row
}
