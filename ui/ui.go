package ui

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/eAlexandrohin/tuickly/cmds"
	"github.com/eAlexandrohin/tuickly/ctx"
	"github.com/eAlexandrohin/tuickly/ui/footer"
	"github.com/eAlexandrohin/tuickly/ui/header"
	"github.com/eAlexandrohin/tuickly/ux"
	"github.com/eAlexandrohin/tuickly/ux/streamlist"
)

type Model struct {
	Ctx    *ctx.Ctx
	Footer footer.Model
	Header header.Model
	UX     ux.UX
}

// type UI struct {
// 	Ctx    *ctx.Ctx
// 	Footer footer.Model
// 	Header header.Model
// 	// Styles styles.Styles
// 	UX ux.UX
// }

// type Tab struct {
// 	Title   string
// 	Content string
// }

func New(ctx *ctx.Ctx) Model {
	return Model{
		Ctx:    ctx,
		Footer: footer.New(ctx),
		Header: header.New(ctx),
		UX:     ux.New(ctx),
	}
}

func (m Model) Init() tea.Cmd {
	return cmds.Live(m.Ctx)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cmds.LiveMsg:
		// m.UX.List.Mdl = list.New(msg, streamlist.New(m.Ctx), m.Ctx.Window.Width, m.Ctx.Window.Height-(m.Header.Height+m.Footer.Height))
		log.Println(m.Ctx.Window)
		m.UX.List.Mdl = list.New(msg, streamlist.New(m.Ctx), m.Ctx.Window.Width, 33)
		m.UX.List.Mdl.SetShowTitle(false)
		m.UX.List.Mdl.SetShowFilter(false)
		m.UX.List.Mdl.SetShowStatusBar(false)
		m.UX.List.Ready = true

		return m, nil
	}

	if m.UX.List.Ready {
		var cmd tea.Cmd

		m.UX.List.Mdl, cmd = m.UX.List.Mdl.Update(msg)

		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	// return m.Main.UX.List.View()

	var s strings.Builder

	s.WriteString(m.Header.View())

	// log.Println(m.Ctx.Auth)

	s.WriteString("\n")

	if m.UX.List.Ready {
		s.WriteString(m.UX.List.Mdl.View() + "\n")
	}

	s.WriteString(m.Footer.View())

	return s.String()
}
