package enemy

import (
	g_c "github.com/pseudoelement/galaga/src/game/game-constants"
	g_m "github.com/pseudoelement/galaga/src/game/models"
)

type OctopusEnemy struct {
	*Enemy
}

func NewOctopusEnemy(x, y int16, health int16) g_m.IEnemy {
	cells := []g_m.ICell{
		//1st
		g_m.NewCell(g_m.CellParams(x, y, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+1, y, "#ff0101ff", "")),
		g_m.NewCell(g_m.CellParams(x+2, y, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+3, y, "#ff0101ff", "")),
		g_m.NewCell(g_m.CellParams(x+4, y, "#ffffffff", "")),
		//2nd
		g_m.NewCell(g_m.CellParams(x, y+1, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+1, y+1, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+2, y+1, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+3, y+1, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+4, y+1, "#ffffffff", "")),
		//3th
		g_m.NewCell(g_m.CellParams(x, y+2, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+1, y+2, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+2, y+2, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+3, y+2, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+4, y+2, "#ffffffff", "")),
		//4th
		g_m.NewCell(g_m.CellParams(x, y+3, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+1, y+3, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+2, y+3, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+3, y+3, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+4, y+3, "#ffffffff", "")),
		//5th
		g_m.NewCell(g_m.CellParams(x, y+4, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+2, y+4, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+4, y+4, "#ffffffff", "")),
		//6th
		g_m.NewCell(g_m.CellParams(x, y+5, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+2, y+5, "#ffffffff", "")),
		g_m.NewCell(g_m.CellParams(x+4, y+5, "#ffffffff", "")),
	}
	movementPattern := []g_m.MoveDir{
		g_m.MoveLeftBottomX1_Y1(),
		g_m.MoveRightBottomX1_Y1(),
	}

	return &OctopusEnemy{
		Enemy: NewEnemy(health, cells, movementPattern),
	}
}

func (e *OctopusEnemy) Name() string {
	return g_c.OCTOPUS
}

func (e *OctopusEnemy) Price() int32 {
	return 1
}

func (e *OctopusEnemy) MovementDelay(tickMs int) int {
	return tickMs * 8
}

var _ g_m.IEnemy = (*OctopusEnemy)(nil)
