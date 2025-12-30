// Package ui provides the main UI Bubble Tea model
package ui

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ealexandrohin/tuickly/cmds"
	"github.com/ealexandrohin/tuickly/ctx"
	"github.com/ealexandrohin/tuickly/msgs"
	"github.com/ealexandrohin/tuickly/ui/footer"
	"github.com/ealexandrohin/tuickly/ui/header"
	"github.com/ealexandrohin/tuickly/ui/keymap"
	"github.com/ealexandrohin/tuickly/ux"
	"github.com/ealexandrohin/tuickly/ux/sidelist"
	"github.com/ealexandrohin/tuickly/ux/streamlist"
)

type Model struct {
	Ctx       *ctx.Ctx
	ClockTick time.Time
	Footer    footer.Model
	Header    header.Model
	KeyMap    keymap.KeyMap
	UX        ux.UX
}

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
		KeyMap:    keymap.New(),
		UX:        ux.New(ctx),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(cmds.ClockTick(), m.UX.Pacman.Tick, cmds.Live(m.Ctx))
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		// absolute garbage just so it works
		case key.Matches(msg, m.KeyMap.Up):
			m.Ctx.States.Tabs.Focused = true
			m.Ctx.States.StreamList.Focused = false
			m.Ctx.States.SideList.Focused = false
		case key.Matches(msg, m.KeyMap.Down):
			m.Ctx.States.Tabs.Focused = false
			m.Ctx.States.StreamList.Focused = true
			m.Ctx.States.SideList.Focused = false
		case key.Matches(msg, m.KeyMap.Left):
			m.Ctx.States.StreamList.Focused = false
			m.Ctx.States.SideList.Focused = true
		case key.Matches(msg, m.KeyMap.Right):
			m.Ctx.States.StreamList.Focused = true
			m.Ctx.States.SideList.Focused = false
		}
	case msgs.LiveMsg:
		m.UX.StreamList = list.New(
			msg,
			streamlist.New(m.Ctx),
			m.Ctx.Window.Width-m.Ctx.Styles.Sizes.SideList.Width,
			m.Ctx.Window.Height-m.Header.Height-m.Footer.Height,
		)
		m.UX.StreamList.SetShowTitle(false)
		m.UX.StreamList.SetShowFilter(false)
		m.UX.StreamList.SetShowStatusBar(false)
		m.UX.StreamList.SetHorizontalEnabled(true)
		m.UX.StreamList.SetStatusBarItemName("stream", "streams")

		m.Ctx.States.StreamList.Ready = true
		m.Ctx.States.StreamList.Focused = true

		m.UX.SideList = list.New(
			msg,
			sidelist.New(m.Ctx),
			m.Ctx.Styles.Sizes.SideList.Width,
			m.Ctx.Window.Height-m.Header.Height-m.Footer.Height,
		)
		m.UX.SideList.SetShowTitle(false)
		m.UX.SideList.SetShowFilter(false)
		m.UX.SideList.SetShowStatusBar(false)
		m.UX.SideList.SetShowHelp(false)
		m.UX.SideList.SetShowPagination(false)
		m.UX.SideList.SetStatusBarItemName("stream", "streams")

		m.Ctx.States.SideList.Ready = true

		return m, nil
	case spinner.TickMsg:
		var cmd tea.Cmd

		m.UX.Pacman, cmd = m.UX.Pacman.Update(msg)

		return m, cmd
	default:
		log.Printf("UI MSG:\n%+v", msg)
	}

	if m.Ctx.States.StreamList.Ready && m.Ctx.States.StreamList.Focused {
		var cmd tea.Cmd

		m.UX.StreamList, cmd = m.UX.StreamList.Update(msg)

		return m, cmd
	}

	if m.Ctx.States.SideList.Ready && m.Ctx.States.SideList.Focused {
		var cmd tea.Cmd

		m.UX.SideList, cmd = m.UX.SideList.Update(msg)

		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	var s strings.Builder
	var screen []string
	var content []string

	screen = append(screen, m.Header.View())

	// split is for future
	if m.Ctx.States.SideList.Ready {
		content = append(content, m.UX.SideList.View())
	}

	if m.Ctx.States.StreamList.Ready {
		content = append(content, m.UX.StreamList.View())
	}

	if !m.Ctx.States.StreamList.Ready && !m.Ctx.States.SideList.Ready {
		content = append(content,
			lipgloss.NewStyle().
				Height(
					m.Ctx.Window.Height-
						m.Header.Height-
						m.Footer.Height,
				).
				Width(m.Ctx.Window.Width).
				AlignHorizontal(lipgloss.Center).
				AlignVertical(lipgloss.Center).
				Render(m.UX.Pacman.View()))
	}

	screen = append(screen, lipgloss.JoinHorizontal(lipgloss.Left, content...))
	screen = append(screen, m.Footer.View())

	fmt.Fprint(&s, lipgloss.JoinVertical(lipgloss.Left, screen...))

	return s.String()
}
