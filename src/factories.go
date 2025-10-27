package app

import (
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

func (af *AppFactories) EnemyFactory(diffLevel game_constants.DifficultyLevel) game_models.IEnemy {
	return enemy.CreateEnemy(diffLevel, af.injector)
}

func (af *AppFactories) BoostFactory(diffLevel game_constants.DifficultyLevel) game_models.IBoost {
	return game_objects.CreateBoost(diffLevel, af.injector)
}
