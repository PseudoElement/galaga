package player

import (
	"math"

	g_o "github.com/pseudoelement/galaga/src/game/game-objects"
	game_objects "github.com/pseudoelement/galaga/src/game/game-objects"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type DefaultPlayer struct {
	*g_o.GameObject

	injector models.IAppInjector
}

func NewDefaultPlayer(injector models.IAppInjector) *DefaultPlayer {
	width := injector.Storage().WindowSize().Width
	height := injector.Storage().WindowSize().Height

	playerTopY := int16(height - 5)
	playerLeftX := int16(math.Floor(float64(width) / 2))

	cells := []g_m.ICell{
		// 1st
		g_o.NewCell(g_m.CellParams(playerLeftX+3, playerTopY, "#eb1eda")),
		// 2nd
		g_o.NewCell(g_m.CellParams(playerLeftX+3, playerTopY+1, "#eb1eda")),
		// 3th
		g_o.NewCell(g_m.CellParams(playerLeftX+3, playerTopY+2, "#eb1eda")),
		// 4th
		g_o.NewCell(g_m.CellParams(playerLeftX, playerTopY+3, "#eb1eda")),
		g_o.NewCell(g_m.CellParams(playerLeftX+1, playerTopY+3, "#eb1eda")),
		g_o.NewCell(g_m.CellParams(playerLeftX+2, playerTopY+3, "#eb1eda")),
		g_o.NewCell(g_m.CellParams(playerLeftX+3, playerTopY+3, "#eb1eda")),
		g_o.NewCell(g_m.CellParams(playerLeftX+4, playerTopY+3, "#eb1eda")),
		g_o.NewCell(g_m.CellParams(playerLeftX+5, playerTopY+3, "#eb1eda")),
		g_o.NewCell(g_m.CellParams(playerLeftX+6, playerTopY+3, "#eb1eda")),
	}

	return &DefaultPlayer{
		GameObject: g_o.NewGameObject(cells),
		injector:   injector,
	}
}

func (dp *DefaultPlayer) Shot() []g_m.IBullet {
	// bullets thrown from top cells of the ship
	topMidCell := dp.Cells()[0]
	bullet := game_objects.NewPlayerBullet(topMidCell.Coords(), "#d7cc05ff")

	return []g_m.IBullet{bullet}
}

func (dp *DefaultPlayer) TakeBoost(boostItem g_m.BoostItem) {}

func (dp *DefaultPlayer) Health() int16 {
	return 1
}

func (dp *DefaultPlayer) GetHeal(plusHealthAmount int16) {}

func (dp *DefaultPlayer) GetDamage(minusHealthAmount int16) {}

var _ g_m.IPlayer = (*DefaultPlayer)(nil)
