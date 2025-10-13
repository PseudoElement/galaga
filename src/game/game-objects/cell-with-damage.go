package game_objects

import game_models "github.com/pseudoelement/galaga/src/game/models"

type CellWithDamage struct {
	*Cell
	damageCount int16
}

func NewCellWithDamage(params game_models.CellConstructorParams) *CellWithDamage {
	return &CellWithDamage{
		Cell: NewCell(params),
	}
}

func (c *CellWithDamage) CanDamage() bool {
	return c.damageCount > 0
}

func (c *CellWithDamage) Damage(player game_models.IPlayer) {
	player.GetDamage(c.damageCount)
}

func (c *CellWithDamage) SetDamageCount(damageCount int16) {
	c.damageCount = damageCount
}

var _ game_models.ICellWithDamage = (*CellWithDamage)(nil)
