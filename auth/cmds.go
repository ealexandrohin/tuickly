package auth

import (
	"encoding/gob"
	"log"
	"os"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/eAlexandrohin/tuickly/ctx"
	"github.com/eAlexandrohin/tuickly/errs"
	"github.com/eAlexandrohin/tuickly/vars"
	helix "github.com/nicklaw5/helix"
)

type (
	AuthExistsMsg   bool
	AuthTickMsg     time.Time
	ContinueTickMsg bool
)

type URIMsg struct {
	URI        string
	DeviceCode string
	// Response   *helix.DeviceVerificationURIResponse
}

type TokenMsg struct {
	Token   string
	Refresh string
	// Response *helix.DeviceAccessTokenResponse
}

type TokenUserMsg struct {
	Token   string
	Refresh string
	User    helix.User
	// Response *helix.UsersResponse
}

type AuthMsg struct {
	Auth ctx.Auth
}

func checkAuth() tea.Cmd {
	return func() tea.Msg {
		if _, err := os.Stat(vars.AuthPath); os.IsNotExist(err) {
			log.Println("auth doesnt exist")
			return AuthExistsMsg(false)
		}

		log.Println("auth exists")
		return AuthExistsMsg(true)
	}
}

func newAuth() tea.Cmd {
	return func() tea.Msg {
		r, err := vars.Client.RequestDeviceVerificationURI(vars.Scopes)
		if err != nil {
			return errs.ErrorMsg{Err: err}
		}

		if r.ErrorMessage != "" {
			return errs.ErrorMsg{Msg: r.ErrorMessage}
		}

		return URIMsg{
			URI:        r.Data.VerificationURI,
			DeviceCode: r.Data.DeviceCode,
			// Response:   r,
		}
	}
}

func AuthTick() tea.Cmd {
	return tea.Tick(5*time.Second, func(t time.Time) tea.Msg {
		log.Println("AuthTick")
		return AuthTickMsg(t)
	})
}

func newToken(m URIMsg) tea.Cmd {
	return func() tea.Msg {
		r, err := vars.Client.RequestDeviceAccessToken(m.DeviceCode, vars.Scopes)
		if err != nil {
			return errs.ErrorMsg{Err: err}
		}

		if r.ErrorMessage == "authorization_pending" {
			return ContinueTickMsg(true)
			// return AuthTick()
		}

		if r.ErrorMessage != "" {
			return errs.ErrorMsg{Msg: r.ErrorMessage}
		}

		log.Printf("auth %+v %+v", r.Data.AccessToken, r.Data.RefreshToken)

		return TokenMsg{
			Token:   r.Data.AccessToken,
			Refresh: r.Data.RefreshToken,
			// Response: r,
		}
	}
}

func checkToken(m TokenMsg) tea.Cmd {
	return func() tea.Msg {
		vars.Client.SetDeviceAccessToken(m.Token)
		vars.Client.SetRefreshToken(m.Refresh)

		r, err := vars.Client.GetUsers(&helix.UsersParams{})
		if err != nil {
			return errs.ErrorMsg{Err: err}
		}

		return TokenUserMsg{
			Token:   m.Token,
			Refresh: m.Refresh,
			User:    r.Data.Users[0],
			// Response: r,
		}
	}
}

func saveAuth(m TokenUserMsg) tea.Cmd {
	return func() tea.Msg {
		auth := ctx.Auth{
			Is:   true,
			User: m.User,
			Opts: helix.Options{
				ClientID:          vars.ClientID,
				DeviceAccessToken: m.Token,
				RefreshToken:      m.Refresh,
			},
		}

		if _, err := os.Stat(vars.ConfigPath); os.IsNotExist(err) {
			os.MkdirAll(filepath.Join(vars.ConfigPath), os.ModePerm)
		}

		file, err := os.Create(vars.AuthPath)
		if err != nil {
			return errs.ErrorMsg{Err: err}
		}
		defer file.Close()

		e := gob.NewEncoder(file)
		if err := e.Encode(auth); err != nil {
			return errs.ErrorMsg{Err: err}
		}

		return AuthMsg{auth}
	}
}

func loadAuth() tea.Cmd {
	return func() tea.Msg {
		file, err := os.Open(vars.AuthPath)
		if err != nil {
			return AuthExistsMsg(false)
		}
		defer file.Close()

		var auth ctx.Auth

		d := gob.NewDecoder(file)
		if err := d.Decode(&auth); err != nil {
			return errs.ErrorMsg{Err: err}
		}

		vars.Client, err = helix.NewClient(&auth.Opts)
		if err != nil {
			return AuthExistsMsg(false)
		}

		return AuthMsg{auth}
	}
}
