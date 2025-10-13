package models

import game_models "github.com/pseudoelement/galaga/src/game/models"

type IAppGameSrv interface {
	View() string
	StartGame()
	EndGame()
	SpawnPlayer(player game_models.IPlayer)
	Player() game_models.IPlayer
}
