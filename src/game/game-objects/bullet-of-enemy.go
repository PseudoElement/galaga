package game_objects

import g_m "github.com/pseudoelement/galaga/src/game/models"

type EnemyBullet struct {
	*Bullet
}

func NewEnemyBullet(coords g_m.Coords, color string) *EnemyBullet {
	return &EnemyBullet{
		Bullet: NewBullet(coords, color),
	}
}

func (b *EnemyBullet) Owner() string {
	return "Enemy"
}

var _ g_m.IBullet = (*EnemyBullet)(nil)
