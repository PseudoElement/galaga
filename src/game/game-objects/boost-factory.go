package game_objects

import (
	"math/rand"

	consts "github.com/pseudoelement/galaga/src/constants"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

func CreateBoost(diffLevel consts.DifficultyLevel, spawnNewShip bool, injector models.IAppInjector) g_m.IBoost {
	windowSize := injector.Storage().WindowSize()

	x := int16(rand.Intn(windowSize.Width)-10) + 5
	y := int16(-2)

	if spawnNewShip {
		playerTier := injector.GameSrv().Player().Tier()
		if playerTier == 1 {
			return NewDoubleGunShipBoost(g_m.Coords{X: x, Y: y}, injector)
		} else {
			return NewTripleGunShipBoost(g_m.Coords{X: x, Y: y}, injector)
		}
	}

	return NewHpBoost(g_m.Coords{X: x, Y: y}, injector)
}
