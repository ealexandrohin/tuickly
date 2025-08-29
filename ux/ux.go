package ux

import "github.com/charmbracelet/bubbles/list"

type UX struct {
	List struct {
		Mdl   list.Model
		Ready bool
	}
}
