package storage

import (
	consts "github.com/pseudoelement/galaga/src/constants"
	"github.com/pseudoelement/galaga/src/models"
)

type AppStorage struct {
	language       consts.Language
	windowSize     models.WindowSize
	gameDifficulty consts.DifficultyLevel
}

func NewAppStorage() *AppStorage {
	return &AppStorage{
		language:       consts.EN,
		gameDifficulty: consts.EASY,
	}
}

func (as *AppStorage) Language() consts.Language {
	return as.language
}

func (as *AppStorage) SetLanguage(lang consts.Language) {
	as.language = lang
}

func (as *AppStorage) WindowSize() models.WindowSize {
	return as.windowSize
}

func (as *AppStorage) SetWindowSize(size models.WindowSize) {
	as.windowSize = size
}

func (as *AppStorage) GameDifficulty() consts.DifficultyLevel {
	return as.gameDifficulty
}

func (as *AppStorage) SetGameDifficulty(gameDifficulty consts.DifficultyLevel) {
	as.gameDifficulty = gameDifficulty
}
