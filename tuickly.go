// Package tuickly is Twitch TUI.
package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/eAlexandrohin/tuickly/auth"
	"github.com/eAlexandrohin/tuickly/errs"
	"github.com/eAlexandrohin/tuickly/ui"
	"github.com/eAlexandrohin/tuickly/ux"
)

type Model struct {
	Auth     *auth.Auth
	AuthMdl  *auth.Model
	ErrorMsg errs.ErrorMsg
	UI       *ui.Model
	UX       *ux.UX
}

func initialModel() *Model {
	a := &auth.Auth{}

	return &Model{
		AuthMdl: auth.New(),
		Auth:    a,
		UI:      ui.New(a),
		// UX:      &ux.UX{},
	}

	// m := &Model{}
	//
	// m.AuthMdl = auth.New()
	// m.Auth = &auth.Auth{}
	// m.UI = ui.New()
	// m.UX = &ux.UX{}
	//
	// return m
}

func (m Model) Init() tea.Cmd {
	return m.AuthMdl.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			log.Println("quit")
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.UI.Window.Width = msg.Width
		m.UI.Window.Height = msg.Height
	case auth.AuthMsg:
		log.Println("before", m.Auth)

		m.Auth = msg.Auth

		log.Println("after", m.Auth)
	case errs.ErrorMsg:
		m.ErrorMsg = msg
	}

	if !m.Auth.Is {
		mdl, cmd := m.AuthMdl.Update(msg)
		m.AuthMdl = &mdl

		return m, cmd
	}

	mdl, cmd := m.UI.Update(msg)
	m.UI = &mdl

	return m, cmd

	// if m.Content.Loaded {
	// 	var cmd tea.Cmd
	// 	m.UX.List, cmd = m.UX.List.Update(msg)
	//
	// 	return m, cmd
	// }

	// return m, nil
}

func (m Model) View() string {
	if m.ErrorMsg != (errs.ErrorMsg{}) {
		return m.ErrorMsg.Error()
	}

	if !m.Auth.Is {
		return m.AuthMdl.View()
	}

	return m.UI.View()
}

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()

		log.SetOutput(f)
	}

	p := tea.NewProgram(initialModel(),
		tea.WithAltScreen(),
		// tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
