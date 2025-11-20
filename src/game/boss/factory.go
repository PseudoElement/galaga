package boss

import (
	"math"
	"math/rand"

	consts "github.com/pseudoelement/galaga/src/constants"
	game_constants "github.com/pseudoelement/galaga/src/game/game-constants"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

var _bossCounter int = 0

func CreateBoss(diffLevel consts.DifficultyLevel, injector models.IAppInjector) g_m.IBossEnemy {
	windowSize := injector.Storage().WindowSize()
	x := int16(math.Floor(float64(rand.Intn(windowSize.Width)) / 2))
	y := int16(1)

	bossName := selectNextBossName()
	switch bossName {
	case game_constants.BOSS_JUGGERNAUT:
		return NewJuggernautBoss(x, y, getJuggernautHP(diffLevel), injector)
	case game_constants.BOSS_LAZER:
		return NewLazerBoss(x, y, getLazerHP(diffLevel), injector)
	default:
		panic(bossName + " is invalid.")
	}
}

func selectNextBossName() game_constants.EnemyType {
	bossIdx := _bossCounter % len(game_constants.BOSS_NAMES)
	_bossCounter++

	return game_constants.BOSS_NAMES[bossIdx]
}

func getJuggernautHP(diffLevel consts.DifficultyLevel) int16 {
	switch diffLevel {
	case consts.EASY:
		return 400
	case consts.MEDIUM:
		return 700
	case consts.HARD:
		return 1000
	default:
		return 400
	}
}

func getLazerHP(diffLevel consts.DifficultyLevel) int16 {
	switch diffLevel {
	case consts.EASY:
		return 100
	case consts.MEDIUM:
		return 200
	case consts.HARD:
		return 300
	default:
		return 300
	}
}
