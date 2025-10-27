package game_srv

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	g_c "github.com/pseudoelement/galaga/src/game/game-constants"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
	game_srv_styles "github.com/pseudoelement/galaga/src/services/game-service/styles"
)

const (
	GAME_LOOP_TICK_DELAY_MS = 20
)

type AppGameSrv struct {
	injector models.IAppInjector

	arena          [][]g_m.ICell
	gameDurationMs int
	objectsPool    []g_m.IGameObject
	player         g_m.IPlayer
	score          int32
	stop           bool
}

func NewAppGameSrv(injector models.IAppInjector, player g_m.IPlayer) models.IAppGameSrv {
	return &AppGameSrv{
		injector: injector,

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

func (gs *AppGameSrv) IsPlaying() bool {
	return gs.player != nil && !gs.stop
}

func (gs *AppGameSrv) View() string {
	// align rows horizontally
	rows := make([]string, 0)
	for _, arenaRow := range gs.arena {
		row := lipgloss.JoinHorizontal(lipgloss.Left, cellsToView(arenaRow)...)
		rows = append(rows, row)
	}

	arenaView := lipgloss.JoinVertical(lipgloss.Top, rows...)

	if !gs.IsPlaying() {
		return ""
	}
	hpView := fmt.Sprintf("❤️ %v", gs.player.Health())
	scoreView := fmt.Sprintf(" Score: %d", gs.score)

	header := lipgloss.JoinHorizontal(lipgloss.Center, hpView, scoreView)
	view := lipgloss.JoinVertical(lipgloss.Center, header, arenaView)

	return view
}

func (gs *AppGameSrv) StartGame() {
	gs.stop = false
	// at this point gs.injector.Storage().WindowSize() should be already defined
	width, height := gs.getArenaSize()
	for heightCount := range height { // y
		arenaRow := make([]g_m.ICell, 0, width)
		for widthCount := range width { // x
			arenaRow = append(arenaRow, g_m.NewCell(cellParams(
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

	gs.injector.TeaProgram().Send(models.EndGameTrigger{})
}

func (gs *AppGameSrv) SetPlayer(player g_m.IPlayer) {
	if gs.player != nil {
		var prevPlayerIdx int
		for idx, obj := range gs.objectsPool {
			oldPlayer, isPlayer := obj.(g_m.IPlayer)
			if isPlayer {
				prevPlayerIdx = idx
				oldPlayer.Destroy()
			}
		}
		gs.objectsPool = append(gs.objectsPool[0:prevPlayerIdx], gs.objectsPool[prevPlayerIdx+1:]...)
	}

	gs.player = player
	gs.objectsPool = append(gs.objectsPool, player)
}

func (gs *AppGameSrv) runLoop() {
	x, y := gs.getArenaSize()
	for !gs.stop && gs.player != nil {
		if gs.player.Health() <= 0 {
			gs.EndGame()
		}

		time.Sleep(GAME_LOOP_TICK_DELAY_MS * time.Millisecond)
		gs.injector.TeaProgram().Send(models.UpdateTrigger{})
		gs.gameDurationMs += GAME_LOOP_TICK_DELAY_MS

		gs.clearPrevCellsOfObjectsOnTick()
		gs.drawNewCellsOfObjectsOnTick()

		gs.handleEnemySpawn()
		gs.handleBoostSpawn()
		gs.handleShot()
		gs.handleCollision()

		for _, object := range gs.objectsPool {
			switch obj := object.(type) {
			case g_m.IBullet:
				gs.handleBulletBehaviourOnTick(obj)
			case g_m.IEnemy:
				gs.handleEnemyBehaviourOnTick(obj)
			case g_m.IBoost:
				gs.handleBoostBehaviourOnTick(obj)
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
		for _, cell := range obj.PrevCells() {
			coords := cell.Coords()
			if !isCellOutOfArena(cell, width, height) {
				gs.arena[coords.Y][coords.X] = g_m.NewCell(cellParams(coords.X, coords.Y, game_srv_styles.BlackColor))
			}
		}
	}

	for _, obj := range gs.objectsPool {
		if obj.Destroyed() {
			for _, cell := range obj.Cells() {
				coords := cell.Coords()
				if !isCellOutOfArena(cell, width, height) {
					gs.arena[coords.Y][coords.X] = g_m.NewCell(cellParams(coords.X, coords.Y, game_srv_styles.BlackColor))
				}
			}
		} else {
			updatedObjectsPool = append(updatedObjectsPool, obj)
		}
	}

	gs.objectsPool = updatedObjectsPool
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
}

func (gs *AppGameSrv) handleCollision() {
	objectPoolCellsMap := make(map[g_m.Coords]g_m.IGameObject, len(gs.objectsPool)+1)
	for _, currentObj := range gs.objectsPool {
		for _, cell := range currentObj.Cells() {
			crossedObject, crossed := objectPoolCellsMap[cell.Coords()]
			// means 2 objects collided each other in specific cell
			if crossed {
				handleCollisionScenarios(currentObj, crossedObject)
			} else {
				objectPoolCellsMap[cell.Coords()] = currentObj
			}
		}
	}
}

func (gs *AppGameSrv) handleEnemySpawn() {
	spawnLatency := GAME_LOOP_TICK_DELAY_MS * 150
	if gs.gameDurationMs%spawnLatency == 0 {
		spawnedEnemy := gs.injector.Factories().EnemyFactory(g_c.EASY)
		gs.objectsPool = append(gs.objectsPool, spawnedEnemy)

		gs.increaseScore()
	}
}

func (gs *AppGameSrv) handleBoostSpawn() {
	spawnLatency := GAME_LOOP_TICK_DELAY_MS * 150
	if gs.gameDurationMs%spawnLatency == 0 {
		spawnedEnemy := gs.injector.Factories().BoostFactory(g_c.EASY)
		gs.objectsPool = append(gs.objectsPool, spawnedEnemy)
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

func (gs *AppGameSrv) handleBulletBehaviourOnTick(bullet g_m.IBullet) {
	delay := bullet.MovementDelay(GAME_LOOP_TICK_DELAY_MS)
	if gs.gameDurationMs%delay == 0 {
		isBulletOfPlayer := false
		for _, playerName := range g_c.PLAYER_NAMES {
			if playerName == bullet.Owner().Name() {
				isBulletOfPlayer = true
			}
		}
		if isBulletOfPlayer {
			bullet.Move(g_m.MoveTopX0_Y1())
		} else {
			bullet.Move(g_m.MoveBottomX0_Y1())
		}
	}
}

func (gs *AppGameSrv) handleBoostBehaviourOnTick(boost g_m.IBoost) {
	delay := boost.MovementDelay(GAME_LOOP_TICK_DELAY_MS)
	if gs.gameDurationMs%delay == 0 {
		boost.Move(g_m.MoveBottomX0_Y1())
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
