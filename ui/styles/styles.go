package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/eAlexandrohin/tuickly/ui/colors"
)

// colors from lipgloss example
// Foreground(lipgloss.Color("#C1C6B2")).
// Background(lipgloss.Color("#353533")).

type Styles struct {
	Colors *colors.Colors

	Tab struct {
		Style  lipgloss.Style
		Border lipgloss.Border

		Active struct {
			Style  lipgloss.Style
			Border lipgloss.Border
		}

		Gap struct {
			Style  lipgloss.Style
			Border lipgloss.Border
		}
	}

	Status struct {
		Style lipgloss.Style
		Left  lipgloss.Style
		Right lipgloss.Style
	}

	List struct {
		Normal struct {
			Style       lipgloss.Style
			UserName    lipgloss.Style
			GameName    lipgloss.Style
			Uptime      lipgloss.Style
			ViewerCount lipgloss.Style
			Title       lipgloss.Style
		}

		Selected struct {
			Style       lipgloss.Style
			UserName    lipgloss.Style
			GameName    lipgloss.Style
			Uptime      lipgloss.Style
			ViewerCount lipgloss.Style
			Title       lipgloss.Style
		}

		Dimmed struct {
			Style       lipgloss.Style
			UserName    lipgloss.Style
			GameName    lipgloss.Style
			Uptime      lipgloss.Style
			ViewerCount lipgloss.Style
			Title       lipgloss.Style
		}

		FilterMatch lipgloss.Style
	}
}

func New() *Styles {
	s := Styles{}
	s.Colors = colors.New()

	// tabs borders

	s.Tab.Border = lipgloss.Border{
		Top:         "─",
		Bottom:      "━",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "┷",
		BottomRight: "┷",
	}

	s.Tab.Active.Border = lipgloss.Border{
		Top:         "━",
		Bottom:      " ",
		Left:        "┃",
		Right:       "┃",
		TopLeft:     "┏",
		TopRight:    "┓",
		BottomLeft:  "┛",
		BottomRight: "┗",
	}

	s.Tab.Gap.Border = lipgloss.Border{
		Bottom:      "━",
		BottomLeft:  "━",
		BottomRight: "━",
	}

	// tabs

	s.Tab.Style = lipgloss.NewStyle().
		Border(s.Tab.Border).
		BorderForeground(s.Colors.Primary).
		Padding(0, 1)

	s.Tab.Active.Style = s.Tab.Style.Border(s.Tab.Active.Border)

	s.Tab.Gap.Style = lipgloss.NewStyle().
		Border(s.Tab.Gap.Border).
		BorderForeground(s.Colors.Primary)

	// statusbar

	s.Status.Left = lipgloss.NewStyle().
		Foreground(s.Colors.Text).
		Background(s.Colors.Primary).
		Padding(0, 1)

	s.Status.Style = lipgloss.NewStyle().
		Foreground(s.Colors.Text).
		Background(s.Colors.Twitch).
		// Faint(true).
		Padding(0, 1)

	s.Status.Right = s.Status.Left.Align(lipgloss.Right)

	// list
	// normal

	s.List.Normal.Style = lipgloss.NewStyle().
		Foreground(s.Colors.Text)

	s.List.Normal.UserName = s.List.Normal.Style.
		Background(s.Colors.Twitch).
		Padding(0, 0, 0, 2)

	s.List.Normal.GameName = s.List.Normal.Style

	s.List.Normal.Uptime = s.List.Normal.Style

	s.List.Normal.ViewerCount = s.List.Normal.Style.
		Background(s.Colors.Important) // Classic

	s.List.Normal.Title = s.List.Normal.Style.
		Padding(0, 0, 0, 2)

	// selected

	s.List.Selected.Style = lipgloss.NewStyle().
		Foreground(s.Colors.Text).
		Background(s.Colors.Primary)

	s.List.Selected.UserName = s.List.Selected.Style.
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(s.Colors.Primary).
		Foreground(s.Colors.Twitch).
		Background(s.Colors.Primary).
		Padding(0, 0, 0, 1)

	s.List.Selected.GameName = s.List.Selected.Style

	s.List.Selected.Uptime = s.List.Selected.Style

	s.List.Selected.ViewerCount = s.List.Selected.Style.
		Foreground(s.Colors.Important)

	s.List.Selected.Title = s.List.Selected.Style.
		Padding(0, 0, 0, 1)

	// dimmed

	s.List.Dimmed.Style = lipgloss.NewStyle().
		Foreground(s.Colors.Text).
		Faint(true)

	s.List.Dimmed.UserName = s.List.Dimmed.Style.
		Background(s.Colors.Twitch).
		Padding(0, 0, 0, 2)

	s.List.Dimmed.GameName = s.List.Dimmed.Style

	s.List.Dimmed.Uptime = s.List.Dimmed.Style

	s.List.Dimmed.ViewerCount = s.List.Dimmed.Style.
		Background(s.Colors.Important)

	s.List.Dimmed.Title = s.List.Dimmed.Style.
		Padding(0, 0, 0, 2)

	s.List.FilterMatch = s.List.Selected.Style
	// s.List.FilterMatch = lipgloss.NewStyle().Underline(true)

	return &s
}
