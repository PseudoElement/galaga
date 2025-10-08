package ui

import "github.com/pseudoelement/galaga/src/models"

type SelectableElement struct {
	*BaseViewElement
	selected bool
}

func NewSelectableElement(params models.BaseViewElementParams) *SelectableElement {
	return &SelectableElement{BaseViewElement: NewBaseViewElement(params)}
}

func (se *SelectableElement) SetSelected(selected bool) {
	se.selected = selected
}

func (se *SelectableElement) IsSelected() bool {
	return se.selected
}

var _ models.IBaseViewElement = (*SelectableElement)(nil)
var _ models.ISelectableElement = (*SelectableElement)(nil)
