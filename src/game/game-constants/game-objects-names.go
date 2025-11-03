package game_constants

// enemies
type EnemyType = string

const (
	OCTOPUS          = "OctopusEnemy"
	SMALL_SPACE_SHIP = "SmallSpaceShipEnemy"
	TNT              = "TNTEnemy"
)

var EMEMY_NAMES = [2]EnemyType{
	OCTOPUS,
	SMALL_SPACE_SHIP,
}

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
	BOOST_HP              = "HpBoost"
)

// others
const (
	BULLET       = "Bullet"
	BLAST_OF_TNT = "BlastOfTNT"
)
