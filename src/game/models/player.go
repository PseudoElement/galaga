package game_models

type IPlayer interface {
	IGameObjectWithHP
	IMoveable
	IShooter
	TakeBoost(boostItem BoostItem)
}
