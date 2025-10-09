package consts

import "github.com/pseudoelement/galaga/src/models"

const (
	EN models.Language = iota
	RU
)

var LanguageMap = map[models.Language]string{
	EN: "English",
	RU: "Russian",
}
