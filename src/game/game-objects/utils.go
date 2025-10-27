package game_objects

import (
	game_constants "github.com/pseudoelement/galaga/src/game/game-constants"
	game_models "github.com/pseudoelement/galaga/src/game/models"
)

func IsEnemy(gameObject game_models.IGameObject) bool {
	for _, enemyName := range game_constants.EMEMY_NAMES {
		if enemyName == gameObject.Name() {
			return true
		}
	}
	return false
}

func IsPlayer(gameObject game_models.IGameObject) bool {
	for _, enemyName := range game_constants.PLAYER_NAMES {
		if enemyName == gameObject.Name() {
			return true
		}
	}
	return false
}
