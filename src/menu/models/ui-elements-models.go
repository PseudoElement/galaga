package models

import "github.com/charmbracelet/lipgloss"

type IBaseViewElement interface {
	LipglossElement() lipgloss.Style
	UpdateStyle(newStyle lipgloss.Style)
	UpdateText(text string)
	ViewString() string
	Value() string
	InitialValue() string
}

type ISelectableElement interface {
	IBaseViewElement
	SetSelected(selected bool)
	IsSelected() bool
}

type IRedirectableElement interface {
	ISelectableElement
	PageToRedirect() IPage
}

type IActionElement interface {
	ISelectableElement
	Action()
}
