package game_objects

import (
	g_c "github.com/pseudoelement/galaga/src/game/game-constants"
	g_m "github.com/pseudoelement/galaga/src/game/models"
)

type Bullet struct {
	*g_m.GameObject

	damage    int16
	ownerName string
}

func NewBullet(coords g_m.Coords, color string, ownerName string) *Bullet {
	cells := []g_m.ICell{g_m.NewCell(g_m.CellParams(coords.X, coords.Y, color, ""), g_c.BULLET)}

	return &Bullet{
		GameObject: g_m.NewGameObject(cells, g_c.BULLET),
		damage:     1,
		ownerName:  ownerName,
	}
}

func (b *Bullet) Name() string {
	return g_c.BULLET
}

func (b *Bullet) Owner() string {
	return b.ownerName
}

func (b *Bullet) Damage(object g_m.IDamageable) {
	object.GetDamage(b.damage)
}

func (b *Bullet) MovementDelay(tickMs int) int {
	// 1 Move() call per 5 ticks
	return tickMs * 3
}

var _ g_m.IBullet = (*Bullet)(nil)
