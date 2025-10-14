package game_srv

import game_models "github.com/pseudoelement/galaga/src/game/models"

func cellParams(x, y int16, color string) game_models.CellConstructorParams {
	return game_models.CellConstructorParams{
		Color:  color,
		Coords: game_models.Coords{X: x, Y: y},
	}
}

func cellsToView(cells []game_models.ICell) []string {
	view := make([]string, len(cells))
	for idx, cell := range cells {
		view[idx] = cell.Style().Render()
	}
	return view
}
