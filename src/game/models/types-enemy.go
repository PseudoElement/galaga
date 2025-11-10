package game_models

type IEnemy interface {
	IGameObjectWithHP
	IAutoMoveable
	IWithPrice
	Name() string
	MoveDir(arenaWidth, arenaHeight int) MoveDir
	ChangeMovementPattern(newPattern []MoveDir)
	IsBoss() bool
}

type IEnemyShooter interface {
	IEnemy
	IAutoShooter
}

type IEnemyBomb interface {
	IEnemy
	Blast() []IGameObject
}
