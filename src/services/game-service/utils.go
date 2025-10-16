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

func isObjectOutOfArena(obj game_models.IGameObject, arenaWidth, arenaHeight int) bool {
	for _, cell := range obj.Cells() {
		coords := cell.Coords()
		if coords.X < 0 ||
			coords.Y < 0 ||
			coords.X > int16(arenaWidth) ||
			coords.Y > int16(arenaHeight) {
			return true
		}
	}
	return false
}

func isCellOutOfArena(cell game_models.ICell, arenaWidth, arenaHeight int) bool {
	coords := cell.Coords()
	if coords.X < 0 ||
		coords.Y < 0 ||
		coords.X > int16(arenaWidth) ||
		coords.Y > int16(arenaHeight) {
		return true
	}
	return false
}
