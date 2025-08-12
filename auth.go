package main

import (
	"encoding/gob"
	"os"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	helix "github.com/nicklaw5/helix"
)

var (
	client   *helix.Client
	clientID = "cqyppegp5st5bk2tg1nglqfd5krd4l"
	scopes   = []string{"user:read:follows", "user:read:subscriptions", "channel:read:subscriptions"}

	homePath   string
	configPath string
	authPath   string
)

type Auth struct {
	Is   bool
	User *helix.User
	Opts *helix.Options
}

type (
	AuthExistsMsg bool
	AuthTickMsg   time.Time
)

type URIMsg struct {
	URI        string
	DeviceCode string
	Response   *helix.DeviceVerificationURIResponse
}

type TokenMsg struct {
	Token    string
	Refresh  string
	Response *helix.DeviceAccessTokenResponse
}

type TokenUserMsg struct {
	Token    string
	Refresh  string
	User     helix.User
	Response *helix.UsersResponse
}

type AuthMsg struct {
	Auth Auth
}

func init() {
	var err error

	homePath, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configPath = filepath.Join(homePath, ".config", "tuickly")
	authPath = filepath.Join(configPath, "auth.gob")

	client, err = helix.NewClient(&helix.Options{
		ClientID: clientID,
	})
	if err != nil {
		panic(err)
	}
}

func checkAuth() tea.Cmd {
	return func() tea.Msg {
		if _, err := os.Stat(authPath); os.IsNotExist(err) {
			return AuthExistsMsg(false)
		}

		return AuthExistsMsg(true)
	}
}

func newAuth() tea.Cmd {
	return func() tea.Msg {
		r, err := client.RequestDeviceVerificationURI(scopes)
		if err != nil {
			return ErrorMsg{Err: err}
		}

		if r.ErrorMessage != "" {
			return ErrorMsg{Msg: r.ErrorMessage}
		}

		return URIMsg{
			URI:        r.Data.VerificationURI,
			DeviceCode: r.Data.DeviceCode,
			Response:   r,
		}
	}
}

func newToken(m *URIMsg) tea.Cmd {
	return func() tea.Msg {
		r, err := client.RequestDeviceAccessToken(m.DeviceCode, scopes)
		if err != nil {
			return ErrorMsg{Err: err}
		}

		if r.ErrorMessage == "authorization_pending" {
			return AuthTick()
		}

		if r.ErrorMessage != "" {
			return ErrorMsg{Msg: r.ErrorMessage}
		}

		return TokenMsg{
			Token:    r.Data.AccessToken,
			Refresh:  r.Data.RefreshToken,
			Response: r,
		}
	}
}

func checkToken(m *TokenMsg) tea.Cmd {
	return func() tea.Msg {
		client.SetDeviceAccessToken(m.Token)
		client.SetRefreshToken(m.Refresh)

		r, err := client.GetUsers(&helix.UsersParams{})
		if err != nil {
			return ErrorMsg{Err: err}
		}

		return TokenUserMsg{
			Token:    m.Token,
			Refresh:  m.Refresh,
			User:     r.Data.Users[0],
			Response: r,
		}
	}
}

func saveAuth(m *TokenUserMsg) tea.Cmd {
	return func() tea.Msg {
		auth := Auth{
			Is:   true,
			User: &m.User,
			Opts: &helix.Options{
				ClientID:          clientID,
				DeviceAccessToken: m.Token,
				RefreshToken:      m.Refresh,
			},
		}

		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			os.MkdirAll(filepath.Join(configPath), os.ModePerm)
		}

		file, err := os.Create(authPath)
		if err != nil {
			return ErrorMsg{Err: err}
		}
		defer file.Close()

		e := gob.NewEncoder(file)
		if err := e.Encode(auth); err != nil {
			return ErrorMsg{Err: err}
		}

		return AuthMsg{auth}
	}
}

func loadAuth() tea.Cmd {
	return func() tea.Msg {
		file, err := os.Open(authPath)
		if err != nil {
			return AuthExistsMsg(false)
		}
		defer file.Close()

		var auth Auth

		d := gob.NewDecoder(file)
		if err := d.Decode(&auth); err != nil {
			return ErrorMsg{Err: err}
		}

		client, err = helix.NewClient(auth.Opts)
		if err != nil {
			return AuthExistsMsg(false)
		}

		return AuthMsg{auth}
	}
}
