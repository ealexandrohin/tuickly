package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// twitch colors
	purple = lipgloss.Color("#9146FF") // twitch primary
	black  = lipgloss.Color("#000000") // twitch black ops
	ice    = lipgloss.Color("#F0F0FF") // twitch ice

	pacman = lipgloss.Color("#FFBF00") // twitch pacman

	// base colors
	white = lipgloss.Color("#FFFFFF")

	// colors
	primary = purple
	text    = white

	// tab borders
	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "━",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "┷",
		BottomRight: "┷",
	}

	activeTabBorder = lipgloss.Border{
		Top:         "━",
		Bottom:      " ",
		Left:        "┃",
		Right:       "┃",
		TopLeft:     "┏",
		TopRight:    "┓",
		BottomLeft:  "┛",
		BottomRight: "┗",
	}

	tabGapBorder = lipgloss.Border{
		Bottom:      "━",
		BottomLeft:  "━",
		BottomRight: "━",
	}

	// tabs
	tab       = lipgloss.NewStyle().Border(tabBorder).BorderForeground(primary).Padding(0, 1)
	activeTab = tab.Border(activeTabBorder)
	tabGap    = lipgloss.NewStyle().Border(tabGapBorder).BorderForeground(primary)

	// statusbar
	status     = lipgloss.NewStyle().Foreground(lipgloss.Color(text)).Background(lipgloss.Color(primary)).Padding(0, 1)
	statusBar  = lipgloss.NewStyle().Foreground(lipgloss.Color("#C1C6B2")).Background(lipgloss.Color("#353533")).Padding(0, 1)
	statusTime = status.Align(lipgloss.Right)

	// dialog
	dialog            = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#874BFD")).Padding(1, 0)
	buttonStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFF7DB")).Background(lipgloss.Color("#888B7E")).Padding(0, 3).MarginTop(1)
	activeButtonStyle = buttonStyle.Foreground(lipgloss.Color("#FFF7DB")).Background(lipgloss.Color("#F25D94")).MarginRight(2).Underline(true)
)

func URIDialog(m Model) string {
	URI := lipgloss.NewStyle().Width(len(m.Msg.(URIMsg).URI) + 4).Align(lipgloss.Center).Render(m.Msg.(URIMsg).URI)

	return lipgloss.Place(m.Window.Width, m.Window.Height,
		lipgloss.Center, lipgloss.Center,
		dialog.Render(URI),
	)
}

// func (m Model) headerView() string {
// 	row := lipgloss.JoinHorizontal(
// 		lipgloss.Center,
// 		tab.Render("Live"),
// 		activeTab.Render("Follows"),
// 		tab.Render("ealexandrohin"),
// 	)
//
// 	gap := tabGap.Render(strings.Repeat(" ", max(0, m.Viewport.Width-lipgloss.Width(row))))
// 	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
//
// 	return row
// }
//
// func (m Model) footerView() string {
// 	tuickly := status.Render("tuickly")
// 	clock := statusTime.Render(time.Now().Format(time.TimeOnly))
// 	tag := statusBar.Width(m.Viewport.Width - lipgloss.Width(tuickly) - lipgloss.Width(clock)).Render("@ealexandrohin")
//
// 	bar := lipgloss.JoinHorizontal(lipgloss.Top, tuickly, tag, clock)
//
// 	return bar
// }
