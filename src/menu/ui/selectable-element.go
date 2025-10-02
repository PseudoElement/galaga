package ui

import "github.com/charmbracelet/lipgloss"

type SelectableElement struct {
	el           lipgloss.Style
	selected     bool
	initialValue string
}

func NewSelectableElement(text string) *SelectableElement {
	el := lipgloss.NewStyle().Height(2).Bold(true).SetString(text)
	return &SelectableElement{el: el, initialValue: text}
}

func (se *SelectableElement) LipglossElement() lipgloss.Style {
	return se.el
}

func (se *SelectableElement) UpdateStyle(newStyle lipgloss.Style) {
	se.el = newStyle
}

func (se *SelectableElement) UpdateText(text string) {
	se.el = se.el.SetString(text)
}

func (se *SelectableElement) String() string {
	return se.el.Render()
}

func (se *SelectableElement) Value() string {
	return se.el.Value()
}

func (se *SelectableElement) InitialValue() string {
	return se.initialValue
}

func (se *SelectableElement) SetSelected(selected bool) {
	se.selected = selected
}

func (se *SelectableElement) IsSelected() bool {
	return se.selected
}
