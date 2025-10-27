package game_srv

import (
	g_o "github.com/pseudoelement/galaga/src/game/game-objects"
	game_models "github.com/pseudoelement/galaga/src/game/models"
)

func handleCollisionScenarios(obj1 game_models.IGameObject, obj2 game_models.IGameObject, gameSrv *AppGameSrv) {
	_checkCollisionPlayerWithEnemy(obj1, obj2, gameSrv)
	_checkCollisionPlayerWithEnemy(obj2, obj1, gameSrv)

	_checkCollisionWithBullet(obj1, obj2, gameSrv)
	_checkCollisionWithBullet(obj2, obj1, gameSrv)

	_checkCollisionWithBoost(obj1, obj2)
	_checkCollisionWithBoost(obj2, obj1)
}

func _checkCollisionWithBullet(obj game_models.IGameObject, otherObject game_models.IGameObject, gameSrv *AppGameSrv) {
	bulletObject, isObjectBullet := obj.(game_models.IBullet)
	if isObjectBullet && !bulletObject.Destroyed() {
		playerObject, isOtherObjPlayer := otherObject.(game_models.IPlayer)
		enemyObject, isOtherObjEnemy := otherObject.(game_models.IEnemy)
		// enemy bullet should not damage other enemies
		if isOtherObjEnemy && !g_o.IsEnemy(bulletObject.Owner()) {
			bulletObject.Damage(enemyObject)
			bulletObject.Destroy()
			if enemyObject.Destroyed() {
				gameSrv.increaseScore(enemyObject.Price())
			}
			return
		}
		// player bullet should not damage player
		if isOtherObjPlayer && !g_o.IsPlayer(bulletObject.Owner()) {
			bulletObject.Damage(playerObject)
			bulletObject.Destroy()
		}
	}
}

func _checkCollisionWithBoost(obj game_models.IGameObject, otherObject game_models.IGameObject) {
	boostObject, isObjectBoost := obj.(game_models.IBoost)
	if isObjectBoost && !boostObject.Destroyed() {
		player, isPlayer := otherObject.(game_models.IPlayer)
		if isPlayer {
			boostObject.Boost(player)
			boostObject.Destroy()
		}
	}
}

func _checkCollisionPlayerWithEnemy(obj game_models.IGameObject, otherObject game_models.IGameObject, gameSrv *AppGameSrv) {
	enemyObject, isObjectEnemy := obj.(game_models.IEnemy)
	if isObjectEnemy {
		player, isPlayer := otherObject.(game_models.IPlayer)
		if isPlayer && !enemyObject.Destroyed() {
			player.GetDamage(player.Health())
			enemyObject.Destroy()
		}
	}
}
