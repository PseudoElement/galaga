package models

type IAppInjector interface {
	Storage() IAppStorage
}
