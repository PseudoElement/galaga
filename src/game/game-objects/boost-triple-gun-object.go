package game_objects

import (
	g_c "github.com/pseudoelement/galaga/src/game/game-constants"
	game_constants "github.com/pseudoelement/galaga/src/game/game-constants"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type TripleGunShipBoost struct {
	*g_m.GameObject

	injector models.IAppInjector
}

func NewTripleGunShipBoost(coords g_m.Coords, injector models.IAppInjector) *TripleGunShipBoost {
	cells := []g_m.ICell{
		//1st
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y, "#ff0505ff", "")),
		//2nd
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+1, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+1, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+1, "#ff0505ff", "")),
		//3th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+2, "#ff0505ff", "3")),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+2, "#ff0505ff", "x")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+2, "#ff0505ff", "G")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+2, "#ff0505ff", "U")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+2, "#ff0505ff", "N")),
		//4th
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+3, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+3, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+3, "#ff0505ff", "")),
		//5th
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+4, "#ff0505ff", "")),
	}

	return &TripleGunShipBoost{
		GameObject: g_m.NewGameObject(cells),
		injector:   injector,
	}
}

func (b *TripleGunShipBoost) Name() string {
	return g_c.BOOST_TRIPLE_GUN_SHIP
}

func (b *TripleGunShipBoost) Boost(p g_m.IPlayer) {
	b.injector.GameSrv().SetPlayer(
		b.injector.Factories().PlayerFactory(
			p.Cells()[0].Coords(),
			game_constants.PLAYER_TRIPLE_GUN,
		),
	)
}

func (e *TripleGunShipBoost) MovementDelay(tickMs int) int {
	return tickMs * 12
}

var _ g_m.IBoost = (*TripleGunShipBoost)(nil)
