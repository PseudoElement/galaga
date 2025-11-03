package pages

import (
	"syscall"

	consts "github.com/pseudoelement/galaga/src/constants"
	pages_styles "github.com/pseudoelement/galaga/src/menu/pages/styles"
	"github.com/pseudoelement/galaga/src/menu/ui"
	"github.com/pseudoelement/galaga/src/models"
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

func NewQuitButton(injector models.IAppInjector) *QuitButton {
	return &QuitButton{SelectableElement: ui.NewSelectableElement(models.BaseViewElementParams{
		InitialStyle: pages_styles.BoldTextStyle,
		Text:         injector.LanguageSrv().Translate("menu.buttons.quit"),
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
				Text:         injector.LanguageSrv().Translate("menu.language.english"),
				InitialStyle: pages_styles.BoldTextStyle,
			},
			RedirectCallback: func() models.IPage { return NewPageFirst(injector) },
		},
		func() { injector.Storage().SetLanguage(consts.EN) },
	)
}

func NewRussianLanguageButton(injector models.IAppInjector) *LanguageButton {
	return NewLanguageButton(
		models.RedirectableElementParams{
			BaseViewElementParams: models.BaseViewElementParams{
				Injector:     injector,
				Text:         injector.LanguageSrv().Translate("menu.language.russian"),
				InitialStyle: pages_styles.BoldTextStyle,
			},
			RedirectCallback: func() models.IPage { return NewPageFirst(injector) },
		},
		func() { injector.Storage().SetLanguage(consts.RU) },
	)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////

type DifficultyButton struct {
	*ui.RedirectableElement
	callback func()
}

func NewDifficultyButton(redirectableElParams models.RedirectableElementParams, callback func()) *DifficultyButton {
	return &DifficultyButton{
		RedirectableElement: ui.NewRedirectableElement(redirectableElParams),
		callback:            callback,
	}
}

func (lb *DifficultyButton) Action() {
	lb.callback()
}

var _ models.IActionElement = (*DifficultyButton)(nil)
var _ models.IRedirectableElement = (*DifficultyButton)(nil)

func NewEasyDifficultyButton(injector models.IAppInjector) *DifficultyButton {
	return NewDifficultyButton(
		models.RedirectableElementParams{
			BaseViewElementParams: models.BaseViewElementParams{
				Injector:     injector,
				Text:         injector.LanguageSrv().Translate("menu.difficulty.easy"),
				InitialStyle: pages_styles.BoldTextStyle,
			},
			RedirectCallback: func() models.IPage { return NewPageFirst(injector) },
		},
		func() { injector.Storage().SetGameDifficulty(consts.EASY) },
	)
}

func NewMediumDifficultyButton(injector models.IAppInjector) *DifficultyButton {
	return NewDifficultyButton(
		models.RedirectableElementParams{
			BaseViewElementParams: models.BaseViewElementParams{
				Injector:     injector,
				Text:         injector.LanguageSrv().Translate("menu.difficulty.medium"),
				InitialStyle: pages_styles.BoldTextStyle,
			},
			RedirectCallback: func() models.IPage { return NewPageFirst(injector) },
		},
		func() { injector.Storage().SetGameDifficulty(consts.MEDIUM) },
	)
}

func NewHardDifficultyButton(injector models.IAppInjector) *DifficultyButton {
	return NewDifficultyButton(
		models.RedirectableElementParams{
			BaseViewElementParams: models.BaseViewElementParams{
				Injector:     injector,
				Text:         injector.LanguageSrv().Translate("menu.difficulty.hard"),
				InitialStyle: pages_styles.BoldTextStyle,
			},
			RedirectCallback: func() models.IPage { return NewPageFirst(injector) },
		},
		func() { injector.Storage().SetGameDifficulty(consts.HARD) },
	)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
