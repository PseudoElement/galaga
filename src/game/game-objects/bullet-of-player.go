package game_objects

import g_m "github.com/pseudoelement/galaga/src/game/models"

type PlayerBullet struct {
	*Bullet
}

func NewPlayerBullet(coords g_m.Coords, color string) *PlayerBullet {
	return &PlayerBullet{
		Bullet: NewBullet(coords, color),
	}
}

func (b *PlayerBullet) Owner() string {
	return "Player"
}

var _ g_m.IBullet = (*PlayerBullet)(nil)
