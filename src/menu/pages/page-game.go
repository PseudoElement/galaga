package pages

import (
	"math"

	game_models "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/game/player"
	"github.com/pseudoelement/galaga/src/models"
)

type PageGame struct {
	*Page
}

func NewPageGame(injector models.IAppInjector) models.IPage {
	p := &PageGame{Page: NewPage(injector)}

	width := injector.Storage().WindowSize().Width
	height := injector.Storage().WindowSize().Height

	playerTopY := int16(height - 5)
	playerLeftX := int16(math.Floor(float64(width) / 2))

	injector.GameSrv().SetPlayer(player.NewDefaultPlayer(game_models.Coords{X: playerLeftX, Y: playerTopY}, injector))
	injector.GameSrv().StartGame()

	return p
}

func (p *PageGame) View() string {
	if p.injector.GameSrv().IsPlaying() {
		return p.injector.GameSrv().View()
	}
	return ""
}
