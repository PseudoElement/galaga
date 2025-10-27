package game_objects

import (
	g_c "github.com/pseudoelement/galaga/src/game/game-constants"
	game_constants "github.com/pseudoelement/galaga/src/game/game-constants"
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
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y, "#2db9f0ff", "")),
		//2nd
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+1, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+1, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+1, "#2db9f0ff", "")),
		//3th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+2, "#2db9f0ff", "2")),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+2, "#2db9f0ff", "x")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+2, "#2db9f0ff", "G")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+2, "#2db9f0ff", "U")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+2, "#2db9f0ff", "N")),
		//4th
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+3, "#2db9f0ff", "")),
		//5th
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+4, "#2db9f0ff", "")),
	}

	return &DoubleGunShipBoost{
		GameObject: g_m.NewGameObject(cells),
		injector:   injector,
	}
}

func (b *DoubleGunShipBoost) Name() string {
	return g_c.BOOST_DOUBLE_GUN_SHIP
}

func (b *DoubleGunShipBoost) Boost(p g_m.IPlayer) {
	b.injector.GameSrv().SetPlayer(
		b.injector.Factories().PlayerFactory(
			p.Cells()[0].Coords(),
			game_constants.PLAYER_DOBLE_GUN,
		),
	)
}

func (e *DoubleGunShipBoost) MovementDelay(tickMs int) int {
	return tickMs * 12
}

var _ g_m.IBoost = (*DoubleGunShipBoost)(nil)
