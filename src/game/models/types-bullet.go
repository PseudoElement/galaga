package game_models

type IBullet interface {
	IGameObject
	IAutoMoveable
	IOwnerable
	Damage(damageable IDamageable)
}
