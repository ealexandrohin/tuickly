package utils

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func Gradient(base lipgloss.Style, s string) string {
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

func Paginate[T any](
	request func(after string) ([]T, string, error),
) ([]T, error) {
	var (
		all    []T
		cursor string
	)

	for {
		items, nextCursor, err := request(cursor)
		if err != nil {
			return nil, err
		}

		all = append(all, items...)
		if nextCursor == "" {
			break
		}
		cursor = nextCursor
	}

	return all, nil
}

func ConvertToItems[S any](sourceSlice []S, mapper func(item S) list.Item) []list.Item {
	items := make([]list.Item, len(sourceSlice))

	for i, s := range sourceSlice {
		items[i] = mapper(s)
	}

	return items
}
