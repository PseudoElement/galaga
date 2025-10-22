package game_objects

import g_m "github.com/pseudoelement/galaga/src/game/models"

type CellWithDamage struct {
	*g_m.Cell
	damageCount int16
}

func NewCellWithDamage(params g_m.CellConstructorParams) *CellWithDamage {
	return &CellWithDamage{
		Cell: g_m.NewCell(params),
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

var _ g_m.ICellWithDamage = (*CellWithDamage)(nil)
