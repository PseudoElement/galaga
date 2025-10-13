package pages

import (
	"github.com/pseudoelement/galaga/src/game/player"
	"github.com/pseudoelement/galaga/src/models"
)

type PageGame struct {
	*Page
}

func NewPageGame(injector models.IAppInjector) models.IPage {
	p := &PageGame{Page: NewPage(injector)}

	injector.GameSrv().SpawnPlayer(player.NewDefaultPlayer(injector))
	injector.GameSrv().StartGame()

	return p
}

func (p *PageGame) View() string {
	return p.injector.GameSrv().View()
}
