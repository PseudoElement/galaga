package game_constants

// enemies
const (
	OCTOPUS          = "OctopusEnemy"
	SMALL_SPACE_SHIP = "SmallSpaceShipEnemy"
)

// players
type PlayerType = string

const (
	PLAYER_DEFAULT    = "DefaultPlayer"
	PLAYER_DOBLE_GUN  = "DoubleGunPlayer"
	PLAYER_TRIPLE_GUN = "TripleGunPlayer"
)

var PLAYER_NAMES = [3]PlayerType{
	PLAYER_DEFAULT,
	PLAYER_DOBLE_GUN,
	PLAYER_TRIPLE_GUN,
}

// boosts
const (
	BOOST_DOUBLE_GUN_SHIP = "DoubleGunShipBoost"
	BOOST_TRIPLE_GUN_SHIP = "TripleGunShipBoost"
)

// others
const (
	BULLET = "Bullet"
)
