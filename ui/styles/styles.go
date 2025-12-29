package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/ealexandrohin/tuickly/ui/colors"
	"github.com/ealexandrohin/tuickly/ui/sizes"
)

type Styles struct {
	Sizes  sizes.Sizes
	Colors colors.Colors

	Err struct {
		Style lipgloss.Style
	}

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
		Style lipgloss.Style

		Normal    lipgloss.Style
		Active    lipgloss.Style
		Separator lipgloss.Style

		Left  lipgloss.Style
		Right lipgloss.Style
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

	SideList struct {
		Style  lipgloss.Style
		Border lipgloss.Border

		Normal struct {
			Style lipgloss.Style

			Top struct {
				Style lipgloss.Style

				Left  lipgloss.Style
				Right lipgloss.Style
			}

			Bottom lipgloss.Style
		}

		Selected struct {
			Style lipgloss.Style

			Top struct {
				Style lipgloss.Style

				Left  lipgloss.Style
				Right lipgloss.Style
			}

			Bottom lipgloss.Style
		}
	}
}

func New() Styles {
	s := Styles{
		Sizes:  sizes.New(),
		Colors: colors.New(),
	}

	// errors

	s.Err.Style = lipgloss.NewStyle()

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

	// tabs

	s.Tab.Style = lipgloss.NewStyle().
		Foreground(s.Colors.Text)

	s.Tab.Separator = s.Tab.Style

	// tab normal

	s.Tab.Normal = s.Tab.Style.
		Padding(0, s.Sizes.Padding)

	// tab active

	s.Tab.Active = s.Tab.Normal.
		Background(s.Colors.Primary).
		Bold(true)

	// tab left

	s.Tab.Left = s.Tab.Style.
		Background(s.Colors.Twitch).
		Padding(0, 3).
		Bold(true)

	// tab right

	s.Tab.Right = s.Tab.Left.
		Padding(0, s.Sizes.Padding)

	// statusbar

	s.Status.Style = lipgloss.NewStyle().
		Foreground(s.Colors.Text).
		Background(s.Colors.Twitch).
		Padding(0, 2)

	s.Status.Left = s.Status.Style.
		Background(s.Colors.Primary).
		Bold(true)

	s.Status.Right = s.Status.Left.Align(lipgloss.Right)

	// streamlist normal

	s.StreamList.Normal.Style = lipgloss.NewStyle().
		Foreground(s.Colors.Text)

	s.StreamList.Normal.Border = lipgloss.HiddenBorder()

	s.StreamList.Normal.Top.Style = s.StreamList.Normal.Style.
		Width(s.Sizes.StreamList.Inner.Width)

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
		Background(s.Colors.Primary).
		Bold(true)

	s.StreamList.Selected.Top.Style = s.StreamList.Selected.Style.
		Width(s.Sizes.StreamList.Inner.Width)

	s.StreamList.Selected.Bottom.Left = s.StreamList.Selected.Style.
		Align(lipgloss.Left).
		Foreground(s.Colors.Twitch)

	s.StreamList.Selected.Bottom.Middle = s.StreamList.Selected.Style.
		Align(lipgloss.Center)

	s.StreamList.Selected.Bottom.Right = s.StreamList.Selected.Style.
		Align(lipgloss.Right).
		Foreground(s.Colors.Important)

	s.StreamList.Selected.Style = s.StreamList.Selected.Style.
		Border(s.StreamList.Selected.Border)

	// sidelist

	// s.SideList.Border = lipgloss.NormalBorder()
	s.SideList.Border = lipgloss.Border{
		Right: "│",
	}

	s.SideList.Style = lipgloss.NewStyle().
		PaddingRight(1).
		Border(s.SideList.Border, false, true, false, false).
		BorderForeground(s.Colors.Background)

	// sidelist normal

	s.SideList.Normal.Style = lipgloss.NewStyle().
		Foreground(s.Colors.Text)

	s.SideList.Normal.Top.Style = s.SideList.Normal.Style

	s.SideList.Normal.Top.Left = s.SideList.Normal.Top.Style.
		Align(lipgloss.Left).
		Foreground(s.Colors.Twitch)

	s.SideList.Normal.Top.Right = s.SideList.Normal.Top.Style.
		Align(lipgloss.Right).
		Foreground(s.Colors.Important)

	s.SideList.Normal.Bottom = s.SideList.Normal.Style

	// sidelist selected

	s.SideList.Selected.Style = lipgloss.NewStyle().
		Foreground(s.Colors.Text).
		Background(s.Colors.Primary)

	s.SideList.Selected.Top.Style = s.SideList.Selected.Style.
		Bold(true)

	s.SideList.Selected.Top.Left = s.SideList.Selected.Top.Style.
		Align(lipgloss.Left).
		Foreground(s.Colors.Twitch)

	s.SideList.Selected.Top.Right = s.SideList.Selected.Top.Style.
		Align(lipgloss.Right).
		Foreground(s.Colors.Important)

	s.SideList.Selected.Bottom = s.SideList.Selected.Style

	return s
}
