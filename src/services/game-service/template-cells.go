package game_srv

import (
	game_models "github.com/pseudoelement/galaga/src/game/models"
	game_srv_styles "github.com/pseudoelement/galaga/src/services/game-service/styles"
)

func blackCell(cell game_models.ICell) *game_models.Cell {
	return game_models.NewCell(
		cellParams(
			cell.Coords().X,
			cell.Coords().Y,
			game_srv_styles.BlackColor,
		),
		"",
	)
}
