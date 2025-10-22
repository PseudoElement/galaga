package game_models

import g_c "github.com/pseudoelement/galaga/src/game/game-constants"

type IEnemyFactory interface {
	SpawnEnemy(diffLevel g_c.DifficultyLevel) IEnemy
}
