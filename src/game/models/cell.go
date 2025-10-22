package game_models

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type Cell struct {
	color     string
	style     lipgloss.Style
	coords    Coords
	destroyed bool
}

func NewCell(params CellConstructorParams) *Cell {
	style := lipgloss.NewStyle().
		Height(1).
		Width(1).
		Background(lipgloss.Color(params.Color)).
		SetString(params.Text)

	return &Cell{
		color:     params.Color,
		coords:    params.Coords,
		style:     style,
		destroyed: false,
	}
}

func (c *Cell) Style() lipgloss.Style {
	return c.style
}

func (c *Cell) SetStyle(newStyle lipgloss.Style) {
	c.style = newStyle
}

func (c *Cell) Color() string {
	r, g, b, _ := c.style.GetBackground().RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func (c *Cell) SetColor(color string) {
	c.style = c.style.Background(lipgloss.Color(color))
}

func (c *Cell) Coords() Coords {
	return c.coords
}

func (c *Cell) SetCoords(coords Coords) {
	c.coords = coords
}

func (obj *Cell) Destroy() {
	obj.destroyed = true
}

func (obj *Cell) Destroyed() bool {
	return obj.destroyed
}

var _ ICell = (*Cell)(nil)
