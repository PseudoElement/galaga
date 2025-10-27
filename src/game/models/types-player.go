package game_models

type IPlayer interface {
	IGameObjectWithHP
	IMoveable
	IShooter
}

type PlayerFactory func(coords Coords, injector any) IPlayer
