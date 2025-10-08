package models

type Language int

type IAppStorage interface {
	Language() Language
	SetLanguage(lang Language)
	WindowSize() WindowSize
	SetWindowSize(size WindowSize)
}

type WindowSize struct {
	Width  int
	Height int
}
