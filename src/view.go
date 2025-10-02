package app

import menu_types "github.com/pseudoelement/galaga/src/menu/types"

type AppView struct {
	page menu_types.IPage
}

func NewAppView() AppView {
	return AppView{}
}

func (av *AppView) SetPage(page menu_types.IPage) {
	av.page = page
}

func (av *AppView) View() string {
	return av.page.View()
}

func (av *AppView) CurrentPage() menu_types.IPage {
	return av.page
}
