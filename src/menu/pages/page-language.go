package pages

import (
	"fmt"

	pages_styles "github.com/pseudoelement/galaga/src/menu/pages/styles"
	"github.com/pseudoelement/galaga/src/menu/ui"
	"github.com/pseudoelement/galaga/src/models"
	"github.com/pseudoelement/galaga/src/storage"
)

type PageLanguage struct {
	*Page
}

func NewPageLanguage(injector models.IAppInjector) models.IPage {
	p := &PageLanguage{Page: NewPage(injector)}
	langStr := storage.LanguageMap[injector.Storage().Language()]

	// init list
	p.elementsList.PushBack(ui.NewBaseViewElement(models.BaseViewElementParams{
		InitialStyle: pages_styles.BoldTextStyle,
		Text:         `Select language`,
	}))
	p.elementsList.PushBack(ui.NewBaseViewElement(models.BaseViewElementParams{
		InitialStyle: pages_styles.AccentTextStyle,
		Text:         fmt.Sprintf("Current language: %v\n", langStr),
	}))
	p.elementsList.PushBack(NewEnglishLanguageButton(p.injector))
	p.elementsList.PushBack(NewRussianLanguageButton(p.injector))
	p.elementsList.PushBack(NewMenuRedirectButton("Back", func() models.IPage { return NewPageFirst(injector) }))
	p.SelectFirstSelectable()

	return p
}
