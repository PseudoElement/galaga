package game_objects

import (
	"time"

	g_c "github.com/pseudoelement/galaga/src/game/game-constants"
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
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+1, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+1, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+1, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+1, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+1, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+1, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+1, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		//2nd
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+2, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+2, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+2, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+2, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+2, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+2, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+2, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		//3th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+3, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+3, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+3, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+3, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+3, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+3, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+3, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		//4th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+4, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+4, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+4, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+4, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+4, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+4, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+4, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		//5th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+5, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+5, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+5, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+5, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+5, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+5, "#fff82eff", ""), g_c.BLAST_OF_TNT),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+5, "#fff82eff", ""), g_c.BLAST_OF_TNT),
	}

	msNow := time.Now().UnixMilli()
	destructTime := int(msNow + 150)

	return &TNTBlast{
		GameObject:   g_m.NewGameObject(cells, g_c.BLAST_OF_TNT),
		destructTime: destructTime,
	}
}

func (b *TNTBlast) Name() string {
	return g_c.BLAST_OF_TNT
}

func (c *TNTBlast) DestructInMs() int {
	return c.destructTime
}

var _ g_m.IBlast = (*TNTBlast)(nil)
