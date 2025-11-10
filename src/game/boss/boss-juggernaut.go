package boss

import (
	"math"

	"github.com/pseudoelement/galaga/src/game/enemy"
	game_constants "github.com/pseudoelement/galaga/src/game/game-constants"
	g_o "github.com/pseudoelement/galaga/src/game/game-objects"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type JuggernautBoss struct {
	*enemy.Enemy
}

func NewJuggernautBoss(x, y int16, health int16, injector models.IAppInjector) g_m.IBossEnemy {
	cells := []g_m.ICell{
		// 1st
		g_m.NewCell(g_m.CellParams(x, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+1, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+2, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+3, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+4, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+5, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+6, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+7, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+8, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+9, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+10, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+11, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+12, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+13, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+14, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+15, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+16, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+17, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+18, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+19, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+20, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+21, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+22, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+23, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+24, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+25, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+26, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+27, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+28, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+29, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+30, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+31, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+32, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+33, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+34, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+35, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+36, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+37, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+38, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+39, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+40, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+41, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+42, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+43, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+44, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+45, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+46, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+47, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+48, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+49, y, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+50, y, "#950f03ff", "")),
		//2nd
		g_m.NewCell(g_m.CellParams(x+3, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+4, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+5, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+6, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+7, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+8, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+9, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+10, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+11, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+12, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+13, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+14, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+15, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+16, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+17, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+18, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+19, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+20, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+21, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+22, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+23, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+24, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+25, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+26, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+27, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+28, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+29, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+30, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+31, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+32, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+33, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+34, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+35, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+36, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+37, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+38, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+39, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+40, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+41, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+42, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+43, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+44, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+45, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+46, y+1, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+47, y+1, "#950f03ff", "")),
		//3th
		g_m.NewCell(g_m.CellParams(x+3, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+7, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+8, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+9, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+10, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+11, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+12, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+13, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+14, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+15, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+16, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+17, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+18, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+19, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+20, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+21, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+22, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+23, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+24, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+25, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+26, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+27, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+28, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+29, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+30, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+31, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+32, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+33, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+34, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+35, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+36, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+37, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+38, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+39, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+40, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+41, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+42, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+43, y+2, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+47, y+2, "#950f03ff", "")),
		//4th
		g_m.NewCell(g_m.CellParams(x+3, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+7, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+8, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+9, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+10, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+11, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+12, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+13, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+14, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+15, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+16, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+17, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+18, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+19, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+20, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+21, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+22, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+23, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+24, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+25, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+26, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+27, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+28, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+29, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+30, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+31, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+32, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+33, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+34, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+35, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+36, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+37, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+38, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+39, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+40, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+41, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+42, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+43, y+3, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+47, y+3, "#950f03ff", "")),
		//5th
		g_m.NewCell(g_m.CellParams(x+3, y+4, "#950f03ff", "")), // left small gun end
		g_m.NewCell(g_m.CellParams(x+10, y+4, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+11, y+4, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+12, y+4, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+18, y+4, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+19, y+4, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+20, y+4, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+30, y+4, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+31, y+4, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+32, y+4, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+38, y+4, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+39, y+4, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+40, y+4, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+47, y+4, "#950f03ff", "")), // left small gun end
		//6th
		g_m.NewCell(g_m.CellParams(x+10, y+5, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+11, y+5, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+12, y+5, "#950f03ff", "")),

		g_m.NewCell(g_m.CellParams(x+18, y+5, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+19, y+5, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+20, y+5, "#950f03ff", "")),

		g_m.NewCell(g_m.CellParams(x+30, y+5, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+31, y+5, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+32, y+5, "#950f03ff", "")),

		g_m.NewCell(g_m.CellParams(x+38, y+5, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+39, y+5, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+40, y+5, "#950f03ff", "")),
		//7th
		g_m.NewCell(g_m.CellParams(x+10, y+6, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+11, y+6, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+12, y+6, "#950f03ff", "")),

		g_m.NewCell(g_m.CellParams(x+18, y+6, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+19, y+6, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+20, y+6, "#950f03ff", "")),

		g_m.NewCell(g_m.CellParams(x+30, y+6, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+31, y+6, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+32, y+6, "#950f03ff", "")),

		g_m.NewCell(g_m.CellParams(x+38, y+6, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+39, y+6, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+40, y+6, "#950f03ff", "")),
		//8th
		g_m.NewCell(g_m.CellParams(x+10, y+7, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+11, y+7, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+12, y+7, "#950f03ff", "")),

		g_m.NewCell(g_m.CellParams(x+18, y+7, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+19, y+7, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+20, y+7, "#950f03ff", "")),

		g_m.NewCell(g_m.CellParams(x+30, y+7, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+31, y+7, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+32, y+7, "#950f03ff", "")),

		g_m.NewCell(g_m.CellParams(x+38, y+7, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+39, y+7, "#950f03ff", "")),
		g_m.NewCell(g_m.CellParams(x+40, y+7, "#950f03ff", "")),
		//9th
		g_m.NewCell(g_m.CellParams(x+10, y+8, "#950f03ff", "")), // big gun end
		g_m.NewCell(g_m.CellParams(x+11, y+8, "#950f03ff", "")), // big gun end
		g_m.NewCell(g_m.CellParams(x+12, y+8, "#950f03ff", "")), // big gun end

		g_m.NewCell(g_m.CellParams(x+18, y+8, "#950f03ff", "")), // big gun end
		g_m.NewCell(g_m.CellParams(x+19, y+8, "#950f03ff", "")), // big gun end
		g_m.NewCell(g_m.CellParams(x+20, y+8, "#950f03ff", "")), // big gun end

		g_m.NewCell(g_m.CellParams(x+30, y+8, "#950f03ff", "")), // big gun end
		g_m.NewCell(g_m.CellParams(x+31, y+8, "#950f03ff", "")), // big gun end
		g_m.NewCell(g_m.CellParams(x+32, y+8, "#950f03ff", "")), // big gun end

		g_m.NewCell(g_m.CellParams(x+38, y+8, "#950f03ff", "")), // big gun end
		g_m.NewCell(g_m.CellParams(x+39, y+8, "#950f03ff", "")), // big gun end
		g_m.NewCell(g_m.CellParams(x+40, y+8, "#950f03ff", "")), // big gun end
	}

	arenaWidth, _ := injector.GameSrv().ArenaSize()
	// full row of move right + full row of move left
	movementPattern := make([]g_m.MoveDir, 0, arenaWidth+2)
	halfOfWidth := int(math.Ceil(float64(arenaWidth) / 2))

	for range halfOfWidth {
		movementPattern = append(movementPattern, g_m.MoveRightX2_Y0())
	}
	for range halfOfWidth {
		movementPattern = append(movementPattern, g_m.MoveLeftX2_Y0())
	}

	return &JuggernautBoss{
		Enemy: enemy.NewEnemy(health, cells, movementPattern),
	}
}

func (b *JuggernautBoss) Name() string {
	return game_constants.BOSS_JUGGERNAUT
}

func (b *JuggernautBoss) Price() int32 {
	return 200
}

func (b *JuggernautBoss) MovementDelay(tickMs int) int {
	return tickMs * 5
}

func (b *JuggernautBoss) IsBoss() bool {
	return true
}

func (e *JuggernautBoss) ShootingDelay(tickMs int) int {
	return tickMs * 200
}

func (b *JuggernautBoss) Shot() []g_m.IBullet {
	/*
	 * 9 bullets per 1 big bullet
	 * 2 bullets of small guns
	 * 2 small guns and 4 big guns
	 * 4 * 9 + 2 = 38
	 */
	bullets := make([]g_m.IBullet, 0, 38)
	leftSmallGunBullet := g_o.NewBullet(b.Cells()[174].Coords(), "#d48003ff", b)
	rightSmallGunBullet := g_o.NewBullet(b.Cells()[187].Coords(), "#d48003ff", b)

	bullets = append(bullets, leftSmallGunBullet, rightSmallGunBullet)

	for i := len(b.Cells()) - 1; i > len(b.Cells())-13; i-- {
		cell := b.Cells()[i]

		topBulletCoords := cell.Coords()
		middleBulletCoords := g_m.Coords{X: cell.Coords().X, Y: cell.Coords().Y + 1}
		bottomBulletCoords := g_m.Coords{X: cell.Coords().X, Y: cell.Coords().Y + 2}

		topBullet := g_o.NewBullet(topBulletCoords, "#d48003ff", b)
		middleBullet := g_o.NewBullet(middleBulletCoords, "#d48003ff", b)
		bottomBullet := g_o.NewBullet(bottomBulletCoords, "#d48003ff", b)

		bullets = append(bullets, topBullet, middleBullet, bottomBullet)
	}

	return bullets
}
