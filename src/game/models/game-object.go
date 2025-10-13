package game_models

type IGameObject interface {
	Cells() []ICell
	/* use to get position before move and redraw bg without rerendering whole arena */
	PrevCells() []ICell
	Move(dir MoveDir)
	Destroy()
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////
type MoveDir struct {
	X int8
	Y int8
}

func MoveTop() MoveDir { return MoveDir{X: 0, Y: -1} }

func MoveBottom() MoveDir { return MoveDir{X: 0, Y: 1} }

func MoveLeft() MoveDir { return MoveDir{X: -1, Y: 0} }

func MoveRight() MoveDir { return MoveDir{X: 1, Y: 0} }

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
