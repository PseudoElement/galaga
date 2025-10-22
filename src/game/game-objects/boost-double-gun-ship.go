package game_objects

import (
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type DoubleGunShipBoost struct {
	*g_m.GameObject

	injector models.IAppInjector
}

func NewDoubleGunShipBoost(coords g_m.Coords, injector models.IAppInjector) *DoubleGunShipBoost {
	cells := []g_m.ICell{
		//1st
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y, "#37c60bff", "")),
		//2nd
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+1, "#37c60bff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+1, "#37c60bff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+1, "#37c60bff", "")),
		//3th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+2, "#37c60bff", "B")),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+2, "#37c60bff", "O")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+2, "#37c60bff", "O")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+2, "#37c60bff", "S")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+2, "#37c60bff", "T")),
		//4th
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+3, "#37c60bff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+3, "#37c60bff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+3, "#37c60bff", "")),
		//5th
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+4, "#37c60bff", "")),
	}

	return &DoubleGunShipBoost{
		GameObject: g_m.NewGameObject(cells),
		injector:   injector,
	}
}

func (b *DoubleGunShipBoost) Boost(p g_m.IPlayer) {
	p = b.injector.Factories().PlayerFactory(p.Cells()[0].Coords(), models.DOUBLE_GUN_PLAYER)
}

func (e *DoubleGunShipBoost) MovementDelay(tickMs int) int {
	return tickMs * 12
}

var _ g_m.IBoost = (*DoubleGunShipBoost)(nil)
