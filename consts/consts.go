// Package consts consits of global constants
package consts

import (
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ealexandrohin/tuickly/msgs"
	helix "github.com/nicklaw5/helix/v2"
)

const ClientID = "cqyppegp5st5bk2tg1nglqfd5krd4l"

var (
	Program *tea.Program
	Client  *helix.Client
	Scopes  = []string{"user:read:follows", "user:read:subscriptions", "channel:read:subscriptions"}

	HomePath   string
	ConfigPath string
	AuthPath   string
)

func init() {
	var err error

	HomePath, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	ConfigPath = filepath.Join(HomePath, ".config", "tuickly")
	AuthPath = filepath.Join(ConfigPath, "auth.gob")

	Client, err = helix.NewClient(&helix.Options{
		ClientID: ClientID,
	})
	if err != nil {
		panic(err)
	}

	Client.OnUserAccessTokenRefreshed(func(newAccessToken string, newRefreshToken string) {
		// apparently this doesnt work
		// TODO: fix refreshing tokens
		Program.Send(msgs.RefreshTokenMsg{
			Token:   newAccessToken,
			Refresh: newRefreshToken,
		})
	})
}
