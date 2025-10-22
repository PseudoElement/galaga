package game_models

type IBoost interface {
	IGameObject
	IAutoMoveable
	Boost(player IPlayer)
}
