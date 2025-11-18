package game_objects

import (
	game_constants "github.com/pseudoelement/galaga/src/game/game-constants"
)

func IsEnemy(gameObjectName string) bool {
	for _, enemyName := range game_constants.EMEMY_NAMES {
		if enemyName == gameObjectName {
			return true
		}
	}
	return false
}

func IsPlayer(gameObjectName string) bool {
	for _, enemyName := range game_constants.PLAYER_NAMES {
		if enemyName == gameObjectName {
			return true
		}
	}
	return false
}
