// Package sidelist provides the item delegate for sidebar list
package sidelist

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/ealexandrohin/tuickly/ctx"
	"github.com/ealexandrohin/tuickly/utils"
	"github.com/ealexandrohin/tuickly/ux/items/stream"
)

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
	return d.Ctx.Styles.Sizes.SideList.Height
}

func (d Delegate) Width() int {
	return d.Ctx.Styles.Sizes.SideList.Width
}

func (d Delegate) Spacing() int {
	return d.Ctx.Styles.Sizes.SideList.Spacing
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
		userName    string
		gameName    string
		viewerCount string
		s           = &d.Ctx.Styles
		innerWidth  = s.Sizes.SideList.Inner.Width
	)

	// viewerCount = strconv.Itoa(i.ViewerCount)
	viewerCount = " " + utils.Humanize(i.ViewerCount)
	userName = ansi.Truncate(i.UserName, innerWidth-lipgloss.Width(viewerCount), "…")
	gameName = ansi.Truncate(i.GameName, innerWidth, "…")

	isSelected := index == m.Index()

	var (
		topLeft   string
		topRight  string
		bottomRow string
	)

	if isSelected && d.Ctx.States.SideList.Focused {
		topRight = s.SideList.Selected.Top.Right.Width(lipgloss.Width(viewerCount)).Render(viewerCount)
		topLeft = s.SideList.Selected.Top.Left.Width(innerWidth - lipgloss.Width(topRight)).Render(userName)
		bottomRow = s.SideList.Selected.Bottom.Width(innerWidth).Render(gameName)
	} else {
		topRight = s.SideList.Normal.Top.Right.Width(lipgloss.Width(viewerCount)).Render(viewerCount)
		topLeft = s.SideList.Normal.Top.Left.Width(innerWidth - lipgloss.Width(topRight)).Render(userName)
		bottomRow = s.SideList.Normal.Bottom.Width(innerWidth).Render(gameName)
	}

	// stream := lipgloss.JoinVertical(lipgloss.Left, lipgloss.JoinHorizontal(lipgloss.Left, topLeft, topRight), bottomRow)
	stream := s.SideList.Style.Render(lipgloss.JoinVertical(lipgloss.Left, lipgloss.JoinHorizontal(lipgloss.Left, topLeft, topRight), bottomRow))

	fmt.Fprint(w, stream)
}
