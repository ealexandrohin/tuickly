// Package tuickly is Twitch TUI.
package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ActiveTab int
	Auth      Auth
	AuthTick  time.Time
	Content   string
	ErrorMsg  ErrorMsg
	Msg       tea.Msg
	Tabs      []Tab
	UI        UI
	Window    Window
	// View      string
}

func initialModel() Model {
	return Model{
		AuthTick: time.Now(),
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
	}
}

type ErrorMsg struct {
	Msg string
	Err error
}

func (e ErrorMsg) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Msg
}

func AuthTick() tea.Cmd {
	return tea.Tick(5*time.Second, func(t time.Time) tea.Msg {
		return AuthTickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return checkAuth()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case AuthExistsMsg:
		if !msg {
			return m, newAuth()
		}

		return m, loadAuth()
	case URIMsg:
		m.Msg = msg

		return m, AuthTick()
	case AuthTickMsg:
		var cmds []tea.Cmd

		if uriMsg, ok := m.Msg.(URIMsg); ok {
			cmds = append(cmds, newToken(&uriMsg))
		}

		cmds = append(cmds, AuthTick())

		return m, tea.Batch(cmds...)
	case TokenMsg:
		m.Msg = nil

		return m, checkToken(&msg)
	case TokenUserMsg:
		return m, saveAuth(&msg)
	case AuthMsg:
		m.Auth = msg.Auth

		return m, start(&m)
	case ContentMsg:
		m.Content += string(msg) + "\n"

		return m, nil
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.Window.Width = msg.Width
		m.Window.Height = msg.Height
	case ErrorMsg:
		m.ErrorMsg = msg
	}

	return m, nil
}

func (m Model) View() string {
	if m.ErrorMsg != (ErrorMsg{}) {
		return m.ErrorMsg.Error()
	}

	if !m.Auth.Is {
		if _, ok := m.Msg.(URIMsg); ok {
			return URIDialog(m)
		}
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.Content, m.footerView())
	// return m.Content
}

func main() {
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}
