// Package tuickly is Twitch TUI.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Auth      Auth
	Tabs      []Tab
	ActiveTab int
	Window    Window
	Msg       tea.Msg
}

type Window struct {
	Width  int
	Height int
}

type Tab struct {
	Title   string
	Content string
}

func (m Model) Init() tea.Cmd {
	return checkAuth()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.Msg = msg

	// var (
	// 	cmd  tea.Cmd
	// 	cmds []tea.Cmd
	// )

	switch msg := msg.(type) {
	case AuthExistsMsg:
		if !msg {
			return m, newAuth()
		}

		return m, loadAuth(m)
	case URIMsg:

		// return m, newToken(msg)
		return m, nil
	case TokenMsg:
		return m, checkToken(msg)
	case TokenUserMsg:
		return m, saveAuth(msg)
	case AuthMsg:
		if !msg {
			return m, loadAuth(m)
		}

		return m, start(m)
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.Window.Width = msg.Width
		m.Window.Height = msg.Height

		// headerHeight := lipgloss.Height(m.headerView())
		// footerHeight := lipgloss.Height(m.footerView())
		// verticalMarginHeight := headerHeight + footerHeight
		//
		// if !m.Ready {
		// 	m.Viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
		// 	m.Viewport.YPosition = headerHeight
		// 	m.Viewport.SetContent(m.Content)
		// 	m.Ready = true
		// } else {
		// 	m.Viewport.Width = msg.Width
		// 	m.Viewport.Height = msg.Height - verticalMarginHeight
		// }
	}

	// // Handle keyboard and mouse events in the viewport
	// m.Viewport, cmd = m.Viewport.Update(msg)
	// cmds = append(cmds, cmd)
	//
	// return m, tea.Batch(cmds...)
	return m, nil
}

func (m Model) View() string {
	if _, ok := m.Msg.(URIMsg); ok {
		return URIDialog(m)
	}

	return ""
	// if !m.Ready {
	// 	return "\nInitializing..."
	// }
	// return m.Content
	// return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func main() {
	p := tea.NewProgram(
		Model{
			Tabs: []Tab{
				{
					Title:   "Live",
					Content: "Live",
				},
				{
					Title:   "Follows",
					Content: "Follows",
				},
				{
					Title:   "ealexandrohin",
					Content: "ealexandrohin",
				},
				{
					Title:   "About",
					Content: "About",
				},
			},
		},
		tea.WithAltScreen(),      // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseAllMotion(), // turn on mouse support so we can track the mouse wheel
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
