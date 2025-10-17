package sidelist

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
	)

	userName = ansi.Truncate(i.UserName, d.Width()/2, "…")
	viewerCount = strconv.Itoa(i.ViewerCount)
	gameName = ansi.Truncate(i.GameName, d.Width(), "…")

	isSelected := index == m.Index()

	var (
		topLeft   string
		topRight  string
		bottomRow string
	)

	if isSelected {
		topLeft = s.SideList.Selected.Top.Left.Render(userName)
		topRight = s.SideList.Selected.Top.Right.Render(viewerCount)
		bottomRow = s.SideList.Selected.Bottom.Width(d.Width()).Render(gameName)
	} else {
		topLeft = s.SideList.Normal.Top.Left.Render(userName)
		topRight = s.SideList.Normal.Top.Right.Render(viewerCount)
		bottomRow = s.SideList.Normal.Bottom.Width(d.Width()).Render(gameName)
	}

	stream := lipgloss.JoinVertical(lipgloss.Left, lipgloss.JoinHorizontal(lipgloss.Left, topLeft, topRight), bottomRow)

	fmt.Fprint(w, stream)
}
