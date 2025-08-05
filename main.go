package main

import (
	// "fmt"
	// helix "github.com/nicklaw5/helix"
	"encoding/json"
	browser "github.com/pkg/browser"
	"net/http"
	"net/url"
)

const clientID = "cqyppegp5st5bk2tg1nglqfd5krd4l"

func auth() {
	response, err := http.PostForm("https://id.twitch.tv/oauth2/device", url.Values{
		"client_id": {clientID},
		"scopes":    {"user:read:follows user:read:subscriptions channel:read:subscriptions"},
	})

	if err != nil {
		panic(err)
	}

	var body map[string]interface{}

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		panic(err)
	}

	browser.OpenURL(body["verification_uri"].(string))
}

func main() {
	auth()
}
