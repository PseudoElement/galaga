package game_models

type IBossEnemy interface {
	IEnemyShooter
}

type IDirectionChanger interface {
	UpdateMovementPattern(player IPlayer)
}
