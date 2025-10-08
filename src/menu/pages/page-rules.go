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

	// init list
	p.elementsList.PushBack(ui.NewBaseViewElement(models.BaseViewElementParams{
		InitialStyle: pages_styles.AccentTextStyle.Width(30),
		Text: `Galaga rules
Your main goal is to protect yourself and live as much as possible.
You can kill them or just dodge shots.
Avoid bombs(TNT), they blast with range 4*4 when you shot them.
Good luck.
`,
	}))
	p.elementsList.PushBack(NewMenuRedirectButton("Back", func() models.IPage { return NewPageFirst(injector) }))
	p.SelectFirstSelectable()

	return p
}
