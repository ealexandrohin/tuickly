package auth

import (
	"encoding/gob"
	"log"
	"os"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ealexandrohin/tuickly/consts"
	"github.com/ealexandrohin/tuickly/ctx"
	"github.com/ealexandrohin/tuickly/errs"
	"github.com/ealexandrohin/tuickly/msgs"
	helix "github.com/nicklaw5/helix"
)

func checkAuth() tea.Cmd {
	return func() tea.Msg {
		if _, err := os.Stat(consts.AuthPath); os.IsNotExist(err) {
			log.Println("auth doesnt exist")
			return msgs.AuthExistsMsg(false)
		}

		log.Println("auth exists")
		return msgs.AuthExistsMsg(true)
	}
}

func newAuth() tea.Cmd {
	return func() tea.Msg {
		r, err := consts.Client.RequestDeviceVerificationURI(consts.Scopes)
		if err != nil {
			return errs.ErrorMsg{Err: err}
		}

		if r.ErrorMessage != "" {
			return errs.ErrorMsg{Msg: r.ErrorMessage}
		}

		return msgs.URIMsg{
			URI:        r.Data.VerificationURI,
			DeviceCode: r.Data.DeviceCode,
			// Response:   r,
		}
	}
}

func authTick() tea.Cmd {
	return tea.Tick(5*time.Second, func(t time.Time) tea.Msg {
		log.Println("AuthTick")
		return msgs.AuthTickMsg(t)
	})
}

func newToken(m msgs.URIMsg) tea.Cmd {
	return func() tea.Msg {
		r, err := consts.Client.RequestDeviceAccessToken(m.DeviceCode, consts.Scopes)
		if err != nil {
			return errs.ErrorMsg{Err: err}
		}

		if r.ErrorMessage == "authorization_pending" {
			// return msgs.AuthTickMsg(true)
			return authTick()
		}

		if r.ErrorMessage != "" {
			return errs.ErrorMsg{Msg: r.ErrorMessage}
		}

		log.Printf("auth %+v %+v", r.Data.AccessToken, r.Data.RefreshToken)

		return msgs.TokenMsg{
			Token:   r.Data.AccessToken,
			Refresh: r.Data.RefreshToken,
			// Response: r,
		}
	}
}

func checkToken(m msgs.TokenMsg) tea.Cmd {
	return func() tea.Msg {
		consts.Client.SetDeviceAccessToken(m.Token)
		consts.Client.SetRefreshToken(m.Refresh)

		r, err := consts.Client.GetUsers(&helix.UsersParams{})
		if err != nil {
			return errs.ErrorMsg{Err: err}
		}

		return msgs.TokenUserMsg{
			Token:   m.Token,
			Refresh: m.Refresh,
			User:    r.Data.Users[0],
		}
	}
}

func saveAuth(m msgs.TokenUserMsg) tea.Cmd {
	return func() tea.Msg {
		auth := ctx.Auth{
			Is:   true,
			User: m.User,
			Opts: helix.Options{
				ClientID:          consts.ClientID,
				DeviceAccessToken: m.Token,
				RefreshToken:      m.Refresh,
			},
		}

		if _, err := os.Stat(consts.ConfigPath); os.IsNotExist(err) {
			os.MkdirAll(filepath.Join(consts.ConfigPath), os.ModePerm)
		}

		file, err := os.Create(consts.AuthPath)
		if err != nil {
			return errs.ErrorMsg{Err: err}
		}
		defer file.Close()

		e := gob.NewEncoder(file)
		if err := e.Encode(auth); err != nil {
			return errs.ErrorMsg{Err: err}
		}

		return msgs.AuthMsg{auth}
	}
}

func loadAuth() tea.Cmd {
	return func() tea.Msg {
		file, err := os.Open(consts.AuthPath)
		if err != nil {
			return msgs.AuthExistsMsg(false)
		}
		defer file.Close()

		var auth ctx.Auth

		d := gob.NewDecoder(file)
		if err := d.Decode(&auth); err != nil {
			return errs.ErrorMsg{Err: err}
		}

		consts.Client, err = helix.NewClient(&auth.Opts)
		if err != nil {
			return msgs.AuthExistsMsg(false)
		}

		return msgs.AuthMsg{Auth: auth}
	}
}

func refreshToken(m msgs.RefreshTokenMsg) tea.Cmd {
	return func() tea.Msg {
		consts.Client.SetDeviceAccessToken(m.Token)
		consts.Client.SetRefreshToken(m.Refresh)

		r, err := consts.Client.GetUsers(&helix.UsersParams{})
		if err != nil {
			return errs.ErrorMsg{Err: err}
		}

		auth := ctx.Auth{
			Is:   true,
			User: r.Data.Users[0],
			Opts: helix.Options{
				ClientID:          consts.ClientID,
				DeviceAccessToken: m.Token,
				RefreshToken:      m.Refresh,
			},
		}

		if _, err := os.Stat(consts.ConfigPath); os.IsNotExist(err) {
			os.MkdirAll(filepath.Join(consts.ConfigPath), os.ModePerm)
		}

		file, err := os.Create(consts.AuthPath)
		if err != nil {
			return errs.ErrorMsg{Err: err}
		}
		defer file.Close()

		e := gob.NewEncoder(file)
		if err := e.Encode(auth); err != nil {
			return errs.ErrorMsg{Err: err}
		}

		return msgs.RefreshAuthMsg{Auth: auth}
	}
}
