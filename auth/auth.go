package auth

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/eAlexandrohin/tuickly/ctx"
)

type Model struct {
	AuthTick time.Time
	Ctx      *ctx.Ctx
	URIMsg   URIMsg
}

func New(ctx *ctx.Ctx) Model {
	return Model{
		AuthTick: time.Now(),
		Ctx:      ctx,
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
		m.URIMsg = msg

		return m, AuthTick()
	case AuthTickMsg:
		return m, newToken(m.URIMsg)
	case ContinueTickMsg:
		return m, AuthTick()
	case TokenMsg:
		return m, checkToken(&msg)
	case TokenUserMsg:
		return m, saveAuth(&msg)
		// case AuthMsg:
		// log.Println("before", m.Ctx.Auth)

		// m.Ctx.Auth = msg.Auth

		// log.Println("after", m.Ctx.Auth)
	}

	return m, nil
}

func (m Model) View() string {
	if m.URIMsg != (URIMsg{}) {
		return m.AuthDialog()
	}

	return ""
}
