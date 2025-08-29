package ux

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/ansi"
	"github.com/eAlexandrohin/tuickly/ui/styles"
)

type StreamItem struct {
	UserID      string
	UserLogin   string
	UserName    string
	GameName    string
	Title       string
	ViewerCount int
	StartedAt   time.Time
}

func (s StreamItem) FilterValue() string {
	return fmt.Sprintf("%s %s %s", s.UserLogin, s.GameName, s.Title)
}

func (s StreamItem) Get() StreamItem {
	return s
}

type StreamDelegate struct {
	Styles styles.Styles
	// UpdateFunc      func(tea.Msg, *Model) tea.Cmd
	// ShortHelpFunc   func() []key.Binding
	// FullHelpFunc    func() [][]key.Binding
}

// func (d StreamDelegate) SetHeight(h int) {
// 	d.height = h
// }

func (d StreamDelegate) Height() int {
	// return d.height
	return 2
}

// func (d StreamDelegate) SetSpacing(s int) {
// 	d.spacing = s
// }

func (d StreamDelegate) Spacing() int {
	// return d.spacing
	return 1
}

func (d StreamDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d StreamDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	var (
		userName    string
		gameName    string
		uptime      string
		viewerCount string
		title       string
		// matchedRunes []int
		s = &d.Styles
	)

	if i, ok := item.(StreamItem); ok {
		userName = i.UserName
		gameName = i.GameName

		d := time.Since(i.StartedAt)
		h := int(d / time.Hour)
		d -= time.Duration(h) * time.Hour
		m := int(d / time.Minute)
		d -= time.Duration(m) * time.Minute
		s := int(d / time.Second)
		uptime = fmt.Sprintf("%02d:%02d:%02d", h, m, s)
		// god i wish for .toLocaleTimeString()

		viewerCount = strconv.Itoa(i.ViewerCount)
		title = i.Title
	} else {
		return
	}

	if m.Width() <= 0 {
		return
	}

	textwidth := m.Width() - s.List.Normal.Title.GetPaddingLeft() - s.List.Normal.Title.GetPaddingRight()
	title = ansi.Truncate(title, textwidth, "...")

	var (
		isSelected  = index == m.Index()
		emptyFilter = m.FilterState() == list.Filtering && m.FilterValue() == ""
		// isFiltered  = m.FilterState() == list.Filtering || m.FilterState() == list.FilterApplied
	)

	// if isFiltered {
	// 	matchedRunes = m.MatchesForItem(index)
	// }

	if emptyFilter {
		userName = s.List.Dimmed.UserName.Render(userName)
		gameName = s.List.Dimmed.GameName.Render(gameName)
		uptime = s.List.Dimmed.Uptime.Render(uptime)
		viewerCount = s.List.Dimmed.ViewerCount.Render(viewerCount)
		title = s.List.Dimmed.Title.Render(title)
	} else if isSelected && m.FilterState() != list.Filtering {
		// if isFiltered {
		// 	unmatched := s.List.Selected.Style.Inline(true)
		// 	matched := unmatched.Inherit(s.List.FilterMatch)
		// 	title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
		// }
		userName = s.List.Selected.UserName.Render(userName)
		gameName = s.List.Selected.GameName.Render(gameName)
		uptime = s.List.Selected.Uptime.Render(uptime)
		viewerCount = s.List.Selected.ViewerCount.Render(viewerCount)
		title = s.List.Selected.Title.Render(title)
	} else {
		// if isFiltered {
		// 	// Highlight matches
		// 	unmatched := s.NormalTitle.Inline(true)
		// 	matched := unmatched.Inherit(s.FilterMatch)
		// 	title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
		// }
		userName = s.List.Normal.UserName.Render(userName)
		gameName = s.List.Normal.GameName.Render(gameName)
		uptime = s.List.Normal.Uptime.Render(uptime)
		viewerCount = s.List.Normal.ViewerCount.Render(viewerCount)
		title = s.List.Normal.Title.Render(title)
	}

	// fmt.Fprintf(w, "%s\n%s", title, desc) //nolint: errcheck
	fmt.Sprintf("%v --- %v --- %v --- %v\n\t>>>%v", userName, gameName, uptime, viewerCount, title)

	// i, ok := item.(StreamItem)
	// if !ok {
	// 	return
	// }
	//
	// var style lipgloss.Style
	// if index == m.Index() {
	// 	style = lipgloss.NewStyle().
	// 		Background(lipgloss.Color("205")).
	// 		Foreground(lipgloss.Color("235"))
	// } else {
	// 	style = lipgloss.NewStyle().
	// 		Foreground(lipgloss.Color("250"))
	// }
	//
	// str := fmt.Sprintf("%v --- %v --- %v --- %v\n\t>>>%v", i.UserName, i.GameName, i.StartedAt, i.ViewerCount, i.Title)
	//
	// fmt.Fprint(w, style.Render(str))
}
