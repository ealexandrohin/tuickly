package auth

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ealexandrohin/tuickly/ctx"
	"github.com/ealexandrohin/tuickly/msgs"
)

type Model struct {
	AuthTick time.Time
	Ctx      *ctx.Ctx
	URIMsg   msgs.URIMsg
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
	case msgs.AuthExistsMsg:
		if !msg {
			return m, newAuth()
		}

		return m, loadAuth()
	case msgs.URIMsg:
		m.URIMsg = msg

		return m, authTick()
	case msgs.AuthTickMsg:
		return m, tea.Batch(newToken(m.URIMsg), authTick())
	// case msgs.ContinueTickMsg:
	// 	return m, authTick()
	case msgs.TokenMsg:
		return m, checkToken(msg)
	case msgs.TokenUserMsg:
		return m, saveAuth(msg)
	case msgs.RefreshTokenMsg:
		return m, refreshToken(msg)
	case msgs.RefreshAuthMsg:
		log.Println("AUTH REFRESHAUTHMSG")
		// case AuthMsg:
		// log.Println("before", m.Ctx.Auth)

		// m.Ctx.Auth = msg.Auth

		// log.Println("after", m.Ctx.Auth)
	}

	return m, nil
}

func (m Model) View() string {
	if m.URIMsg != (msgs.URIMsg{}) {
		return m.AuthDialog()
	}

	return ""
}
