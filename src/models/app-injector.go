package models

import tea "github.com/charmbracelet/bubbletea"

type IAppInjector interface {
	Storage() IAppStorage
	LanguageSrv() IAppLanguageSrv
	GameSrv() IAppGameSrv
	TeaProgram() *tea.Program
	Factories() IAppFactories
	View() IAppView
}
