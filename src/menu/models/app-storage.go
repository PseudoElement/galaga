package models

type Language int

type IAppStorage interface {
	SetLanguage(lang Language)
	Language() Language
}
