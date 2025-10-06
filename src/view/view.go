package app_view

import "github.com/pseudoelement/galaga/src/menu/models"

type AppView struct {
	page models.IPage
}

func NewAppView(page models.IPage) *AppView {
	return &AppView{page: page}
}

func (av *AppView) SetPage(page models.IPage) {
	av.page = page
}

func (av *AppView) View() string {
	return av.page.View()
}

func (av *AppView) CurrentPage() models.IPage {
	return av.page
}
