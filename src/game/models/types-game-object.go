package game_models

type IGameObject interface {
	IDestroyable
	IMoveable
	Cells() []ICell
	/* use to get position before move and redraw bg without rerendering whole arena */
	PrevCells() []ICell
	Name() string
	Size() (width, height int16)
}

type IGameObjectWithHP interface {
	IGameObject
	IHealable
	IDamageable
}

type IMoveable interface {
	Move(dir MoveDir)
	PrevMoveDir() MoveDir
}

type IAutoMoveable interface {
	IMoveable
	MovementDelay(tickMs int) int
}

type IDestroyable interface {
	Destroy()
	Destroyed() bool
}

type IShooter interface {
	Shot() []IBullet
}

type IAutoShooter interface {
	IShooter
	ShootingDelay(tickMs int) int
}

type IHealable interface {
	Health() int16
	GetHeal(plusHealthAmount int16)
}

type IDamageable interface {
	GetDamage(minusHealthAmount int16)
}

type IOwnerable interface {
	Owner() IGameObject
}

type IWithPrice interface {
	Price() int32
}

type IAutoDestructable interface {
	DestructInMs() int
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////
type MoveDir struct {
	X int16
	Y int16
}

func MoveNoneX0_Y0() MoveDir { return MoveDir{X: 0, Y: 0} }

func MoveTopX0_Y1() MoveDir { return MoveDir{X: 0, Y: -1} }

func MoveBottomX0_Y1() MoveDir { return MoveDir{X: 0, Y: 1} }

func MoveLeftX1_Y0() MoveDir { return MoveDir{X: -1, Y: 0} }

func MoveRightX1_Y0() MoveDir { return MoveDir{X: 1, Y: 0} }

func MoveLeftX2_Y0() MoveDir { return MoveDir{X: -2, Y: 0} }

func MoveRightX2_Y0() MoveDir { return MoveDir{X: 2, Y: 0} }

func MoveTopX0_Y3() MoveDir { return MoveDir{X: 0, Y: -3} }

func MoveBottomX0_Y3() MoveDir { return MoveDir{X: 0, Y: 3} }

func MoveLeftX5_Y0() MoveDir { return MoveDir{X: -5, Y: 0} }

func MoveRightX5_Y0() MoveDir { return MoveDir{X: 5, Y: 0} }

func MoveLeftBottomX1_Y1() MoveDir { return MoveDir{X: -1, Y: 1} }

func MoveRightBottomX1_Y1() MoveDir { return MoveDir{X: 1, Y: 1} }

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
