package ux

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/ealexandrohin/tuickly/ctx"
)

type UX struct {
	Ctx *ctx.Ctx

	StreamList struct {
		Mdl   list.Model
		Ready bool
	}

	SideList struct {
		Mdl   list.Model
		Ready bool
	}
}

func New(ctx *ctx.Ctx) UX {
	return UX{
		Ctx: ctx,
	}
}
