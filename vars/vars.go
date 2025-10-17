package vars

import (
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ealexandrohin/tuickly/msgs"
	helix "github.com/nicklaw5/helix"
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

	// proxyURL, _ := url.Parse("http://t879694:raiTh+e3hoot@192.168.1.240:3128")
	// http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	Client, err = helix.NewClient(&helix.Options{
		ClientID: ClientID,
	})
	if err != nil {
		panic(err)
	}

	Client.OnUserAccessTokenRefreshed(func(newAccessToken string, newRefreshToken string) {
		Program.Send(msgs.RefreshTokenMsg{
			Token:   newAccessToken,
			Refresh: newRefreshToken,
		})
	})
}
