package enemy

import (
	g_c "github.com/pseudoelement/galaga/src/game/game-constants"
	g_o "github.com/pseudoelement/galaga/src/game/game-objects"
	g_m "github.com/pseudoelement/galaga/src/game/models"
)

type SmallSpaceShipEnemy struct {
	*Enemy
}

func NewSmallSpaceShipEnemy(x, y int16, health int16) g_m.IEnemyShooter {
	cells := []g_m.ICell{
		//1st
		g_m.NewCell(g_m.CellParams(x, y, "#a60bc6ff", "")),
		g_m.NewCell(g_m.CellParams(x+1, y, "#a60bc6ff", "")),
		g_m.NewCell(g_m.CellParams(x+2, y, "#a60bc6ff", "")),
		//2nd
		g_m.NewCell(g_m.CellParams(x+1, y+1, "#a60bc6ff", "")),
		//3th
		g_m.NewCell(g_m.CellParams(x+1, y+2, "#a60bc6ff", "")),
	}
	movementPattern := []g_m.MoveDir{
		g_m.MoveLeftBottomX1_Y1(),
		g_m.MoveLeftBottomX1_Y1(),
		g_m.MoveLeftBottomX1_Y1(),
		g_m.MoveLeftBottomX1_Y1(),
		g_m.MoveLeftBottomX1_Y1(),
		g_m.MoveLeftBottomX1_Y1(),
		g_m.MoveRightBottomX1_Y1(),
		g_m.MoveRightBottomX1_Y1(),
		g_m.MoveRightBottomX1_Y1(),
		g_m.MoveRightBottomX1_Y1(),
		g_m.MoveRightBottomX1_Y1(),
		g_m.MoveRightBottomX1_Y1(),
	}

	return &SmallSpaceShipEnemy{
		Enemy: NewEnemy(health, cells, movementPattern),
	}
}

func (e *SmallSpaceShipEnemy) Name() string {
	return g_c.SMALL_SPACE_SHIP
}

func (e *SmallSpaceShipEnemy) MovementDelay(tickMs int) int {
	return tickMs * 12
}

func (e *SmallSpaceShipEnemy) Shot() []g_m.IBullet {
	bottomMidCell := e.Cells()[len(e.Cells())-1]
	bullet := g_o.NewEnemyBullet(bottomMidCell.Coords(), "#d48003ff")

	return []g_m.IBullet{bullet}
}

var _ g_m.IEnemyShooter = (*SmallSpaceShipEnemy)(nil)
