package game_objects

import (
	g_m "github.com/pseudoelement/galaga/src/game/models"
)

type Bullet struct {
	*GameObject

	damage int16
}

func NewBullet(coords g_m.Coords, color string) *Bullet {
	cells := []g_m.ICell{NewCell(g_m.CellParams(coords.X, coords.Y, color))}

	return &Bullet{
		GameObject: NewGameObject(cells),
		damage:     1,
	}
}

func (b *Bullet) Damage(object g_m.IGameObjectWithHP) {
	object.GetDamage(b.damage)
}

func (b *Bullet) MovementDelay(tickMs int) int {
	// 1 Move() call per 5 ticks
	return tickMs * 3
}
