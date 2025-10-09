package injector

import (
	"github.com/pseudoelement/galaga/src/models"
)

type AppInjector struct {
	storage     models.IAppStorage
	view        models.IAppView
	languageSrv models.IAppLanguageSrv
}

func NewAppInjector() *AppInjector {
	return &AppInjector{}
}

func (i *AppInjector) Storage() models.IAppStorage {
	return i.storage
}

func (i *AppInjector) SetStorage(storage models.IAppStorage) {
	i.storage = storage
}

func (i *AppInjector) LanguageSrv() models.IAppLanguageSrv {
	return i.languageSrv
}

func (i *AppInjector) SetLanguageSrv(languageSrv models.IAppLanguageSrv) {
	i.languageSrv = languageSrv
}

func (i *AppInjector) View() models.IAppView {
	return i.view
}

func (i *AppInjector) SetView(view models.IAppView) {
	i.view = view
}
