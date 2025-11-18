package enemy

import (
	g_c "github.com/pseudoelement/galaga/src/game/game-constants"
	game_objects "github.com/pseudoelement/galaga/src/game/game-objects"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type TNTEnemy struct {
	*Enemy
	injector models.IAppInjector
}

func NewTNTEnemy(x, y int16, health int16, injector models.IAppInjector) g_m.IEnemy {
	cells := []g_m.ICell{
		//1st
		g_m.NewCell(g_m.CellParams(x, y, "#ff0101ff", ""), g_c.TNT),
		g_m.NewCell(g_m.CellParams(x+1, y, "#ff0101ff", ""), g_c.TNT),
		g_m.NewCell(g_m.CellParams(x+2, y, "#ff0101ff", ""), g_c.TNT),
		//2nd
		g_m.NewCell(g_m.CellParams(x, y+1, "#ff0101ff", "T"), g_c.TNT),
		g_m.NewCell(g_m.CellParams(x+1, y+1, "#ff0101ff", "N"), g_c.TNT),
		g_m.NewCell(g_m.CellParams(x+2, y+1, "#ff0101ff", "T"), g_c.TNT),
		//3th
		g_m.NewCell(g_m.CellParams(x, y+2, "#ff0101ff", ""), g_c.TNT),
		g_m.NewCell(g_m.CellParams(x+1, y+2, "#ff0101ff", ""), g_c.TNT),
		g_m.NewCell(g_m.CellParams(x+2, y+2, "#ff0101ff", ""), g_c.TNT),
	}

	arenaWidth, _ := injector.GameSrv().ArenaSize()
	// full row of move right + 1 step down + full row of move left + 1 step down
	movementPattern := make([]g_m.MoveDir, 0, arenaWidth*2+2)

	for range arenaWidth {
		movementPattern = append(movementPattern, g_m.MoveRightX1_Y0())
	}
	movementPattern = append(movementPattern, g_m.MoveBottomX0_Y3())
	for range arenaWidth {
		movementPattern = append(movementPattern, g_m.MoveLeftX1_Y0())
	}
	movementPattern = append(movementPattern, g_m.MoveBottomX0_Y3())

	return &TNTEnemy{
		Enemy:    NewEnemy(health, cells, movementPattern, g_c.TNT),
		injector: injector,
	}
}

func (e *TNTEnemy) Name() string {
	return g_c.TNT
}

func (e *TNTEnemy) Price() int32 {
	return 3
}

func (e *TNTEnemy) MovementDelay(tickMs int) int {
	return tickMs * 3
}

func (e *TNTEnemy) IsBoss() bool {
	return false
}

func (e *TNTEnemy) Blast() []g_m.IGameObject {
	// !!! supposed that [0]  element in cells list is top left corner of object
	tntLeftCornerCoordX := e.Cells()[0].Coords().X
	tntLeftCornerCoordY := e.Cells()[0].Coords().Y
	blastLeftCornerCoordX := tntLeftCornerCoordX - 2
	blastLeftCornerCoordY := tntLeftCornerCoordY - 2

	blast := game_objects.NewTNTBlast(g_m.Coords{
		X: blastLeftCornerCoordX,
		Y: blastLeftCornerCoordY,
	},
		e.injector,
	)

	return []g_m.IGameObject{blast}
}

var _ g_m.IEnemyBomb = (*TNTEnemy)(nil)
