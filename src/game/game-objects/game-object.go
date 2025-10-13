package game_objects

import (
	game_models "github.com/pseudoelement/galaga/src/game/models"
)

type GameObject struct {
	cells     []game_models.ICell
	prevCells []game_models.ICell
}

func NewGameObject(cells []game_models.ICell) *GameObject {
	prevCells := make([]game_models.ICell, len(cells))
	copy(prevCells, cells)
	return &GameObject{cells: cells, prevCells: prevCells}
}

func (obj *GameObject) Cells() []game_models.ICell {
	return obj.cells
}

func (obj *GameObject) PrevCells() []game_models.ICell {
	return obj.prevCells
}

func (obj *GameObject) Move(dir game_models.MoveDir) {
	for _, cell := range obj.cells {
		prevCoords := cell.Coords()
		// obj.prevCells[idx].SetCoords(prevCoords)
		cell.SetCoords(game_models.Coords{
			X: prevCoords.X + dir.X,
			Y: prevCoords.Y + dir.Y,
		})
	}
}

func (obj *GameObject) Destroy() {
	obj.cells = nil
	obj = nil
}

var _ game_models.IGameObject = (*GameObject)(nil)
