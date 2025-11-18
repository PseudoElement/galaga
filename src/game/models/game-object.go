package game_models

import (
	"math"
)

type GameObject struct {
	cells       []ICell
	prevCells   []ICell
	prevMoveDir MoveDir

	width  int16
	height int16
}

func NewGameObject(cells []ICell, name string) *GameObject {
	// Create deep copies of each cell
	prevCells := make([]ICell, len(cells))
	for i, cell := range cells {
		// Create a new cell with the same properties
		prevCells[i] = NewCell(
			CellConstructorParams{
				Color:  cell.Color(),
				Coords: cell.Coords(),
			},
			name,
		)
	}

	object := &GameObject{
		cells:       cells,
		prevCells:   prevCells,
		prevMoveDir: MoveDir{X: 0, Y: 0},
	}
	object.setSize()

	return object
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
	if dir == MoveNoneX0_Y0() {
		return
	}

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

func (obj *GameObject) Size() (width, height int16) {
	return obj.width, obj.height
}

func (obj *GameObject) setSize() {
	leftX, rightX, topY, bottomY := int16(math.MaxInt16), int16(math.MinInt16), int16(math.MaxInt16), int16(math.MinInt16)
	for _, cell := range obj.Cells() {
		if cell.Coords().X < leftX {
			leftX = cell.Coords().X
		}
		if cell.Coords().X > rightX {
			rightX = cell.Coords().X
		}
		if cell.Coords().Y < topY {
			topY = cell.Coords().Y
		}
		if cell.Coords().Y > bottomY {
			bottomY = cell.Coords().Y
		}
	}

	obj.width = rightX - leftX + 1
	obj.height = bottomY - topY + 1
}
