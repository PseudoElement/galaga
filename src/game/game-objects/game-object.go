package game_objects

import (
	game_models "github.com/pseudoelement/galaga/src/game/models"
)

type GameObject struct {
	cells     []game_models.ICell
	prevCells []game_models.ICell
	destroyed bool
}

func NewGameObject(cells []game_models.ICell) *GameObject {
	// Create deep copies of each cell
	prevCells := make([]game_models.ICell, len(cells))
	for i, cell := range cells {
		// Create a new cell with the same properties
		prevCells[i] = NewCell(game_models.CellConstructorParams{
			Color:  cell.Color(),
			Coords: cell.Coords(),
		})
	}

	return &GameObject{cells: cells, prevCells: prevCells, destroyed: false}
}

func (obj *GameObject) Cells() []game_models.ICell {
	return obj.cells
}

func (obj *GameObject) PrevCells() []game_models.ICell {
	return obj.prevCells
}

func (obj *GameObject) Move(dir game_models.MoveDir) {
	for idx, cell := range obj.cells {
		prevCoords := cell.Coords()
		prevCellState := obj.prevCells[idx]
		prevCellState.SetCoords(prevCoords)

		cell.SetCoords(game_models.Coords{
			X: prevCoords.X + dir.X,
			Y: prevCoords.Y + dir.Y,
		})
	}
}

func (obj *GameObject) Destroy() {
	obj.destroyed = true
}

func (obj *GameObject) Destroyed() bool {
	return obj.destroyed
}

var _ game_models.IGameObject = (*GameObject)(nil)
