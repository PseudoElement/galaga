package game_srv

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/pseudoelement/galaga/src/game/enemy"
	g_c "github.com/pseudoelement/galaga/src/game/game-constants"
	g_o "github.com/pseudoelement/galaga/src/game/game-objects"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
	game_srv_styles "github.com/pseudoelement/galaga/src/services/game-service/styles"
)

const (
	GAME_LOOP_TICK_DELAY_MS = 20
)

type AppGameSrv struct {
	injector     models.IAppInjector
	enemyFactory g_m.IEnemyFactory

	arena          [][]g_m.ICell
	gameDurationMs int
	objectsPool    []g_m.IGameObject
	player         g_m.IPlayer
	score          int32
	stop           bool
}

func NewAppGameSrv(injector models.IAppInjector, player g_m.IPlayer) models.IAppGameSrv {
	return &AppGameSrv{
		injector:     injector,
		enemyFactory: enemy.NewEnemyFactory(injector),

		arena:          make([][]g_m.ICell, 0),
		objectsPool:    make([]g_m.IGameObject, 0),
		score:          0,
		player:         player,
		stop:           false,
		gameDurationMs: 0,
	}
}

func (gs *AppGameSrv) Player() g_m.IPlayer {
	return gs.player
}

func (gs *AppGameSrv) View() string {
	// align rows horizontally
	rows := make([]string, 0)
	for _, arenaRow := range gs.arena {
		row := lipgloss.JoinHorizontal(lipgloss.Left, cellsToView(arenaRow)...)
		rows = append(rows, row)
	}

	arenaView := lipgloss.JoinVertical(lipgloss.Top, rows...)
	hpView := fmt.Sprintf("❤️ 1")
	scoreView := fmt.Sprintf(" Score: %d", gs.score)

	header := lipgloss.JoinHorizontal(lipgloss.Center, hpView, scoreView)
	view := lipgloss.JoinVertical(lipgloss.Center, header, arenaView)

	return view
}

func (gs *AppGameSrv) StartGame() {
	// at this point gs.injector.Storage().WindowSize() should be already defined
	width, height := gs.getArenaSize()
	for heightCount := range height { // y
		arenaRow := make([]g_m.ICell, 0, width)
		for widthCount := range width { // x
			arenaRow = append(arenaRow, g_o.NewCell(cellParams(
				int16(widthCount),
				int16(heightCount),
				game_srv_styles.BlackColor,
			)))
		}

		gs.arena = append(gs.arena, arenaRow)
	}

	// render player's starship
	for _, cell := range gs.player.Cells() {
		coords := cell.Coords()
		gs.arena[coords.Y][coords.X] = cell
	}

	go gs.runLoop()
}

func (gs *AppGameSrv) EndGame() {
	gs.stop = true
	gs.player = nil
	gs.objectsPool = make([]g_m.IGameObject, 0)
	gs.arena = make([][]g_m.ICell, 0)
	gs.gameDurationMs = 0
	gs.resetScore()
}

func (gs *AppGameSrv) SetPlayer(player g_m.IPlayer) {
	gs.player = player
}

func (gs *AppGameSrv) runLoop() {
	x, y := gs.getArenaSize()
	for !gs.stop && gs.player != nil {
		time.Sleep(GAME_LOOP_TICK_DELAY_MS * time.Millisecond)
		gs.injector.TeaProgram().Send(models.UpdateTrigger{})
		gs.gameDurationMs += GAME_LOOP_TICK_DELAY_MS

		gs.clearPrevCellsOfObjectsOnTick()
		gs.drawNewCellsOfObjectsOnTick()

		gs.handleEnemySpawn()
		gs.handleShot()

		for _, object := range gs.objectsPool {
			switch obj := object.(type) {
			case g_m.IBullet:
				gs.handleBulletBehaviourOnTick(obj)
			case g_m.IEnemy:
				gs.handleEnemyBehaviourOnTick(obj)
			}
		}

		// destroy cells out of arena
		for _, object := range gs.objectsPool {
			for _, cell := range object.Cells() {
				if isCellOutOfArena(cell, x, y) {
					cell.Destroy()
				}
			}
		}
	}
}

