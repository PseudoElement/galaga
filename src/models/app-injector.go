package models

type IAppInjector interface {
	Storage() IAppStorage
	LanguageSrv() IAppLanguageSrv
}
