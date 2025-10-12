package pages

import (
	pages_styles "github.com/pseudoelement/galaga/src/menu/pages/styles"
	"github.com/pseudoelement/galaga/src/menu/ui"
	"github.com/pseudoelement/galaga/src/models"
)

type PageRules struct {
	*Page
}

func NewPageRules(injector models.IAppInjector) models.IPage {
	p := &PageRules{Page: NewPage(injector)}
	windowWidth := injector.Storage().WindowSize().Width

	// init list
	p.elementsList.PushBack(ui.NewBaseViewElement(models.BaseViewElementParams{
		InitialStyle: pages_styles.AccentTextStyle.Width(windowWidth),
		Text:         injector.LanguageSrv().Translate("menu.text.rules"),
	}))
	p.elementsList.PushBack(
		NewMenuRedirectButton(
			injector.LanguageSrv().Translate("menu.buttons.back"),
			func() models.IPage { return NewPageFirst(injector) },
		),
	)

	p.SelectFirstSelectable()

	return p
}
