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
	helix "github.com/nicklaw5/helix/v2"
)

// checkAuth is a Bubble Tea command that checks for the existence of the
// auth file.
// It sends an [msgs.AuthExistsMsg] indicating whether the file was found.
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

// newAuth is a Bubble Tea command that initiates a new DCF with Twitch.
// It requests a device verification URI and code, then sends a [msgs.URIMsg]
// or an [errs.ErrorMsg] if the request fails.
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

// authTick is a Bubble Tea command that sends an [msgs.AuthTickMsg]
// every 5 seconds. This is used to periodically check for device token
// status during the auth flow.
func authTick() tea.Cmd {
	return tea.Tick(5*time.Second, func(t time.Time) tea.Msg {
		log.Println("AuthTick")
		return msgs.AuthTickMsg(t)
	})
}

// newToken is a Bubble Tea command that attempts to request a device access
// token using the provided device code.
// It handles `authorization_pending` by scheduling another [authTick].
// On success, it sends a [msgs.TokenMsg]; otherwise, an [errs.ErrorMsg].
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

// checkToken is a Bubble Tea command that validates the received access
// token by fetching user's data.
// It sets the access and refresh tokens on the global [helix.Client].
// On success, it sends a [msgs.TokenUserMsg]; otherwise, an [errs.ErrorMsg].
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

// saveAuth is a Bubble Tea command that persists the auth data
// [ctx.Auth] to a file using gob encoding.
// It creates the configuration directory if it doesn't exist.
// On success, it sends an [msgs.AuthMsg]; otherwise, an [errs.ErrorMsg].
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

// loadAuth is a Bubble Tea command that loads auth data from the saved file.
// It decodes the gob-encoded data into an [ctx.Auth] and re-initializes
// the global [helix.Client] with the saved options. On success,
// it sends an [msgs.AuthMsg]; otherwise, it sends [msgs.AuthExistsMsg(false)]
// or an [errs.ErrorMsg] if decoding or client re-initialization fails.
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

// refreshToken is a Bubble Tea command that handles an incoming
// [msgs.RefreshTokenMsg]. It updates the global [helix.Client] with
// the new access and refresh tokens, fetches the user's data,
// and then saves the updated auth context to the file.
// On success, it sends a [msgs.RefreshAuthMsg]; otherwise, an [ErrorMsg].
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
