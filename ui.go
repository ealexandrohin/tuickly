package main

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
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
	primary = pacman
	text    = white

	// spinner
	spinnerStyle = lipgloss.NewStyle().Foreground(primary)

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
	tabStyle       = lipgloss.NewStyle().Border(tabBorder).BorderForeground(primary).Padding(0, 1)
	activeTabStyle = tabStyle.Border(activeTabBorder)
	tabGapStyle    = lipgloss.NewStyle().Border(tabGapBorder).BorderForeground(primary)

	// statusbar
	statusStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color(text)).Background(lipgloss.Color(primary)).Padding(0, 1)
	statusBarStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#C1C6B2")).Background(lipgloss.Color("#353533")).Padding(0, 1)
	statusTimeStyle = statusStyle.Align(lipgloss.Right)

	// dialog
	dialogStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color(primary)).Padding(1, 0)
)

type Tab struct {
	Title   string
	Content string
}

type Window struct {
	Width  int
	Height int
}

type UI struct {
	spinner.Model
}

func URIDialog(m Model) string {
	URI := lipgloss.NewStyle().Width(len(m.Msg.(URIMsg).URI) + 4).Align(lipgloss.Center).Render(m.Msg.(URIMsg).URI)

	return lipgloss.Place(m.Window.Width, m.Window.Height,
		lipgloss.Center, lipgloss.Center,
		dialogStyle.Render(URI),
	)
}

func (m Model) headerView() string {
	row := lipgloss.JoinHorizontal(
		lipgloss.Center,
		activeTabStyle.Render("Live"),
		tabStyle.Render("Follows"),
		tabStyle.Render("ealexandrohin"),
	)

	gap := tabGapStyle.Render(strings.Repeat(" ", max(0, m.Window.Width-lipgloss.Width(row))))
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)

	return row
}

func (m Model) contentView() string {
	return ""
	// return docStyle.Render(m.list.View())
}

func (m Model) footerView() string {
	tuickly := statusStyle.Render("tuickly")
	clock := statusTimeStyle.Render(time.Now().Format(time.TimeOnly))
	tag := statusBarStyle.Width(m.Window.Width - lipgloss.Width(tuickly) - lipgloss.Width(clock)).Render("@ealexandrohin")

	bar := lipgloss.JoinHorizontal(lipgloss.Top, tuickly, tag, clock)

	return bar
}
