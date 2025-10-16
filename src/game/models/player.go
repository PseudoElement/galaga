package game_models

type IPlayer interface {
	IGameObjectWithHP
	IMoveable
	Shot() []IBullet
	TakeBoost(boostItem BoostItem)
}
