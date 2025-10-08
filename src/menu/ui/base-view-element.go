package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/pseudoelement/galaga/src/models"
)

type BaseViewElement struct {
	el           lipgloss.Style
	initialValue string
}

func NewBaseViewElement(params models.BaseViewElementParams) *BaseViewElement {
	el := params.InitialStyle.SetString(params.Text)
	return &BaseViewElement{el: el, initialValue: params.Text}
}

func (se *BaseViewElement) LipglossElement() lipgloss.Style {
	return se.el
}

func (se *BaseViewElement) UpdateStyle(newStyle lipgloss.Style) {
	se.el = newStyle
}

func (se *BaseViewElement) UpdateText(text string) {
	se.el = se.el.SetString(text)
}

func (se *BaseViewElement) ViewString() string {
	return se.el.Render()
}

func (se *BaseViewElement) Value() string {
	return se.el.Value()
}

func (se *BaseViewElement) InitialValue() string {
	return se.initialValue
}

var _ models.IBaseViewElement = (*BaseViewElement)(nil)
