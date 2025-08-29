package ui

// func (ui UI) HeaderView() string {
// 	row := lipgloss.JoinHorizontal(
// 		lipgloss.Center,
// 		activeTabStyle.Render("Live"),
// 		tabStyle.Render("Follows"),
// 		tabStyle.Render("ealexandrohin"),
// 	)
//
// 	gap := tabGapStyle.Render(strings.Repeat(" ", max(0, ui.Window.Width-lipgloss.Width(row))))
// 	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
//
// 	return row
// }
//
// func (ui UI) ContentView() string {
// 	return ""
// }
//
// func (ui UI) FooterView() string {
// 	tuickly := statusStyle.Render("tuickly")
// 	clock := statusTimeStyle.Render(time.Now().Format(time.TimeOnly))
// 	tag := statusBarStyle.Width(ui.Window.Width - lipgloss.Width(tuickly) - lipgloss.Width(clock)).Render("@ealexandrohin")
//
// 	bar := lipgloss.JoinHorizontal(lipgloss.Top, tuickly, tag, clock)
//
// 	return bar
// }
