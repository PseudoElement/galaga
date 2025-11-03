package models

import consts "github.com/pseudoelement/galaga/src/constants"

type IAppStorage interface {
	Language() consts.Language
	SetLanguage(lang consts.Language)
	WindowSize() WindowSize
	SetWindowSize(size WindowSize)
	GameDifficulty() consts.DifficultyLevel
	SetGameDifficulty(diffilculty consts.DifficultyLevel)
}

type WindowSize struct {
	Width  int
	Height int
}
