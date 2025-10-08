package pages

import (
	"syscall"

	pages_styles "github.com/pseudoelement/galaga/src/menu/pages/styles"
	"github.com/pseudoelement/galaga/src/menu/ui"
	"github.com/pseudoelement/galaga/src/models"
	"github.com/pseudoelement/galaga/src/storage"
)

func NewMenuRedirectButton(text string, redirectCallback func() models.IPage) models.IRedirectableElement {
	return ui.NewRedirectableElement(models.RedirectableElementParams{
		BaseViewElementParams: models.BaseViewElementParams{
			Text:         text,
			InitialStyle: pages_styles.BoldTextStyle,
		},
		RedirectCallback: redirectCallback,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////

type QuitButton struct {
	*ui.SelectableElement
}

func NewQuitButton() *QuitButton {
	return &QuitButton{SelectableElement: ui.NewSelectableElement(models.BaseViewElementParams{
		InitialStyle: pages_styles.BoldTextStyle,
		Text:         "Quit",
	})}
}

func (qb *QuitButton) Action() {
	syscall.Exit(1)
}

var _ models.IActionElement = (*QuitButton)(nil)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////

type LanguageButton struct {
	*ui.RedirectableElement
	callback func()
}

func NewLanguageButton(redirectableElParams models.RedirectableElementParams, callback func()) *LanguageButton {
	return &LanguageButton{
		RedirectableElement: ui.NewRedirectableElement(redirectableElParams),
		callback:            callback,
	}
}

func (lb *LanguageButton) Action() {
	lb.callback()
}

var _ models.IActionElement = (*LanguageButton)(nil)
var _ models.IRedirectableElement = (*LanguageButton)(nil)

func NewEnglishLanguageButton(injector models.IAppInjector) *LanguageButton {
	return NewLanguageButton(
		models.RedirectableElementParams{
			BaseViewElementParams: models.BaseViewElementParams{
				Injector:     injector,
				Text:         storage.LanguageMap[storage.EN],
				InitialStyle: pages_styles.BoldTextStyle,
			},
			RedirectCallback: func() models.IPage { return NewPageFirst(injector) },
		},
		func() { injector.Storage().SetLanguage(storage.EN) },
	)
}

func NewRussianLanguageButton(injector models.IAppInjector) *LanguageButton {
	return NewLanguageButton(
		models.RedirectableElementParams{
			BaseViewElementParams: models.BaseViewElementParams{
				Injector:     injector,
				Text:         storage.LanguageMap[storage.RU],
				InitialStyle: pages_styles.BoldTextStyle,
			},
			RedirectCallback: func() models.IPage { return NewPageFirst(injector) },
		},
		func() { injector.Storage().SetLanguage(storage.RU) },
	)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
