package game_models

import consts "github.com/pseudoelement/galaga/src/constants"

type IEnemyFactory interface {
	CreateEnemy(diffLevel consts.DifficultyLevel) IEnemy
}
