package consts

type Language int

const (
	EN Language = iota
	RU
)

var LANGUAGE_TO_TRANSLATE_MAP = map[Language]string{
	EN: "english",
	RU: "russian",
}
