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

const (
	DEFAULT_PLAYER BulletType = iota
	DOUBLE_GUN_PLAYER
)

type IAppFactories interface {
	PlayerFactory(coords game_models.Coords, bulletType BulletType) game_models.IPlayer
	EnemyFactory(diffLevel game_constants.DifficultyLevel) game_models.IEnemy
	BoostFactory(diffLevel game_constants.DifficultyLevel) game_models.IBoost
}
