package game_models

type IBullet interface {
	IGameObject
	IAutoMoveable
	Damage(player IGameObjectWithHP)
	Owner() string
}
