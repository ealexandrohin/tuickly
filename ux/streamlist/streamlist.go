// Package streamlist provides the item delegate for main stream list
package streamlist

import (
	"fmt"
	_ "image/jpeg"
	"io"
	"regexp"
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/ealexandrohin/tuickly/ctx"
	"github.com/ealexandrohin/tuickly/ux/items/stream"
)

type Delegate struct {
	Ctx *ctx.Ctx
	// UpdateFunc func(tea.Msg, *list.Model) tea.Cmd
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
		innerWidth  = s.Sizes.StreamList.Inner.Width
	)

	// stupid ass emojis
	regex := regexp.MustCompile(`[\s\x{FE00}-\x{FE0F}\x{1F3FB}-\x{1F3FF}]+`)

	viewerCount = strconv.Itoa(i.ViewerCount)

	title = ansi.Truncate(
		regex.ReplaceAllString(i.Title, " "),
		innerWidth,
		"…",
	)

	gameName = ansi.Truncate(
		i.GameName,
		innerWidth-
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

	if isSelected && d.Ctx.States.StreamList.Focused {
		topRow = s.StreamList.Selected.Top.Style.Render(title)
		bottomLeft = s.StreamList.Selected.Bottom.Left.Render(i.UserName)
		bottomRight = s.StreamList.Selected.Bottom.Right.Render(viewerCount)
		bottomRow = lipgloss.JoinHorizontal(lipgloss.Left,
			bottomLeft,
			s.StreamList.Selected.Bottom.Middle.
				Width(innerWidth-
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
				Width(innerWidth-
					lipgloss.Width(bottomLeft)-
					lipgloss.Width(bottomRight)).
				Render(gameName),
			bottomRight,
		)
	}

	// Remove last unneeded \n
	preview := i.Preview[:len(i.Preview)-1]

	stream := lipgloss.JoinVertical(lipgloss.Left, preview, topRow, bottomRow)

	if isSelected && d.Ctx.States.StreamList.Focused {
		stream = lipgloss.NewStyle().Border(s.StreamList.Selected.Border).BorderForeground(s.Colors.Primary).Render(stream)
	} else {
		stream = lipgloss.NewStyle().Border(s.StreamList.Normal.Border).Render(stream)
	}

	fmt.Fprint(w, stream)
}
