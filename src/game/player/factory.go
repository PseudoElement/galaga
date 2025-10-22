package player

import (
	"fmt"

	game_models "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

func CreatePlayer(coords game_models.Coords, playerType models.PlayerType, injector models.IAppInjector) game_models.IPlayer {
	switch playerType {
	case models.DEFAULT_PLAYER:
		return NewDefaultPlayer(coords, injector)
	case models.DOUBLE_GUN_PLAYER:
		return NewDoubleGunPlayer(coords, injector)
	default:
		err := fmt.Sprintf("%s is unknown playerType.", playerType)
		panic(err)
	}
}
