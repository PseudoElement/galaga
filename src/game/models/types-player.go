package game_models

type IPlayer interface {
	IGameObjectWithHP
	IMoveable
	IShooter
	TakeBoost(boostItem IBoost)
}

type PlayerFactory func(coords Coords, injector any) IPlayer
