package models

import (
	"github.com/charmbracelet/lipgloss"
)

type BaseViewElementParams struct {
	Injector     IAppInjector
	Text         string
	InitialStyle lipgloss.Style
}

type RedirectableElementParams struct {
	BaseViewElementParams
	RedirectCallback func() IPage
}

type ActionElementParams struct {
	BaseViewElementParams
	Callback func()
}
