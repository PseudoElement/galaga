package models

import game_models "github.com/pseudoelement/galaga/src/game/models"

type IAppGameSrv interface {
	View() string
	StartGame()
	EndGame()
	Player() game_models.IPlayer
	SetPlayer(player game_models.IPlayer)
}
