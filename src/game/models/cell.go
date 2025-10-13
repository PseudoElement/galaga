package game_models

type Coords struct {
	X int8
	Y int8
}

type ICell interface {
	Color() string
	SetColor(color string)
	Coords() Coords
	SetCoords(coords Coords)
}

type CellConstructorParams struct {
	/* hex */
	Color  string
	Coords Coords
	// CanDamage bool
}

type ICellWithDamage interface {
	ICell
	CanDamage() bool
	Damage(player IPlayer)
	SetDamageCount(damageCount int16)
}
