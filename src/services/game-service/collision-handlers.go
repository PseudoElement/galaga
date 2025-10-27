package game_srv

import game_models "github.com/pseudoelement/galaga/src/game/models"

func handleCollisionScenarios(obj1 game_models.IGameObject, obj2 game_models.IGameObject) {
	_checkCollisionPlayerWithEnemy(obj1, obj2)
	_checkCollisionPlayerWithEnemy(obj2, obj1)

	_checkCollisionWithBullet(obj1, obj2)
	_checkCollisionWithBullet(obj2, obj1)

	_checkCollisionWithBoost(obj1, obj2)
	_checkCollisionWithBoost(obj2, obj1)
}

func _checkCollisionWithBullet(obj game_models.IGameObject, otherObject game_models.IGameObject) {
	bulletObject, isObjectBullet := obj.(game_models.IBullet)
	if isObjectBullet {
		damageableObject, isDamageable := otherObject.(game_models.IDamageable)
		// @CHECK if bullet of old player can damage new player (taken boost)
		doesBulletCrossedItsOwner := bulletObject.Owner() == otherObject
		if isDamageable && !doesBulletCrossedItsOwner && !bulletObject.Destroyed() {
			bulletObject.Damage(damageableObject)
			bulletObject.Destroy()
		}
	}
}

func _checkCollisionWithBoost(obj game_models.IGameObject, otherObject game_models.IGameObject) {
	boostObject, isObjectBoost := obj.(game_models.IBoost)
	if isObjectBoost {
		player, isPlayer := otherObject.(game_models.IPlayer)
		if isPlayer && !boostObject.Destroyed() {
			boostObject.Boost(player)
			boostObject.Destroy()
		}
	}
}

func _checkCollisionPlayerWithEnemy(obj game_models.IGameObject, otherObject game_models.IGameObject) {
	enemyObject, isObjectEnemy := obj.(game_models.IEnemy)
	if isObjectEnemy {
		player, isPlayer := otherObject.(game_models.IPlayer)
		if isPlayer && !enemyObject.Destroyed() {
			player.GetDamage(player.Health())
			enemyObject.Destroy()
		}
	}
}
