package pages

import (
	pages_styles "github.com/pseudoelement/galaga/src/menu/pages/styles"
	"github.com/pseudoelement/galaga/src/menu/ui"
	"github.com/pseudoelement/galaga/src/models"
)

type PageControls struct {
	*Page
}

func NewPageControls(injector models.IAppInjector) models.IPage {
	p := &PageControls{Page: NewPage(injector)}

	// init list
	p.elementsList.PushBack(ui.NewBaseViewElement(models.BaseViewElementParams{
		InitialStyle: pages_styles.AccentTextStyle.Width(30),
		Text: `Galaga rules
* Move top - W
* Move down - S
* Move left - A
* Move right - D
* Quit game - Escape or Enter
`,
	}))
	p.elementsList.PushBack(NewMenuRedirectButton("Back", func() models.IPage { return NewPageFirst(injector) }))
	p.SelectFirstSelectable()

	return p
}
