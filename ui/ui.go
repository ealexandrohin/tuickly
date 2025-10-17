package ui

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ealexandrohin/tuickly/cmds"
	"github.com/ealexandrohin/tuickly/ctx"
	"github.com/ealexandrohin/tuickly/msgs"
	"github.com/ealexandrohin/tuickly/ui/footer"
	"github.com/ealexandrohin/tuickly/ui/header"
	"github.com/ealexandrohin/tuickly/ux"
	"github.com/ealexandrohin/tuickly/ux/sidelist"
	"github.com/ealexandrohin/tuickly/ux/streamlist"
)

type Model struct {
	Ctx       *ctx.Ctx
	ClockTick time.Time
	Footer    footer.Model
	Header    header.Model
	UX        ux.UX
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
		Ctx:       ctx,
		ClockTick: time.Now(),
		Footer:    footer.New(ctx),
		Header:    header.New(ctx),
		UX:        ux.New(ctx),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(cmds.ClockTick(), cmds.Live(m.Ctx))
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case msgs.LiveMsg:
		m.UX.StreamList.Mdl = list.New(msg, streamlist.New(m.Ctx), m.Ctx.Window.Width-m.Ctx.Styles.Sizes.SideList.Width, m.Ctx.Window.Height-m.Header.Height-m.Footer.Height)
		m.UX.StreamList.Mdl.SetShowTitle(false)
		m.UX.StreamList.Mdl.SetShowFilter(false)
		m.UX.StreamList.Mdl.SetShowStatusBar(false)
		m.UX.StreamList.Mdl.SetHorizontalView(true)
		m.UX.StreamList.Ready = true

		m.UX.SideList.Mdl = list.New(msg, sidelist.New(m.Ctx), m.Ctx.Styles.Sizes.SideList.Width, m.Ctx.Window.Height-m.Header.Height-m.Footer.Height)
		m.UX.SideList.Mdl.SetShowTitle(false)
		m.UX.SideList.Mdl.SetShowFilter(false)
		m.UX.SideList.Mdl.SetShowStatusBar(false)
		m.UX.SideList.Mdl.SetShowHelp(false)
		m.UX.SideList.Mdl.SetShowPagination(false)
		m.UX.SideList.Ready = true

		return m, nil
	case msgs.ClockTick:
		return m, cmds.ClockTick()
	}

	// if m.UX.StreamList.Ready {
	// 	var cmd tea.Cmd
	//
	// 	m.UX.StreamList.Mdl, cmd = m.UX.StreamList.Mdl.Update(msg)
	//
	// 	return m, cmd
	// }

	if m.UX.SideList.Ready {
		var cmd tea.Cmd

		m.UX.SideList.Mdl, cmd = m.UX.SideList.Mdl.Update(msg)

		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	var s strings.Builder

	// s.WriteString(m.Header.View())
	//
	// if m.UX.SideList.Ready {
	// 	m.UX.StreamList.Mdl.StartSpinner()
	// 	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Left, m.UX.StreamList.Mdl.View(), m.UX.SideList.Mdl.View()))
	// 	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Left, m.UX.SideList.Mdl.View(), m.UX.StreamList.Mdl.View()))
	// 	s.WriteString(m.UX.SideList.Mdl.View() + "\n")
	// 	s.WriteString(m.UX.StreamList.Mdl.View() + "\n")
	// }
	//
	// s.WriteString(m.Footer.View())

	// s.WriteString(
	// 	lipgloss.JoinVertical(lipgloss.Left,
	// 		m.Header.View(),
	// 		lipgloss.JoinHorizontal(lipgloss.Left,
	// 			m.Ctx.Styles.SideList.Style.Render(m.UX.SideList.Mdl.View()),
	// 			m.UX.StreamList.Mdl.View(),
	// 		),
	// 		m.Footer.View(),
	// 	),
	// )

	s.WriteString(
		lipgloss.JoinVertical(lipgloss.Left,
			m.Header.View(),
			lipgloss.JoinHorizontal(lipgloss.Left,
				m.UX.SideList.Mdl.View(),
				m.UX.StreamList.Mdl.View(),
			),
			m.Footer.View(),
		),
	)

	return s.String()
}
