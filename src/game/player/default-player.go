package player

import (
	"math"

	game_objects "github.com/pseudoelement/galaga/src/game/game-objects"
	game_models "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type DefaultPlayer struct {
	*game_objects.GameObject

	injector models.IAppInjector
}

func NewDefaultPlayer(injector models.IAppInjector) *DefaultPlayer {
	width := injector.Storage().WindowSize().Width
	height := injector.Storage().WindowSize().Height

	playerTopY := int16(height - 5)
	playerLeftX := int16(math.Floor(float64(width) / 2))

	cellParams := func(x, y int16) game_models.CellConstructorParams {
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

	return &DefaultPlayer{
		GameObject: game_objects.NewGameObject(cells),
		injector:   injector,
	}
}

func (dp *DefaultPlayer) Shot()                                     {}
func (dp *DefaultPlayer) TakeBoost(boostItem game_models.BoostItem) {}
func (dp *DefaultPlayer) Health() int16 {
	return 1
}
func (dp *DefaultPlayer) GetHeal(plusHealthAmount int16)    {}
func (dp *DefaultPlayer) GetDamage(minusHealthAmount int16) {}

var _ game_models.IPlayer = (*DefaultPlayer)(nil)
