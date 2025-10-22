package player

import (
	game_objects "github.com/pseudoelement/galaga/src/game/game-objects"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type DoubleGunPlayer struct {
	*g_m.GameObject

	injector models.IAppInjector
	health   int16
}

func NewDoubleGunPlayer(coords g_m.Coords, injector models.IAppInjector) *DoubleGunPlayer {
	cells := []g_m.ICell{
		// 1st
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y, "#2db9f0ff", "")),
		// 2nd
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+1, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+1, "#2db9f0ff", "")),
		// 3th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+2, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+2, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+2, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+2, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+2, "#2db9f0ff", "")),
		// 4th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+3, "#2db9f0ff", "")),
	}
	return &DoubleGunPlayer{GameObject: g_m.NewGameObject(cells)}
}

func (dp *DoubleGunPlayer) Shot() []g_m.IBullet {
	topLeftCell := dp.Cells()[0]
	topRightCell := dp.Cells()[1]
	bulletLeft := game_objects.NewPlayerBullet(topLeftCell.Coords(), "#d7cc05ff")
	bulletRight := game_objects.NewPlayerBullet(topRightCell.Coords(), "#d7cc05ff")

	return []g_m.IBullet{bulletLeft, bulletRight}
}

func (dp *DoubleGunPlayer) TakeBoost(boostItem g_m.IBoost) {}

func (dp *DoubleGunPlayer) Health() int16 {
	return dp.health
}

func (dp *DoubleGunPlayer) GetHeal(plusHealthAmount int16) {
	dp.health += plusHealthAmount
}

func (dp *DoubleGunPlayer) GetDamage(minusHealthAmount int16) {
	dp.health -= minusHealthAmount
}

var _ g_m.IPlayer = (*DoubleGunPlayer)(nil)
