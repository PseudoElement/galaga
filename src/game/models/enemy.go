package game_models

type IEnemy interface {
	IGameObjectWithHP
	IAutoMoveable
	Name() string
	MoveDir(arenaWidth, arenaHeight int) MoveDir
}

type IEnemyShooter interface {
	IEnemy
	IShooter
}
