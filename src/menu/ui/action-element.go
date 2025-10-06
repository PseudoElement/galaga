package ui

import "github.com/pseudoelement/galaga/src/menu/models"

type ActionElement struct {
	*SelectableElement
	pageToRedirect models.IPage
}

func NewActionElement(params models.BaseViewElementParams) *ActionElement {
	return &ActionElement{}
}

func (re *ActionElement) Action() {

}

var _ models.IActionElement = (*ActionElement)(nil)
