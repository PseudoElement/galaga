package pages

import "github.com/pseudoelement/galaga/src/models"

type PageFirst struct {
	*Page
}

func NewPageFirst(injector models.IAppInjector) models.IPage {
	p := &PageFirst{Page: NewPage(injector)}

	// init list
	p.elementsList.PushBack(
		NewMenuRedirectButton(
			injector.LanguageSrv().Translate("menu.buttons.start"),
			func() models.IPage { return NewPageGame(injector) },
		),
	)
	p.elementsList.PushBack(
		NewMenuRedirectButton(
			injector.LanguageSrv().Translate("menu.buttons.difficulty"),
			func() models.IPage { return NewPageDifficultyLvl(injector) },
		),
	)
	p.elementsList.PushBack(
		NewMenuRedirectButton(
			injector.LanguageSrv().Translate("menu.buttons.language"),
			func() models.IPage { return NewPageLanguage(injector) },
		),
	)
	p.elementsList.PushBack(
		NewMenuRedirectButton(
			injector.LanguageSrv().Translate("menu.buttons.rules"),
			func() models.IPage { return NewPageRules(injector) },
		),
	)
	p.elementsList.PushBack(
		NewMenuRedirectButton(
			injector.LanguageSrv().Translate("menu.buttons.controls"),
			func() models.IPage { return NewPageControls(injector) },
		),
	)
	p.elementsList.PushBack(NewQuitButton(injector))

	p.SelectFirstSelectable()

	return p
}
