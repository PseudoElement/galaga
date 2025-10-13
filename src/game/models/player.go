package game_models

type IPlayer interface {
	IGameObject
	Shot()
	TakeBoost(boostItem BoostItem)
	Health() int16
	GetHeal(plusHealthAmount int16)
	GetDamage(minusHealthAmount int16)
}
