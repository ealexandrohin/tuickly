package main

import tea "github.com/charmbracelet/bubbletea"

type ContentMsg bool

func start(m Model) tea.Cmd {
	return func() tea.Msg {
		return ContentMsg(false)
	}
}
