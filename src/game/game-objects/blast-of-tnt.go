package game_objects

import (
	"time"

	game_constants "github.com/pseudoelement/galaga/src/game/game-constants"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type TNTBlast struct {
	*g_m.GameObject
	destructTime int
}

func NewTNTBlast(coords g_m.Coords, injector models.IAppInjector) *TNTBlast {
	cells := []g_m.ICell{
		//1st
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+1, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+1, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+1, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+1, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+1, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+1, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+1, "#fff82eff", "")),
		//2nd
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+2, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+2, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+2, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+2, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+2, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+2, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+2, "#fff82eff", "")),
		//3th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+3, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+3, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+3, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+3, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+3, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+3, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+3, "#fff82eff", "")),
		//4th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+4, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+4, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+4, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+4, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+4, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+4, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+4, "#fff82eff", "")),
		//5th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+5, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+5, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+5, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+5, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+5, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+5, "#fff82eff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+5, "#fff82eff", "")),
	}

	msNow := time.Now().UnixMilli()
	destructTime := int(msNow + 150)

	return &TNTBlast{
		GameObject:   g_m.NewGameObject(cells),
		destructTime: destructTime,
	}
}

func (b *TNTBlast) Name() string {
	return game_constants.BLAST_OF_TNT
}

func (c *TNTBlast) DestructInMs() int {
	return c.destructTime
}

var _ g_m.IBlast = (*TNTBlast)(nil)
