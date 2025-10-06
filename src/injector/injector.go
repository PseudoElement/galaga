package injector

import (
	"github.com/pseudoelement/galaga/src/menu/models"
)

type AppInjector struct {
	storage models.IAppStorage
	view    models.IAppView
}

func NewAppInjector(storage models.IAppStorage, view models.IAppView) *AppInjector {
	injector := &AppInjector{storage: storage, view: view}
	return injector
}

func (i *AppInjector) Storage() models.IAppStorage {
	return i.storage
}
