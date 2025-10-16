package enemy

import (
	g_o "github.com/pseudoelement/galaga/src/game/game-objects"
	g_m "github.com/pseudoelement/galaga/src/game/models"
)

type OctopusEnemy struct {
	*Enemy
}

func NewOctopusEnemy(x, y int16) g_m.IEnemyShooter {
	cells := []g_m.ICell{
		//1st
		g_o.NewCell(g_m.CellParams(x, y, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+1, y, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+2, y, "#8105bbff")),
		g_o.NewCell(g_m.CellParams(x+3, y, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+4, y, "#ffffffff")),
		//2nd
		g_o.NewCell(g_m.CellParams(x, y+1, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+1, y+1, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+2, y+1, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+3, y+1, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+4, y+1, "#ffffffff")),
		//3th
		g_o.NewCell(g_m.CellParams(x, y+2, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+1, y+2, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+2, y+2, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+3, y+2, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+4, y+2, "#ffffffff")),
		//4th
		g_o.NewCell(g_m.CellParams(x, y+3, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+1, y+3, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+2, y+3, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+3, y+3, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+4, y+3, "#ffffffff")),
		//5th
		g_o.NewCell(g_m.CellParams(x, y+4, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+2, y+4, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+4, y+4, "#ffffffff")),
		//6th
		g_o.NewCell(g_m.CellParams(x, y+5, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+2, y+5, "#ffffffff")),
		g_o.NewCell(g_m.CellParams(x+4, y+5, "#ffffffff")),
	}

	return &OctopusEnemy{
		Enemy: NewEnemy(10, cells),
	}
}

func (dp *OctopusEnemy) MovementDelay(tickMs int) int {
	return tickMs * 8
}

func (dp *OctopusEnemy) Shot() []g_m.IBullet {
	return []g_m.IBullet{}
}

var _ g_m.IEnemy = (*OctopusEnemy)(nil)
var _ g_m.IEnemyShooter = (*OctopusEnemy)(nil)
