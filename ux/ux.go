package ux

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/ealexandrohin/tuickly/ctx"
	"github.com/ealexandrohin/tuickly/ux/pacman"
)

type UX struct {
	Ctx *ctx.Ctx

	Pacman     spinner.Model
	SideList   list.Model
	StreamList list.Model
}

func New(ctx *ctx.Ctx) UX {
	return UX{
		Ctx:    ctx,
		Pacman: spinner.New(spinner.WithSpinner(pacman.Pacman)),
	}
}
