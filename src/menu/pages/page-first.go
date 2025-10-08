package pages

import "github.com/pseudoelement/galaga/src/models"

type PageFirst struct {
	*Page
}

func NewPageFirst(injector models.IAppInjector) models.IPage {
	p := &PageFirst{Page: NewPage(injector)}

	// init list
	p.elementsList.PushBack(NewMenuRedirectButton("Start", func() models.IPage { return NewPageRules(injector) }))
	p.elementsList.PushBack(NewMenuRedirectButton("Language", func() models.IPage { return NewPageLanguage(injector) }))
	p.elementsList.PushBack(NewMenuRedirectButton("Rules", func() models.IPage { return NewPageRules(injector) }))
	p.elementsList.PushBack(NewMenuRedirectButton("Controls", func() models.IPage { return NewPageControls(injector) }))
	p.elementsList.PushBack(NewQuitButton())
	p.SelectFirstSelectable()

	return p
}
