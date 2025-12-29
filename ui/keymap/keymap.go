package keymap

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
}

func New() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("ctrl+up", "ctrl+k"),
			key.WithHelp("^↑/^k", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("ctrl+down", "ctrl+j"),
			key.WithHelp("^↓/^j", "down"),
		),
		Left: key.NewBinding(
			key.WithKeys("ctrl+left", "ctrl+h"),
			key.WithHelp("^←/^h", "left"),
		),
		Right: key.NewBinding(
			key.WithKeys("ctrl+right", "ctrl+l"),
			key.WithHelp("^→/^l", "right"),
		),
	}
}
