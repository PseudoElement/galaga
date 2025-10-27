package player

import (
	game_constants "github.com/pseudoelement/galaga/src/game/game-constants"
	game_objects "github.com/pseudoelement/galaga/src/game/game-objects"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type TripleGunPlayer struct {
	*g_m.GameObject

	injector models.IAppInjector
	health   int16
}

func NewTripleGunPlayer(coords g_m.Coords, injector models.IAppInjector) *TripleGunPlayer {
	cells := []g_m.ICell{
		// 1st
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y, "#ff0505ff", "")),
		// 2nd
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+1, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+1, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+1, "#ff0505ff", "")),
		// 3th
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+2, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+2, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+2, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+2, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+2, "#ff0505ff", "")),
		// 4th
		g_m.NewCell(g_m.CellParams(coords.X, coords.Y+3, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+1, coords.Y+3, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+2, coords.Y+3, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+3, coords.Y+3, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+4, coords.Y+3, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+5, coords.Y+3, "#ff0505ff", "")),
		g_m.NewCell(g_m.CellParams(coords.X+6, coords.Y+3, "#ff0505ff", "")),
	}

	return &TripleGunPlayer{
		GameObject: g_m.NewGameObject(cells),
		health:     15,
	}
}

func (b *TripleGunPlayer) Name() string {
	return game_constants.PLAYER_TRIPLE_GUN
}

func (p *TripleGunPlayer) Shot() []g_m.IBullet {
	topLeftCell := p.Cells()[0]
	topMidCell := p.Cells()[1]
	topRightCell := p.Cells()[2]
	bulletLeft := game_objects.NewBullet(topLeftCell.Coords(), "#d7cc05ff", p)
	bulletMid := game_objects.NewBullet(topMidCell.Coords(), "#d7cc05ff", p)
	bulletRight := game_objects.NewBullet(topRightCell.Coords(), "#d7cc05ff", p)

	return []g_m.IBullet{bulletLeft, bulletMid, bulletRight}
}

func (p *TripleGunPlayer) Health() int16 {
	return p.health
}

func (p *TripleGunPlayer) GetHeal(plusHealthAmount int16) {
	p.health += plusHealthAmount
}

func (p *TripleGunPlayer) GetDamage(minusHealthAmount int16) {
	p.health -= minusHealthAmount

	if p.health <= 0 {
		p.Destroy()
	}
}

var _ g_m.IPlayer = (*TripleGunPlayer)(nil)
