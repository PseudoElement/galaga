package models

import (
	game_constants "github.com/pseudoelement/galaga/src/game/game-constants"
	game_models "github.com/pseudoelement/galaga/src/game/models"
)

type BulletType = int
type PlayerType = int

const (
	PLAYER_BULLET BulletType = iota
	ENEMY_BULLET
)

type IAppFactories interface {
	PlayerFactory(coords game_models.Coords, playerType game_constants.PlayerType) game_models.IPlayer
	EnemyFactory(diffLevel game_constants.DifficultyLevel) game_models.IEnemy
	BoostFactory(diffLevel game_constants.DifficultyLevel) game_models.IBoost
}
