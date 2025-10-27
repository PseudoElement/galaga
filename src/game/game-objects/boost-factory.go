package game_objects

import (
	"math/rand"

	g_c "github.com/pseudoelement/galaga/src/game/game-constants"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

func CreateBoost(diffLevel g_c.DifficultyLevel, injector models.IAppInjector) g_m.IBoost {
	windowSize := injector.Storage().WindowSize()

	x := int16(rand.Intn(windowSize.Width)-5) + 2
	y := int16(-2)

	randBit := rand.Intn(2)
	if randBit == 0 {
		return NewDoubleGunShipBoost(g_m.Coords{X: x, Y: y}, injector)
	} else {
		return NewTripleGunShipBoost(g_m.Coords{X: x, Y: y}, injector)
	}
}
