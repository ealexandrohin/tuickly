package vars

import (
	"os"
	"path/filepath"

	helix "github.com/nicklaw5/helix"
)

const ClientID = "cqyppegp5st5bk2tg1nglqfd5krd4l"

var (
	Client *helix.Client
	Scopes = []string{"user:read:follows", "user:read:subscriptions", "channel:read:subscriptions"}

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
}
