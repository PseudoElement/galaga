package game_models

type Bullet struct {
	*GameObject

	damage int16
}

func NewBullet(coords Coords, color string) *Bullet {
	cells := []ICell{NewCell(CellParams(coords.X, coords.Y, color, ""))}

	return &Bullet{
		GameObject: NewGameObject(cells),
		damage:     1,
	}
}

func (b *Bullet) Damage(object IGameObjectWithHP) {
	object.GetDamage(b.damage)
}

func (b *Bullet) MovementDelay(tickMs int) int {
	// 1 Move() call per 5 ticks
	return tickMs * 3
}
