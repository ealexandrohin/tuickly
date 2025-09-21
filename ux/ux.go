package ux

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/eAlexandrohin/tuickly/ctx"
)

type UX struct {
	Ctx *ctx.Ctx

	List struct {
		Mdl   list.Model
		Ready bool
	}
}

func New(ctx *ctx.Ctx) UX {
	return UX{
		Ctx: ctx,
	}
}
