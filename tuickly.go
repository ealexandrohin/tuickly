// Package tuickly is Twitch TUI.
package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/eAlexandrohin/tuickly/auth"
	"github.com/eAlexandrohin/tuickly/ctx"
	"github.com/eAlexandrohin/tuickly/errs"
	"github.com/eAlexandrohin/tuickly/ui"
	"github.com/eAlexandrohin/tuickly/ui/styles"
	"github.com/eAlexandrohin/tuickly/vars"
)

type Model struct {
	Auth     auth.Model
	Ctx      *ctx.Ctx
	ErrorMsg errs.ErrorMsg
	UI       ui.Model
	// UI       ui.Model
	// UX       ux.UX
}

func initialModel() Model {
	ctx := &ctx.Ctx{
		Styles: styles.New(),
	}

	return Model{
		Ctx:  ctx,
		Auth: auth.New(ctx),
		UI:   ui.New(ctx),
	}
}

func (m Model) Init() tea.Cmd {
	return m.Auth.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" {
			log.Println("quit")
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.Ctx.Window.Width = msg.Width
		m.Ctx.Window.Height = msg.Height
	case errs.ErrorMsg:
		m.ErrorMsg = msg
	case auth.AuthMsg:
		m.Ctx.Auth = msg.Auth

		return m, m.UI.Init()
	}

	if !m.Ctx.Auth.Is {
		mdl, cmd := m.Auth.Update(msg)
		m.Auth = mdl.(auth.Model)

		return m, cmd
	}

	mdl, cmd := m.UI.Update(msg)
	m.UI = mdl.(ui.Model)

	return m, cmd
}

func (m Model) View() string {
	if m.ErrorMsg != (errs.ErrorMsg{}) {
		return m.ErrorMsg.Error()
	}

	if !m.Ctx.Auth.Is {
		return m.Auth.View()
	}

	return m.UI.View()
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	log.SetOutput(f)

	vars.Program = tea.NewProgram(initialModel(),
		tea.WithAltScreen(),
		// tea.WithMouseCellMotion(),
	)

	if _, err := vars.Program.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
