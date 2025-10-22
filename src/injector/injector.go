package injector

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pseudoelement/galaga/src/models"
)

type AppInjector struct {
	storage     models.IAppStorage
	view        models.IAppView
	languageSrv models.IAppLanguageSrv
	gameSrv     models.IAppGameSrv
	factories   models.IAppFactories
	teaProgram  *tea.Program
}

func NewAppInjector() *AppInjector {
	return &AppInjector{}
}

func (i *AppInjector) Factories() models.IAppFactories {
	return i.factories
}

func (i *AppInjector) SetFactories(factories models.IAppFactories) {
	i.factories = factories
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

func (i *AppInjector) GameSrv() models.IAppGameSrv {
	return i.gameSrv
}

func (i *AppInjector) SetGameSrv(gameSrv models.IAppGameSrv) {
	i.gameSrv = gameSrv
}

func (i *AppInjector) TeaProgram() *tea.Program {
	return i.teaProgram
}

func (i *AppInjector) SetTeaProgram(teaProgram *tea.Program) {
	i.teaProgram = teaProgram
}