func (gs *AppGameSrv) clearPrevCellsOfObjectsOnTick() {
	width, height := gs.getArenaSize()
	updatedObjectsPool := make([]g_m.IGameObject, 0)
	for _, obj := range gs.objectsPool {
		if !obj.Destroyed() {
			updatedObjectsPool = append(updatedObjectsPool, obj)
		}
		for _, cell := range obj.PrevCells() {
			coords := cell.Coords()
			if !isCellOutOfArena(cell, width, height) {
				gs.arena[coords.Y][coords.X] = g_o.NewCell(cellParams(coords.X, coords.Y, game_srv_styles.BlackColor))
			}
		}
	}

	gs.objectsPool = updatedObjectsPool

	for _, cell := range gs.player.PrevCells() {
		coords := cell.Coords()
		gs.arena[coords.Y][coords.X] = g_o.NewCell(cellParams(coords.X, coords.Y, game_srv_styles.BlackColor))
	}
}

func (gs *AppGameSrv) drawNewCellsOfObjectsOnTick() {
	width, height := gs.getArenaSize()
	for _, obj := range gs.objectsPool {
		for _, cell := range obj.Cells() {
			coords := cell.Coords()
			if !isCellOutOfArena(cell, width, height) {
				gs.arena[coords.Y][coords.X] = cell
			}
		}
	}
	for _, cell := range gs.player.Cells() {
		coords := cell.Coords()
		gs.arena[coords.Y][coords.X] = cell
	}
}

func (gs *AppGameSrv) handleEnemySpawn() {
	spawnLatency := GAME_LOOP_TICK_DELAY_MS * 150
	if gs.gameDurationMs%spawnLatency == 0 {
		spawnedEnemy := gs.enemyFactory.SpawnEnemy(g_c.EASY)
		gs.objectsPool = append(gs.objectsPool, spawnedEnemy)

		gs.increaseScore()
	}
}

func (gs *AppGameSrv) handleShot() {
	shotLatency := GAME_LOOP_TICK_DELAY_MS * 16
	if gs.gameDurationMs%shotLatency == 0 {
		bullets := gs.player.Shot()
		for _, bullet := range bullets {
			gs.objectsPool = append(gs.objectsPool, bullet)
		}
	}
}

func (gs *AppGameSrv) handleEnemyBehaviourOnTick(enemy g_m.IEnemy) {
	width, height := gs.getArenaSize()
	movementDelay := enemy.MovementDelay(GAME_LOOP_TICK_DELAY_MS)
	if gs.gameDurationMs%movementDelay == 0 {
		enemy.Move(enemy.MoveDir(width, height))
	}

	enemyShooter, isShooter := enemy.(g_m.IEnemyShooter)
	if isShooter {
		shotLatency := GAME_LOOP_TICK_DELAY_MS * 64
		if gs.gameDurationMs%shotLatency == 0 {
			bullets := enemyShooter.Shot()
			for _, bullet := range bullets {
				gs.objectsPool = append(gs.objectsPool, bullet)
			}
		}
	}
}

// @FIX handle bullets from enemy
func (gs *AppGameSrv) handleBulletBehaviourOnTick(bullet g_m.IBullet) {
	delay := bullet.MovementDelay(GAME_LOOP_TICK_DELAY_MS)
	if gs.gameDurationMs%delay == 0 {
		if bullet.Owner() == "Player" {
			bullet.Move(g_m.MoveTopX0_Y1())
		} else {
			bullet.Move(g_m.MoveBottomX0_Y1())
		}
	}
}

func (gs *AppGameSrv) increaseScore() {
	gs.score++
}

func (gs *AppGameSrv) resetScore() {
	gs.score = 0
}

func (gs *AppGameSrv) getArenaSize() (x, y int) {
	windowSize := gs.injector.Storage().WindowSize()
	// first row leave for header with HP and score
	return windowSize.Width, windowSize.Height - 1
}
