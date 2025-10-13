package player

import (
	"math"

	game_objects "github.com/pseudoelement/galaga/src/game/game-objects"
	game_models "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type DefalutPlayer struct {
	*game_objects.GameObject

	injector models.IAppInjector
}

func NewDefaultPlayer(injector models.IAppInjector) *DefalutPlayer {
	width := injector.Storage().WindowSize().Width
	height := injector.Storage().WindowSize().Height

	playerTopY := int8(height - 5)
	playerLeftX := int8(math.Floor(float64(width) / 2))

	cellParams := func(x, y int8) game_models.CellConstructorParams {
		return game_models.CellConstructorParams{
			Color:  "#eb1eda",
			Coords: game_models.Coords{X: x, Y: y},
		}
	}

	cells := []game_models.ICell{
		game_objects.NewCell(cellParams(playerLeftX+2, playerTopY)), game_objects.NewCell(cellParams(playerLeftX+3, playerTopY)),
		game_objects.NewCell(cellParams(playerLeftX+2, playerTopY+1)), game_objects.NewCell(cellParams(playerLeftX+3, playerTopY+1)),
		game_objects.NewCell(cellParams(playerLeftX+2, playerTopY+2)), game_objects.NewCell(cellParams(playerLeftX+3, playerTopY+2)),
		game_objects.NewCell(cellParams(playerLeftX, playerTopY+3)), game_objects.NewCell(cellParams(playerLeftX+1, playerTopY+3)), game_objects.NewCell(cellParams(playerLeftX+2, playerTopY+3)), game_objects.NewCell(cellParams(playerLeftX+3, playerTopY+3)), game_objects.NewCell(cellParams(playerLeftX+4, playerTopY+3)), game_objects.NewCell(cellParams(playerLeftX+5, playerTopY+3)),
	}

	return &DefalutPlayer{
		GameObject: game_objects.NewGameObject(cells),
		injector:   injector,
	}
}

func (dp *DefalutPlayer) Shot()                                     {}
func (dp *DefalutPlayer) TakeBoost(boostItem game_models.BoostItem) {}
func (dp *DefalutPlayer) Health() int16 {
	return 1
}
func (dp *DefalutPlayer) GetHeal(plusHealthAmount int16)    {}
func (dp *DefalutPlayer) GetDamage(minusHealthAmount int16) {}

var _ game_models.IPlayer = (*DefalutPlayer)(nil)
