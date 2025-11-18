package game_objects

import (
	consts "github.com/pseudoelement/galaga/src/constants"
	g_c "github.com/pseudoelement/galaga/src/game/game-constants"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type HpBoost struct {
	*g_m.GameObject

	injector models.IAppInjector
}

func NewHpBoost(coords g_m.Coords, injector models.IAppInjector) *HpBoost {
	cells := []g_m.ICell{
		//1st
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y, "#08ae05ff", ""), g_c.BOOST_HP),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y, "#08ae05ff", ""), g_c.BOOST_HP),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y, "#08ae05ff", ""), g_c.BOOST_HP),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y, "#08ae05ff", ""), g_c.BOOST_HP),
		//2nd
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+1, "#08ae05ff", ""), g_c.BOOST_HP),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+1, "#08ae05ff", "H"), g_c.BOOST_HP),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+1, "#08ae05ff", "P"), g_c.BOOST_HP),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+1, "#08ae05ff", ""), g_c.BOOST_HP),
		//3th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+2, "#08ae05ff", ""), g_c.BOOST_HP),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+2, "#08ae05ff", "H"), g_c.BOOST_HP),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+2, "#08ae05ff", "P"), g_c.BOOST_HP),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+2, "#08ae05ff", ""), g_c.BOOST_HP),
		//4th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+3, "#08ae05ff", ""), g_c.BOOST_HP),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+3, "#08ae05ff", ""), g_c.BOOST_HP),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+3, "#08ae05ff", ""), g_c.BOOST_HP),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+3, "#08ae05ff", ""), g_c.BOOST_HP),
	}

	return &HpBoost{
		GameObject: g_m.NewGameObject(cells, g_c.BOOST_HP),
		injector:   injector,
	}
}

func (b *HpBoost) Name() string {
	return g_c.BOOST_HP
}

func (b *HpBoost) Boost(p g_m.IPlayer) {
	difficultyLvl := b.injector.Storage().GameDifficulty()
	var plusHP int16 = 1
	switch difficultyLvl {
	case consts.EASY:
		plusHP = 3
	case consts.MEDIUM:
		plusHP = 2
	case consts.HARD:
		plusHP = 1
	}

	p.GetHeal(plusHP)
}

func (e *HpBoost) MovementDelay(tickMs int) int {
	return tickMs * 12
}

var _ g_m.IBoost = (*HpBoost)(nil)
