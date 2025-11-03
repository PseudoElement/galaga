package consts

type DifficultyLevel = int

const (
	EASY = iota
	MEDIUM
	HARD
)

var DIFFICULTY_TO_TRANSLATE_MAP = map[DifficultyLevel]string{
	EASY:   "easy",
	MEDIUM: "medium",
	HARD:   "hard",
}
