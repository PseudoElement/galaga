package storage

import "github.com/pseudoelement/galaga/src/menu/models"

type AppStorage struct {
	language models.Language
}

func NewAppStorage() *AppStorage {
	return &AppStorage{language: EN}
}

func (as *AppStorage) SetLanguage(lang models.Language) {
	as.language = lang
}

func (as *AppStorage) Language() models.Language {
	return as.language
}
