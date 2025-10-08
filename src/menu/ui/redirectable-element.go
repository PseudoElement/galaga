package ui

import "github.com/pseudoelement/galaga/src/models"

type RedirectableElement struct {
	*SelectableElement
	redirectCallback func() models.IPage
}

func NewRedirectableElement(params models.RedirectableElementParams) *RedirectableElement {
	return &RedirectableElement{
		SelectableElement: NewSelectableElement(params.BaseViewElementParams),
		redirectCallback:  params.RedirectCallback,
	}
}

func (re *RedirectableElement) PageToRedirect() models.IPage {
	return re.redirectCallback()
}

var _ models.IBaseViewElement = (*RedirectableElement)(nil)
var _ models.ISelectableElement = (*RedirectableElement)(nil)
var _ models.IRedirectableElement = (*RedirectableElement)(nil)
