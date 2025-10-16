package enemy

import (
	"math/rand"

	game_models "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type EnemyFactory struct {
	injector models.IAppInjector
}

func NewEnemyFactory(injector models.IAppInjector) *EnemyFactory {
	return &EnemyFactory{injector: injector}
}

// @TODO add different enemies
func (ef *EnemyFactory) SpawnEnemy() game_models.IEnemy {
	windowSize := ef.injector.Storage().WindowSize()

	x := int16(rand.Intn(windowSize.Width)-5) + 2
	y := int16(0)

	return NewOctopusEnemy(x, y)
}

var _ game_models.IEnemyFactory = (*EnemyFactory)(nil)
