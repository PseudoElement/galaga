package pages

import (
	"fmt"

	consts "github.com/pseudoelement/galaga/src/constants"
	pages_styles "github.com/pseudoelement/galaga/src/menu/pages/styles"
	"github.com/pseudoelement/galaga/src/menu/ui"
	"github.com/pseudoelement/galaga/src/models"
)

type PageLanguage struct {
	*Page
}

func NewPageLanguage(injector models.IAppInjector) models.IPage {
	p := &PageLanguage{Page: NewPage(injector)}
	langStr := consts.LanguageMap[injector.Storage().Language()]

	// init list
	p.elementsList.PushBack(ui.NewBaseViewElement(models.BaseViewElementParams{
		InitialStyle: pages_styles.BoldTextStyle,
		Text:         injector.LanguageSrv().Translate("menu.language.selectLang"),
	}))
	p.elementsList.PushBack(ui.NewBaseViewElement(models.BaseViewElementParams{
		InitialStyle: pages_styles.AccentTextStyle,
		Text:         fmt.Sprintf("%v: %v\n", injector.LanguageSrv().Translate("menu.language.currLang"), langStr),
	}))
	p.elementsList.PushBack(NewEnglishLanguageButton(p.injector))
	p.elementsList.PushBack(NewRussianLanguageButton(p.injector))
	p.elementsList.PushBack(NewMenuRedirectButton(
		injector.LanguageSrv().Translate("menu.buttons.back"),
		func() models.IPage { return NewPageFirst(injector) },
	),
	)
	p.SelectFirstSelectable()

	return p
}
