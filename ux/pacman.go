package ux

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
)

var Pacman = spinner.Spinner{
	Frames: []string{"••••", "𜱭•••", " 𜱭••", "• 𜱭•", "•• 𜱭", "••• "},
	FPS:    time.Second / 2,
}
