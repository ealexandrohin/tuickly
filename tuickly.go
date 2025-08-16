// Package tuickly is Twitch TUI.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	lst "github.com/eAlexandrohin/tuickly/list"
	helix "github.com/nicklaw5/helix"
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
	List      lst.Model
	// View      string
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

type StreamItem struct {
	Stream helix.Stream
}

// func (s StreamItem) Title() string {
// 	return s.Stream.UserName
// }

func (s StreamItem) FilterValue() string {
	return s.Stream.UserName
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

		return m, Live(&m)
	case LiveMsg:
		m.List = lst.Model{
			List: list.New(lst.Items, list.NewDefaultDelegate(), 0, 0),
			H:    m.Window.Height,
			W:    m.Window.Width,
		}
		// items := make([]list.Item, len(msg))
		//
		// for i, s := range msg {
		// 	items[i] = StreamItem{s}
		// }
		//
		// m.List = lst.Model{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}

		// m.List = lst.Model{List: list.New(msg, list.NewDefaultDelegate(), 0, 0)}

		// return m, nil
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

	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.List.View(), m.footerView())
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
