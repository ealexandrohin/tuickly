package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gradient(base lipgloss.Style, s string) string {
	startColor, _ := colorful.Hex("#ffbf00")
	endColor, _ := colorful.Hex("#9146ff")
	n := len([]rune(s))
	var str string
	for i, ss := range []rune(s) {
		t := float64(i) / float64(max(1, n-1))
		c := startColor.BlendLab(endColor, t).Clamped()
		str += base.Foreground(lipgloss.Color(c.Hex())).Render(string(ss))
	}
	return str
}
