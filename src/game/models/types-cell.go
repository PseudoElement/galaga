package game_models

import (
	"github.com/charmbracelet/lipgloss"
)

type Coords struct {
	X int16
	Y int16
}

type ICell interface {
	IDestroyable
	Color() string
	SetColor(color string)
	Coords() Coords
	SetCoords(coords Coords)
	Style() lipgloss.Style
	SetStyle(newStyle lipgloss.Style)
}

type CellConstructorParams struct {
	/* hex */
	Color  string
	Coords Coords
	Text   string
	// CanDamage bool
}

func CellParams(x, y int16, color string, text string) CellConstructorParams {
	return CellConstructorParams{
		Color:  color,
		Coords: Coords{X: x, Y: y},
		Text:   text,
	}
}

type ICellWithDamage interface {
	ICell
	CanDamage() bool
	Damage(player IPlayer)
	SetDamageCount(damageCount int16)
}
