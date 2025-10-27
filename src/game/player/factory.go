package player

import (
	"fmt"

	game_constants "github.com/pseudoelement/galaga/src/game/game-constants"
	game_models "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

func CreatePlayer(coords game_models.Coords, playerType game_constants.PlayerType, injector models.IAppInjector) game_models.IPlayer {
	switch playerType {
	case game_constants.PLAYER_DEFAULT:
		return NewDefaultPlayer(coords, injector)
	case game_constants.PLAYER_DOBLE_GUN:
		return NewDoubleGunPlayer(coords, injector)
	case game_constants.PLAYER_TRIPLE_GUN:
		return NewTripleGunPlayer(coords, injector)
	default:
		err := fmt.Sprintf("%s is unknown playerType.", playerType)
		panic(err)
	}
}
