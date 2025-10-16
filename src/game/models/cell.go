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
}

type CellConstructorParams struct {
	/* hex */
	Color  string
	Coords Coords
	// CanDamage bool
}

func CellParams(x, y int16, color string) CellConstructorParams {
	return CellConstructorParams{
		Color:  color,
		Coords: Coords{X: x, Y: y},
	}
}

type ICellWithDamage interface {
	ICell
	CanDamage() bool
	Damage(player IPlayer)
	SetDamageCount(damageCount int16)
}
