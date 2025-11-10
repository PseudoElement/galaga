package models

import (
	consts "github.com/pseudoelement/galaga/src/constants"
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
	EnemyFactory(diffLevel consts.DifficultyLevel) game_models.IEnemy
	BoostFactory(diffLevel consts.DifficultyLevel, spawnNewShip bool) game_models.IBoost
	BossEnemyFactory(diffLevel consts.DifficultyLevel) game_models.IBossEnemy
}
