package game_models

type IEnemy interface {
	IGameObjectWithHP
	IAutoMoveable
}

type IEnemyShooter interface {
	IGameObjectWithHP
	IAutoMoveable
	IShooter
}
