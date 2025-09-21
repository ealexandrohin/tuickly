package colors

import "github.com/charmbracelet/lipgloss"

type Colors struct {
	Important lipgloss.Color
	Primary   lipgloss.Color
	Text      lipgloss.Color
	Twitch    lipgloss.Color
}

var (
	// official twitch branding colors
	purple = lipgloss.Color("#9146FF") // twitch primary
	black  = lipgloss.Color("#000000") // twitch black ops
	ice    = lipgloss.Color("#F0F0FF") // twitch ice

	// twitch colors
	pacman = lipgloss.Color("#FFBF00") // twitch pacman

	// base colors
	white = lipgloss.Color("#FFFFFF")
	red   = lipgloss.Color("#FF0000")
	green = lipgloss.Color("#00FF00")
	blue  = lipgloss.Color("#0000FF")
)

func New() Colors {
	return Colors{
		Important: red,
		Primary:   pacman,
		Text:      white,
		Twitch:    purple,
	}
}
