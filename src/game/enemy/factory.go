package enemy

import (
	"math/rand"

	g_c "github.com/pseudoelement/galaga/src/game/game-constants"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type EnemyFactory struct {
	injector models.IAppInjector
}

func NewEnemyFactory(injector models.IAppInjector) *EnemyFactory {
	return &EnemyFactory{injector: injector}
}

func (ef *EnemyFactory) SpawnEnemy(diffLevel g_c.DifficultyLevel) g_m.IEnemy {
	windowSize := ef.injector.Storage().WindowSize()

	x := int16(rand.Intn(windowSize.Width)-5) + 2
	y := int16(-2)

	var health int16
	if diffLevel == g_c.EASY {
		health = 3
	} else {
		health = 6
	}

	// randBit := rand.Intn(2)
	// if randBit == 0 {
	// 	return NewOctopusEnemy(x, y, health)
	// } else {
	// 	return NewSmallSpaceShipEnemy(x, y, health)
	// }
	return NewSmallSpaceShipEnemy(x, y, health)
}

var _ g_m.IEnemyFactory = (*EnemyFactory)(nil)
