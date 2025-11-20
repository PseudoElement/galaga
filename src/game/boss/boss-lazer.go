package boss

import (
	"github.com/pseudoelement/galaga/src/game/enemy"
	g_c "github.com/pseudoelement/galaga/src/game/game-constants"
	g_o "github.com/pseudoelement/galaga/src/game/game-objects"
	g_u "github.com/pseudoelement/galaga/src/game/game-utils"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

type LazerBoss struct {
	*enemy.Enemy
	injector models.IAppInjector
}

func NewLazerBoss(x, y int16, health int16, injector models.IAppInjector) g_m.IBossEnemy {
	cells := []g_m.ICell{
		// 1st
		g_m.NewCell(g_m.CellParams(x, y, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+1, y, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+2, y, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+3, y, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+4, y, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+5, y, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+6, y, "#19a053ff", ""), g_c.BOSS_LAZER),
		// 2nd
		g_m.NewCell(g_m.CellParams(x, y+1, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+1, y+1, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+2, y+1, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+3, y+1, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+4, y+1, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+5, y+1, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+6, y+1, "#19a053ff", ""), g_c.BOSS_LAZER),
		// 3th
		g_m.NewCell(g_m.CellParams(x+2, y+2, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+3, y+2, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+4, y+2, "#19a053ff", ""), g_c.BOSS_LAZER),
		// 4th
		g_m.NewCell(g_m.CellParams(x+2, y+3, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+4, y+3, "#19a053ff", ""), g_c.BOSS_LAZER),
		// 5th
		g_m.NewCell(g_m.CellParams(x+2, y+4, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+4, y+4, "#19a053ff", ""), g_c.BOSS_LAZER),
		// 6th
		g_m.NewCell(g_m.CellParams(x+2, y+5, "#19a053ff", ""), g_c.BOSS_LAZER),
		g_m.NewCell(g_m.CellParams(x+4, y+5, "#19a053ff", ""), g_c.BOSS_LAZER),
		// 7th
		g_m.NewCell(g_m.CellParams(x+2, y+6, "#19a053ff", ""), g_c.BOSS_LAZER), // left gun
		g_m.NewCell(g_m.CellParams(x+4, y+6, "#19a053ff", ""), g_c.BOSS_LAZER), // right gun
	}

	movementPattern := []g_m.MoveDir{g_m.MoveNoneX0_Y0()}

	return &LazerBoss{
		Enemy:    enemy.NewEnemy(health, cells, movementPattern, g_c.BOSS_LAZER),
		injector: injector,
	}
}

func (b *LazerBoss) Name() string {
	return g_c.BOSS_LAZER
}

func (b *LazerBoss) Price() int32 {
	return 400
}

func (b *LazerBoss) MovementDelay(tickMs int) int {
	return tickMs * 5
}

func (b *LazerBoss) IsBoss() bool {
	return true
}

func (e *LazerBoss) ShootingDelay(tickMs int) int {
	return tickMs * 200
}

func (b *LazerBoss) UpdateMovementPattern(player g_m.IPlayer) {
	playerMidCoords := g_u.GetObjectMidCoords(player)
	bossMidCoords := g_u.GetObjectMidCoords(b)

	if playerMidCoords.X > bossMidCoords.X {
		b.ChangeMovementPattern([]g_m.MoveDir{g_m.MoveRightX1_Y0()})
	} else if playerMidCoords.X < bossMidCoords.X {
		b.ChangeMovementPattern([]g_m.MoveDir{g_m.MoveLeftX1_Y0()})
	} else {
		b.ChangeMovementPattern([]g_m.MoveDir{g_m.MoveNoneX0_Y0()})
	}
}

func (b *LazerBoss) Shot() []g_m.IBullet {
	/*
	* shipHeight = 7
	* paddingFromLazerEndToArenaBottom = 4
	* bulettsPerGun = arenaHeight - shipHeight - paddingFromLazerEndToArenaBottom
	* gunsCount = 2
	* bulletsCount = gunsCount * bulettsPerGun
	 */
	_, arenaHeight := b.injector.GameSrv().ArenaSize()

	paddingFromLazerEndToArenaBottom := 4
	shipHeight := 7
	bulletsPerGun := arenaHeight - shipHeight - paddingFromLazerEndToArenaBottom
	bulletsCount := 2 * bulletsPerGun

	leftGunCoord := b.Cells()[len(b.Cells())-2].Coords()
	rightGunCoord := b.Cells()[len(b.Cells())-1].Coords()

	bullets := make([]g_m.IBullet, 0, bulletsCount)

	for bulletCoordY := shipHeight; bulletCoordY < arenaHeight-paddingFromLazerEndToArenaBottom; bulletCoordY++ {
		leftGunBullet := g_o.NewBullet(g_m.Coords{X: leftGunCoord.X, Y: int16(bulletCoordY)}, "#cf1374", b.Name())
		rightGunBullet := g_o.NewBullet(g_m.Coords{X: rightGunCoord.X, Y: int16(bulletCoordY)}, "#cf1374", b.Name())

		bullets = append(bullets, leftGunBullet, rightGunBullet)

		bulletCoordY++
	}

	return bullets
}

var _ g_m.IBossEnemy = (*LazerBoss)(nil)
var _ g_m.IDirectionChanger = (*LazerBoss)(nil)
