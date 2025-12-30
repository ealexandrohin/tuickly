// Package msgs defines Bubble Tea messages
package msgs

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/ealexandrohin/tuickly/ctx"
	helix "github.com/nicklaw5/helix/v2"
)

// auth msgs

type (
	AuthExistsMsg bool
	AuthTickMsg   time.Time
	// ContinueTickMsg bool
)

type URIMsg struct {
	URI        string
	DeviceCode string
}

type TokenMsg struct {
	Token   string
	Refresh string
}

type TokenUserMsg struct {
	Token   string
	Refresh string
	User    helix.User
}

type AuthMsg struct {
	Auth ctx.Auth
}

type RefreshTokenMsg struct {
	Token   string
	Refresh string
}

type RefreshAuthMsg struct {
	Auth ctx.Auth
}

// ui msgs

type ClockTick time.Time

// cmds msgs

type LiveMsg []list.Item
