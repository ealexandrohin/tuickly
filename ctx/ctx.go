package ctx

import (
	"github.com/ealexandrohin/tuickly/ui/styles"
	helix "github.com/nicklaw5/helix"
)

type Ctx struct {
	Auth   Auth
	Window Window
	States States
	Styles styles.Styles
}

type Auth struct {
	Is   bool
	User helix.User
	Opts helix.Options
}

type Window struct {
	Width  int
	Height int
}

type States struct {
	Tabs       State
	SideList   State
	StreamList State
}

type State struct {
	Ready   bool
	Focused bool
}
