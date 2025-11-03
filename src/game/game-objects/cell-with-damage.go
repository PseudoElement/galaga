package game_objects

import (
	"time"

	g_m "github.com/pseudoelement/galaga/src/game/models"
)

type CellWithDamage struct {
	*g_m.Cell
	damageCount  int16
	destructTime int
}

func NewCellWithDamage(params g_m.CellWithDamageConstructorParams) *CellWithDamage {
	msNow := time.Now().UnixMilli()
	destructTime := int(msNow + 150)

	return &CellWithDamage{
		Cell:         g_m.NewCell(params.CellConstructorParams),
		damageCount:  params.DamageCount,
		destructTime: destructTime,
	}
}

func (c *CellWithDamage) CanDamage() bool {
	return c.damageCount > 0
}

func (c *CellWithDamage) Damage(player g_m.IPlayer) {
	player.GetDamage(c.damageCount)
}

func (c *CellWithDamage) SetDamageCount(damageCount int16) {
	c.damageCount = damageCount
}

func (c *CellWithDamage) DestructInMs() int {
	return c.destructTime
}

var _ g_m.ICellWithDamage = (*CellWithDamage)(nil)
var _ g_m.IAutoDestructable = (*CellWithDamage)(nil)
