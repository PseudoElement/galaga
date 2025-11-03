package game_models

type IEnemy interface {
	IGameObjectWithHP
	IAutoMoveable
	IWithPrice
	Name() string
	MoveDir(arenaWidth, arenaHeight int) MoveDir
}

type IEnemyShooter interface {
	IEnemy
	IShooter
}

type IEnemyBomb interface {
	IEnemy
	Blast() []IGameObject
}
