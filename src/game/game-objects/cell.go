package game_objects

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	game_models "github.com/pseudoelement/galaga/src/game/models"
)

type Cell struct {
	color  string
	style  lipgloss.Style
	coords game_models.Coords
}

func NewCell(params game_models.CellConstructorParams) *Cell {
	style := lipgloss.NewStyle().Height(1).Width(1).Background(lipgloss.Color(params.Color))

	return &Cell{
		color:  params.Color,
		coords: params.Coords,
		style:  style,
	}
}

func (c *Cell) Style() lipgloss.Style {
	return c.style
}

func (c *Cell) Color() string {
	r, g, b, _ := c.style.GetBackground().RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func (c *Cell) SetColor(color string) {
	c.style = c.style.Background(lipgloss.Color(color))
}

func (c *Cell) Coords() game_models.Coords {
	return c.coords
}

func (c *Cell) SetCoords(coords game_models.Coords) {
	c.coords = coords
}

var _ game_models.ICell = (*Cell)(nil)
