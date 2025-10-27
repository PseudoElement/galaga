package game_models

type PlayerTier = int8
type IPlayer interface {
	IGameObjectWithHP
	IMoveable
	IShooter
	Tier() PlayerTier
}

type PlayerFactory func(coords Coords, injector any) IPlayer
