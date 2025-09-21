package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/eAlexandrohin/tuickly/ui/colors"
)

// colors from lipgloss example
// Foreground(lipgloss.Color("#C1C6B2")).
// Background(lipgloss.Color("#353533")).

type Styles struct {
	Colors colors.Colors

	Auth struct {
		Dialog struct {
			Style  lipgloss.Style
			Border lipgloss.Border
		}

		URI struct {
			Style lipgloss.Style
		}
	}

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

	StreamList struct {
		Normal struct {
			Style  lipgloss.Style
			Border lipgloss.Border

			Top struct {
				Style lipgloss.Style
			}

			Bottom struct {
				Left   lipgloss.Style
				Middle lipgloss.Style
				Right  lipgloss.Style
			}
		}

		Selected struct {
			Style  lipgloss.Style
			Border lipgloss.Border

			Top struct {
				Style lipgloss.Style
			}

			Bottom struct {
				Left   lipgloss.Style
				Middle lipgloss.Style
				Right  lipgloss.Style
			}
		}
	}
}

func New() Styles {
	s := Styles{
		Colors: colors.New(),
	}

	// auth dialog

	s.Auth.Dialog.Border = lipgloss.ThickBorder()

	s.Auth.Dialog.Style = lipgloss.NewStyle().
		Border(s.Auth.Dialog.Border).
		BorderForeground(s.Colors.Primary).
		Padding(1, 0)

	// auth uri

	s.Auth.URI.Style = lipgloss.NewStyle().
		Foreground(s.Colors.Text).
		Align(lipgloss.Center)

	// tab normal

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

	s.Tab.Style = lipgloss.NewStyle().
		Border(s.Tab.Border).
		BorderForeground(s.Colors.Primary).
		Padding(0, 1)

	// tab active

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

	s.Tab.Active.Style = s.Tab.Style.Border(s.Tab.Active.Border)

	// tab gap

	s.Tab.Gap.Border = lipgloss.Border{
		Bottom:      "━",
		BottomLeft:  "━",
		BottomRight: "━",
	}

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

	// streamlist normal

	s.StreamList.Normal.Style = lipgloss.NewStyle().
		Foreground(s.Colors.Text)

	s.StreamList.Normal.Top.Style = s.StreamList.Normal.Style.
		Width(40)

	s.StreamList.Normal.Bottom.Left = s.StreamList.Normal.Style.
		Align(lipgloss.Left).
		Foreground(s.Colors.Twitch).
		Bold(true)

	s.StreamList.Normal.Bottom.Middle = s.StreamList.Normal.Style.
		Align(lipgloss.Center)

	s.StreamList.Normal.Bottom.Right = s.StreamList.Normal.Style.
		Align(lipgloss.Right).
		Foreground(s.Colors.Important).
		Bold(true)

	// streamlist selected

	s.StreamList.Selected.Border = lipgloss.Border{
		Top:      "▄",
		Left:     "█",
		Right:    "█",
		TopLeft:  "▄",
		TopRight: "▄",
	}

	s.StreamList.Selected.Style = s.StreamList.Normal.Style.
		Background(s.Colors.Primary)

	s.StreamList.Selected.Top.Style = s.StreamList.Selected.Style.
		Width(40)

	s.StreamList.Selected.Bottom.Left = s.StreamList.Selected.Style.
		Align(lipgloss.Left).
		Foreground(s.Colors.Twitch).
		Bold(true)

	s.StreamList.Selected.Bottom.Middle = s.StreamList.Selected.Style.
		Align(lipgloss.Center)

	s.StreamList.Selected.Bottom.Right = s.StreamList.Selected.Style.
		Align(lipgloss.Right).
		Foreground(s.Colors.Important).
		Bold(true)

	// s.StreamList.Selected.Style = s.StreamList.Selected.Style.
	// 	Border(s.StreamList.Selected.Border)

	return s
}
