package ui

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/eAlexandrohin/tuickly/auth"
	"github.com/eAlexandrohin/tuickly/cmds"
	"github.com/eAlexandrohin/tuickly/ui/footer"
	"github.com/eAlexandrohin/tuickly/ui/header"
	"github.com/eAlexandrohin/tuickly/ui/styles"
	"github.com/eAlexandrohin/tuickly/ui/window"
	"github.com/eAlexandrohin/tuickly/ux"
)

type Model struct {
	Auth   *auth.Auth
	Footer *footer.Model
	Header *header.Model
	Styles *styles.Styles
	UX     *ux.UX
	Window *window.Window
}

// type Tab struct {
// 	Title   string
// 	Content string
// }

func New(a *auth.Auth) *Model {
	s := styles.New()
	w := &window.Window{}

	return &Model{
		Auth:   a,
		Footer: footer.New(w, s),
		Header: header.New(w, s),
		Styles: s,
		UX:     &ux.UX{},
		Window: w,
	}
}

func (m Model) Init() tea.Cmd {
	return cmds.Live(m.Auth)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cmds.LiveMsg:
		m.UX.List.Mdl = list.New(msg, ux.StreamDelegate{}, m.Window.Width, m.Window.Height-(m.Header.Height+m.Footer.Height))
		m.UX.List.Mdl.Title = "Live"
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
	log.Println(m.Auth)
	if m.UX.List.Ready {
		s.WriteString(m.UX.List.Mdl.View() + "\n")
	}
	s.WriteString(m.Footer.View())

	return s.String()
}
