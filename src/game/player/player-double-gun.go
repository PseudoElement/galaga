package player

import (
	game_constants "github.com/pseudoelement/galaga/src/game/game-constants"
	game_objects "github.com/pseudoelement/galaga/src/game/game-objects"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type DoubleGunPlayer struct {
	*g_m.GameObject

	injector models.IAppInjector
	health   int16
}

func NewDoubleGunPlayer(coords g_m.Coords, injector models.IAppInjector) *DoubleGunPlayer {
	cells := []g_m.ICell{
		// 1st
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y, "#2db9f0ff", "")),
		// 2nd
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+1, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+1, "#2db9f0ff", "")),
		// 3th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+2, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+2, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+2, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+2, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+2, "#2db9f0ff", "")),
		// 4th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+3, "#2db9f0ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+3, "#2db9f0ff", "")),
	}
	return &DoubleGunPlayer{
		GameObject: g_m.NewGameObject(cells),
		health:     10,
	}
}

func (b *DoubleGunPlayer) Name() string {
	return game_constants.PLAYER_DOBLE_GUN
}

func (b *DoubleGunPlayer) Tier() g_m.PlayerTier {
	return 2
}

func (p *DoubleGunPlayer) Shot() []g_m.IBullet {
	topLeftCell := p.Cells()[0]
	topRightCell := p.Cells()[1]
	bulletLeft := game_objects.NewBullet(topLeftCell.Coords(), "#d7cc05ff", p)
	bulletRight := game_objects.NewBullet(topRightCell.Coords(), "#d7cc05ff", p)

	return []g_m.IBullet{bulletLeft, bulletRight}
}

func (p *DoubleGunPlayer) Health() int16 {
	return p.health
}

func (p *DoubleGunPlayer) GetHeal(plusHealthAmount int16) {
	p.health += plusHealthAmount
}

func (p *DoubleGunPlayer) GetDamage(minusHealthAmount int16) {
	p.health -= minusHealthAmount

	if p.health <= 0 {
		p.Destroy()
	}
}

var _ g_m.IPlayer = (*DoubleGunPlayer)(nil)
