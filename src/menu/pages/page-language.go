package pages

import (
	"fmt"

	"github.com/pseudoelement/galaga/src/menu/models"
	pages_styles "github.com/pseudoelement/galaga/src/menu/pages/styles"
	"github.com/pseudoelement/galaga/src/menu/ui"
)

type PageLanguage struct {
	*Page
}

func NewPageLanguage(injector models.IAppInjector) models.IPage {
	p := &PageLanguage{Page: NewPage(injector)}

	// init list
	p.elementsList.PushBack(ui.NewBaseViewElement(models.BaseViewElementParams{
		InitialStyle: pages_styles.BoldTextStyle,
		Text:         `Select language`,
	}))
	p.elementsList.PushBack(ui.NewBaseViewElement(models.BaseViewElementParams{
		InitialStyle: pages_styles.AccentTextStyle,
		Text:         fmt.Sprintf("Current language: %v\n", injector.Storage().Language()),
	}))
	p.elementsList.PushBack(NewEnglishLanguageButton(p.injector))
	p.elementsList.PushBack(NewRussianLanguageButton(p.injector))
	p.elementsList.PushBack(NewMenuRedirectButton("Back", func() models.IPage { return NewPageFirst(injector) }))
	p.SelectFirstSelectable()

	return p
}
