package streamlist

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/ealexandrohin/tuickly/ctx"
	"github.com/ealexandrohin/tuickly/ux/items/stream"
)

// type Item struct {
// 	UserID       string
// 	UserLogin    string
// 	UserName     string
// 	GameName     string
// 	Title        string
// 	ViewerCount  int
// 	StartedAt    time.Time
// 	ThumbnailURL string
// 	Preview      string
// }
//
// func (s Item) FilterValue() string {
// 	return fmt.Sprintf("%s %s %s", s.UserLogin, s.GameName, s.Title)
// }

// func (s Item) Get() Item {
// 	return s
// }

type Delegate struct {
	Ctx *ctx.Ctx
	// UpdateFunc      func(tea.Msg, *Model) tea.Cmd
	// ShortHelpFunc   func() []key.Binding
	// FullHelpFunc    func() [][]key.Binding
}

func New(ctx *ctx.Ctx) Delegate {
	return Delegate{
		Ctx: ctx,
	}
}

func (d Delegate) Height() int {
	return d.Ctx.Styles.Sizes.StreamList.Height
}

func (d Delegate) Width() int {
	return d.Ctx.Styles.Sizes.StreamList.Width
}

func (d Delegate) Spacing() int {
	return d.Ctx.Styles.Sizes.StreamList.Spacing
}

func (d Delegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d Delegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	if m.Width() <= 0 {
		return
	}

	i, ok := item.(stream.Item)
	if !ok {
		return
	}

	var (
		title       string
		gameName    string
		viewerCount string
		s           = &d.Ctx.Styles
	)

	viewerCount = strconv.Itoa(i.ViewerCount)

	title = ansi.Truncate(
		i.Title,
		d.Ctx.Styles.Sizes.StreamList.Preview.Width,
		"…",
	)

	gameName = ansi.Truncate(
		i.GameName,
		d.Ctx.Styles.Sizes.StreamList.Preview.Width-
			len(i.UserName)-
			len(viewerCount)-
			2, // for spacing in between
		"…",
	)

	isSelected := index == m.Index()

	var (
		topRow      string
		bottomLeft  string
		bottomRight string
		bottomRow   string
	)

	if isSelected {
		topRow = s.StreamList.Selected.Top.Style.Render(title)
		bottomLeft = s.StreamList.Selected.Bottom.Left.Render(i.UserName)
		bottomRight = s.StreamList.Selected.Bottom.Right.Render(viewerCount)
		bottomRow = lipgloss.JoinHorizontal(lipgloss.Left,
			bottomLeft,
			s.StreamList.Selected.Bottom.Middle.
				Width(d.Ctx.Styles.Sizes.StreamList.Preview.Width-
					lipgloss.Width(bottomLeft)-
					lipgloss.Width(bottomRight)).
				Render(gameName),
			bottomRight,
		)
	} else {
		topRow = s.StreamList.Normal.Top.Style.Render(title)
		bottomLeft = s.StreamList.Normal.Bottom.Left.Render(i.UserName)
		bottomRight = s.StreamList.Normal.Bottom.Right.Render(viewerCount)
		bottomRow = lipgloss.JoinHorizontal(lipgloss.Left,
			bottomLeft,
			s.StreamList.Normal.Bottom.Middle.
				Width(d.Ctx.Styles.Sizes.StreamList.Preview.Width-
					lipgloss.Width(bottomLeft)-
					lipgloss.Width(bottomRight)).
				Render(gameName),
			bottomRight,
		)
	}

	// Remove last unneeded \n
	preview := i.Preview[:len(i.Preview)-1]

	stream := lipgloss.JoinVertical(lipgloss.Left, preview, topRow, bottomRow)

	if isSelected {
		stream = lipgloss.NewStyle().Border(s.StreamList.Selected.Border).BorderForeground(s.Colors.Primary).Render(stream)
	} else {
		stream = lipgloss.NewStyle().Border(s.StreamList.Normal.Border).Render(stream)
	}

	// stream = s.StreamList.Selected.Style.Border(s.StreamList.Selected.Border).Render(stream)

	fmt.Fprint(w, stream)
}
