package game_objects

import game_models "github.com/pseudoelement/galaga/src/game/models"

type Cell struct {
	color  string
	coords game_models.Coords
}

func NewCell(params game_models.CellConstructorParams) *Cell {
	return &Cell{
		color:  params.Color,
		coords: params.Coords,
	}
}

func (c *Cell) Color() string {
	return c.color
}

func (c *Cell) SetColor(color string) {
	c.color = color
}

func (c *Cell) Coords() game_models.Coords {
	return c.coords
}

func (c *Cell) SetCoords(coords game_models.Coords) {
	c.coords = coords
}

var _ game_models.ICell = (*Cell)(nil)
