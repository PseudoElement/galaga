package boss

import (
	"math"
	"math/rand"

	consts "github.com/pseudoelement/galaga/src/constants"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

func CreateBoss(diffLevel consts.DifficultyLevel, injector models.IAppInjector) g_m.IBossEnemy {
	windowSize := injector.Storage().WindowSize()

	x := int16(math.Floor(float64(rand.Intn(windowSize.Width)) / 2))
	y := int16(0)

	var health int16
	if diffLevel == consts.EASY {
		health = 100
	} else if diffLevel == consts.MEDIUM {
		health = 200
	} else {
		health = 500
	}

	return NewJuggernautBoss(x, y, health, injector)
}
