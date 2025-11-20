package enemy

import (
	"math/rand"

	consts "github.com/pseudoelement/galaga/src/constants"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

func CreateEnemy(diffLevel consts.DifficultyLevel, injector models.IAppInjector) g_m.IEnemy {
	windowSize := injector.Storage().WindowSize()

	x := int16(rand.Intn(windowSize.Width)-10) + 5
	y := int16(-2)

	var health int16
	if diffLevel == consts.EASY {
		health = 3
	} else if diffLevel == consts.MEDIUM {
		health = 5
	} else {
		health = 7
	}

	randBit := rand.Intn(10)
	if randBit >= 5 {
		return NewOctopusEnemy(x, y, health)
	} else if randBit >= 2 {
		return NewSmallSpaceShipEnemy(x, y, health)
	} else {
		return NewTNTEnemy(x, 0, health, injector)
	}
}
