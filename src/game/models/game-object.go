package game_models

type GameObject struct {
	cells       []ICell
	prevCells   []ICell
	prevMoveDir MoveDir
}

func NewGameObject(cells []ICell) *GameObject {
	// Create deep copies of each cell
	prevCells := make([]ICell, len(cells))
	for i, cell := range cells {
		// Create a new cell with the same properties
		prevCells[i] = NewCell(CellConstructorParams{
			Color:  cell.Color(),
			Coords: cell.Coords(),
		})
	}

	return &GameObject{
		cells:       cells,
		prevCells:   prevCells,
		prevMoveDir: MoveDir{X: 0, Y: 0},
	}
}

func (obj *GameObject) Cells() []ICell {
	return obj.cells
}

func (obj *GameObject) PrevCells() []ICell {
	return obj.prevCells
}

func (obj *GameObject) PrevMoveDir() MoveDir {
	return obj.prevMoveDir
}

func (obj *GameObject) Move(dir MoveDir) {
	obj.prevMoveDir = dir
	for idx, cell := range obj.cells {
		prevCoords := cell.Coords()
		prevCellState := obj.prevCells[idx]
		prevCellState.SetCoords(prevCoords)

		cell.SetCoords(Coords{
			X: prevCoords.X + dir.X,
			Y: prevCoords.Y + dir.Y,
		})
	}
}

func (obj *GameObject) Destroy() {
	for _, cell := range obj.cells {
		cell.Destroy()
	}
}

func (obj *GameObject) Destroyed() bool {
	for _, cell := range obj.cells {
		if !cell.Destroyed() {
			return false
		}
	}
	return true
}

// var _ IGameObject = (*GameObject)(nil)
