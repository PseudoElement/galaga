package app

import (
	consts "github.com/pseudoelement/galaga/src/constants"
	"github.com/pseudoelement/galaga/src/game/boss"
	"github.com/pseudoelement/galaga/src/game/enemy"
	game_constants "github.com/pseudoelement/galaga/src/game/game-constants"
	game_objects "github.com/pseudoelement/galaga/src/game/game-objects"
	game_models "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/game/player"
	"github.com/pseudoelement/galaga/src/models"
)

type AppFactories struct {
	injector models.IAppInjector
}

func NewAppFactories(injector models.IAppInjector) models.IAppFactories {
	return &AppFactories{injector: injector}
}

func (af *AppFactories) PlayerFactory(coords game_models.Coords, playerType game_constants.PlayerType) game_models.IPlayer {
	return player.CreatePlayer(coords, playerType, af.injector)
}

func (af *AppFactories) EnemyFactory(diffLevel consts.DifficultyLevel) game_models.IEnemy {
	return enemy.CreateEnemy(diffLevel, af.injector)
}

func (af *AppFactories) BoostFactory(diffLevel consts.DifficultyLevel, spawnNewShip bool) game_models.IBoost {
	return game_objects.CreateBoost(diffLevel, spawnNewShip, af.injector)
}

func (af *AppFactories) BossEnemyFactory(diffLevel consts.DifficultyLevel) game_models.IBossEnemy {
	return boss.CreateBoss(diffLevel, af.injector)
}
