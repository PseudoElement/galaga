package enemy

import (
	g_o "github.com/pseudoelement/galaga/src/game/game-objects"
	g_m "github.com/pseudoelement/galaga/src/game/models"
)

type Enemy struct {
	*g_o.GameObject

	health int16
}

func NewEnemy(health int16, cells []g_m.ICell) *Enemy {
	return &Enemy{health: health, GameObject: g_o.NewGameObject(cells)}
}

func (e *Enemy) Health() int16 {
	return e.health
}

func (e *Enemy) GetHeal(amount int16) {
	if amount < 0 {
		panic("[Enemy_GetHeal] amount should be positive.")
	}
	e.health += amount
}

func (e *Enemy) GetDamage(amount int16) {
	if amount < 0 {
		panic("[Enemy_GetDamage] amount should be positive.")
	}
	e.health -= amount

	if e.health <= 0 {
		e.Destroy()
	}
}

var _ g_m.IHealable = (*Enemy)(nil)
var _ g_m.IDamageable = (*Enemy)(nil)
