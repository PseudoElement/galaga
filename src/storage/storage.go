package storage

import "github.com/pseudoelement/galaga/src/models"

type AppStorage struct {
	language   models.Language
	windowSize models.WindowSize
}

func NewAppStorage() *AppStorage {
	return &AppStorage{language: EN}
}

func (as *AppStorage) Language() models.Language {
	return as.language
}

func (as *AppStorage) SetLanguage(lang models.Language) {
	as.language = lang
}

func (as *AppStorage) WindowSize() models.WindowSize {
	return as.windowSize
}

func (as *AppStorage) SetWindowSize(size models.WindowSize) {
	as.windowSize = size
}
