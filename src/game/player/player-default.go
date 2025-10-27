package player

import (
	game_constants "github.com/pseudoelement/galaga/src/game/game-constants"
	game_objects "github.com/pseudoelement/galaga/src/game/game-objects"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type DefaultPlayer struct {
	*g_m.GameObject

	injector models.IAppInjector
	health   int16
}

func NewDefaultPlayer(coords g_m.Coords, injector models.IAppInjector) *DefaultPlayer {
	cells := []g_m.ICell{
		// 1st
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y, "#eb1eda", "")),
		// 2nd
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+1, "#eb1eda", "")),
		// 3th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+2, "#eb1eda", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+2, "#eb1eda", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+2, "#eb1eda", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+2, "#eb1eda", "")),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+2, "#eb1eda", "")),
		// 4th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+3, "#eb1eda", "")),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+3, "#eb1eda", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+3, "#eb1eda", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+3, "#eb1eda", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+3, "#eb1eda", "")),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+3, "#eb1eda", "")),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+3, "#eb1eda", "")),
	}

	return &DefaultPlayer{
		GameObject: g_m.NewGameObject(cells),
		injector:   injector,
		health:     5,
	}
}

func (b *DefaultPlayer) Name() string {
	return game_constants.PLAYER_DEFAULT
}

func (b *DefaultPlayer) Tier() g_m.PlayerTier {
	return 1
}

func (p *DefaultPlayer) Shot() []g_m.IBullet {
	// bullets thrown from top cells of the ship
	topMidCell := p.Cells()[0]
	bullet := game_objects.NewBullet(topMidCell.Coords(), "#d7cc05ff", p)

	return []g_m.IBullet{bullet}
}

func (p *DefaultPlayer) Health() int16 {
	return p.health
}

func (p *DefaultPlayer) GetHeal(plusHealthAmount int16) {
	p.health += plusHealthAmount
}

func (p *DefaultPlayer) GetDamage(minusHealthAmount int16) {
	p.health -= minusHealthAmount

	if p.health <= 0 {
		p.Destroy()
	}
}

var _ g_m.IPlayer = (*DefaultPlayer)(nil)
