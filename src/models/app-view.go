package models

type IAppView interface {
	SetPage(page IPage)
	View() string
	CurrentPage() IPage
}
