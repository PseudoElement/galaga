package game_srv

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	consts "github.com/pseudoelement/galaga/src/constants"
	g_o "github.com/pseudoelement/galaga/src/game/game-objects"
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
	bestScore      int32
	stop           bool
	memoryUsed     uint64
}

func NewAppGameSrv(injector models.IAppInjector, player g_m.IPlayer) models.IAppGameSrv {
	gameSrv := &AppGameSrv{
		injector:       injector,
		arena:          make([][]g_m.ICell, 0),
		objectsPool:    make([]g_m.IGameObject, 0),
		score:          0,
		player:         player,
		stop:           false,
		gameDurationMs: 0,
		memoryUsed:     0,
	}

	go gameSrv.loadBestScore()

	return gameSrv
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

	memoryInfo := fmt.Sprintf("%s: %v KB | ", gs.injector.LanguageSrv().Translate("game.memoryUsage"), gs.memoryUsed)

	hpView := fmt.Sprintf("❤️ %v", gs.player.Health())
	scoreView := fmt.Sprintf(" %s: %d", gs.injector.LanguageSrv().Translate("game.score"), gs.score)
	besrScoreView := fmt.Sprintf(" %s: %d", gs.injector.LanguageSrv().Translate("game.bestScore"), gs.bestScore)

	gameInfo := lipgloss.JoinHorizontal(lipgloss.Center, hpView, scoreView, besrScoreView)
	header := lipgloss.JoinHorizontal(lipgloss.Left, memoryInfo, gameInfo)

	view := lipgloss.JoinVertical(lipgloss.Center, header, arenaView)

	return view
}

func (gs *AppGameSrv) StartGame() {
	gs.stop = false
	// at this point gs.injector.Storage().WindowSize() should be already defined
	width, height := gs.ArenaSize()
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
	go checkMemoryUsage(gs)
}

func (gs *AppGameSrv) EndGame() {
	go gs.injector.TeaProgram().Send(models.EndGameTrigger{})
	go gs.saveBestScoreIfRecord(gs.score)

	gs.stop = true
	gs.player = nil
	gs.objectsPool = make([]g_m.IGameObject, 0)
	gs.arena = make([][]g_m.ICell, 0)
	gs.gameDurationMs = 0

	gs.resetScore()
}

func (gs *AppGameSrv) SetPlayer(player g_m.IPlayer) {
	if gs.player != nil {
		for _, obj := range gs.objectsPool {
			if g_o.IsPlayer(obj) {
				obj.Destroy()
			}
		}
	}

	gs.player = player
	gs.objectsPool = append(gs.objectsPool, player)
}

func (gs *AppGameSrv) AddObjectsToObjectPool(objects ...g_m.IGameObject) {
	for _, obj := range objects {
		gs.objectsPool = append(gs.objectsPool, obj)
	}
}

func (gs *AppGameSrv) runLoop() {
	x, y := gs.ArenaSize()
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
			case g_m.IBlast:
				gs.handleBlastDisappearance(obj)
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
	width, height := gs.ArenaSize()
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
	width, height := gs.ArenaSize()
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
				handleCollisionScenarios(currentObj, crossedObject, gs)
			} else {
				objectPoolCellsMap[cell.Coords()] = currentObj
			}
		}
	}
}

func (gs *AppGameSrv) handleEnemySpawn() {
	spawnLatency := GAME_LOOP_TICK_DELAY_MS * 150
	if gs.gameDurationMs%spawnLatency == 0 {
		spawnedEnemy := gs.injector.Factories().EnemyFactory(consts.EASY)
		gs.objectsPool = append(gs.objectsPool, spawnedEnemy)
	}
}

func (gs *AppGameSrv) handleBoostSpawn() {
	hpSpawnLatency := GAME_LOOP_TICK_DELAY_MS * 1_500
	shipSpawnLatency := GAME_LOOP_TICK_DELAY_MS * 3_000
	if gs.gameDurationMs%shipSpawnLatency == 0 {
		nextTierShipBoost := gs.injector.Factories().BoostFactory(consts.EASY, true)
		gs.objectsPool = append(gs.objectsPool, nextTierShipBoost)
		return
	}
	if gs.gameDurationMs%hpSpawnLatency == 0 {
		hpBoost := gs.injector.Factories().BoostFactory(consts.EASY, false)
		gs.objectsPool = append(gs.objectsPool, hpBoost)
		return
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

func (gs *AppGameSrv) handleBlastDisappearance(blast g_m.IBlast) {
	msNow := time.Now().UnixMilli()
	if msNow >= int64(blast.DestructInMs()) {
		blast.Destroy()
	}
}

func (gs *AppGameSrv) handleEnemyBehaviourOnTick(enemy g_m.IEnemy) {
	width, height := gs.ArenaSize()
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
		isBulletOfPlayer := g_o.IsPlayer(bullet.Owner())
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

func (gs *AppGameSrv) increaseScore(plusPoints int32) {
	gs.score += plusPoints
}

func (gs *AppGameSrv) resetScore() {
	gs.score = 0
}

func (gs *AppGameSrv) loadBestScore() {
	pwd, _ := os.Getwd()
	path := pwd + "/data" + "/best-score.txt"

	fileBytes, err := os.ReadFile(path)
	throwOnError(err, "AppGameSrv_loadBestScore_ReadFile")

	bestScore, err := strconv.Atoi(string(fileBytes))
	throwOnError(err, "AppGameSrv_loadBestScore_Atoi")

	gs.bestScore = int32(bestScore)
}

func (gs *AppGameSrv) saveBestScoreIfRecord(currentScore int32) {
	if gs.bestScore < currentScore {
		gs.bestScore = currentScore
		newBestScore := strconv.Itoa(int(currentScore))

		pwd, _ := os.Getwd()
		path := pwd + "/data" + "/best-score.txt"

		err := os.WriteFile(path, []byte(newBestScore), 0777)
		throwOnError(err, "AppGameSrv_saveBestScoreIfRecord_WriteFile")
	}
}

func (gs *AppGameSrv) ArenaSize() (width, height int) {
	windowSize := gs.injector.Storage().WindowSize()
	// first row leave for header with HP and score
	return windowSize.Width, windowSize.Height - 1
}
