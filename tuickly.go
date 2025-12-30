// Package tuickly is Twitch TUI
package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ealexandrohin/tuickly/auth"
	"github.com/ealexandrohin/tuickly/consts"
	"github.com/ealexandrohin/tuickly/ctx"
	"github.com/ealexandrohin/tuickly/errs"
	"github.com/ealexandrohin/tuickly/msgs"
	"github.com/ealexandrohin/tuickly/ui"
	"github.com/ealexandrohin/tuickly/ui/styles"
)

type Model struct {
	Auth     auth.Model
	Ctx      *ctx.Ctx
	ErrorMsg errs.ErrorMsg
	UI       ui.Model
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
	case msgs.AuthMsg:
		m.Ctx.Auth = msg.Auth

		return m, m.UI.Init()
	case msgs.RefreshTokenMsg:
		mdl, cmd := m.Auth.Update(msg)
		m.Auth = mdl.(auth.Model)

		return m, cmd
	case msgs.RefreshAuthMsg:
		log.Println("TUICKLY REFRESHAUTHMSG")
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
		return m.Ctx.Styles.Err.Style.Width(m.Ctx.Window.Width).Render(m.ErrorMsg.Error())
	}

	if !m.Ctx.Auth.Is {
		return m.Auth.View()
	}

	return m.UI.View()
}

func main() {
	// weird hack cuz teas logging doesnt support overwriting
	f, _ := os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	f.Close()

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	log.SetOutput(f)

	consts.Program = tea.NewProgram(initialModel(),
		tea.WithAltScreen(),
		// tea.WithMouseCellMotion(),
	)

	if _, err := consts.Program.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
