package game_models

type IEnemyFactory interface {
	SpawnEnemy() IEnemy
}
