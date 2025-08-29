package auth

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	helix "github.com/nicklaw5/helix"
)

type Model struct {
	AuthTick time.Time
	URIMsg   *URIMsg
	Width    int
	Height   int
}

type Auth struct {
	Is   bool
	User *helix.User
	Opts *helix.Options
}

func New() *Model {
	return &Model{
		AuthTick: time.Now(),
	}
}

func (m Model) Init() tea.Cmd {
	return checkAuth()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case AuthExistsMsg:
		if !msg {
			return m, newAuth()
		}

		return m, loadAuth()
	case URIMsg:
		m.URIMsg = &msg

		return m, AuthTick()
	case AuthTickMsg:
		return m, newToken(m.URIMsg)
	case ContinueTickMsg:
		return m, AuthTick()
	case TokenMsg:
		return m, checkToken(&msg)
	case TokenUserMsg:
		return m, saveAuth(&msg)
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}

	return m, nil
}

func (m Model) View() string {
	if m.URIMsg != nil {
		return m.AuthDialog()
	}

	return ""
}
