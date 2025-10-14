package game_srv

import (
	"time"

	"github.com/charmbracelet/lipgloss"
	game_objects "github.com/pseudoelement/galaga/src/game/game-objects"
	game_models "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
	game_srv_styles "github.com/pseudoelement/galaga/src/services/game-service/styles"
)

type AppGameSrv struct {
	injector models.IAppInjector

	arena       [][]game_models.ICell
	objectsPool []game_models.IGameObject
	player      game_models.IPlayer
	score       int32
	stop        bool
}

func NewAppGameSrv(injector models.IAppInjector) models.IAppGameSrv {
	return &AppGameSrv{
		injector:    injector,
		arena:       make([][]game_models.ICell, 0),
		objectsPool: make([]game_models.IGameObject, 0),
		score:       0,
		player:      nil,
		stop:        false,
	}
}

func (gs *AppGameSrv) Player() game_models.IPlayer {
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
	header := lipgloss.JoinHorizontal(lipgloss.Center, "❤️ 1", " Score: 1")
	view := lipgloss.JoinVertical(lipgloss.Center, header, arenaView)

	return view
}

func (gs *AppGameSrv) StartGame() {
	// at this point gs.injector.Storage().WindowSize() should be already defined
	width := gs.injector.Storage().WindowSize().Width
	height := gs.injector.Storage().WindowSize().Height
	// first row leave for header with HP and score
	rowsCount := height - 1
	for heightCount := range rowsCount { // y
		arenaRow := make([]game_models.ICell, 0, width)
		for widthCount := range width { // x
			if (heightCount+widthCount)%2 == 0 {
				arenaRow = append(arenaRow, game_objects.NewCell(cellParams(
					int16(widthCount),
					int16(heightCount),
					game_srv_styles.DarkBgColor,
				)))
			} else {
				arenaRow = append(arenaRow, game_objects.NewCell(cellParams(
					int16(widthCount),
					int16(heightCount),
					game_srv_styles.SemiDarkBgColor,
				)))
			}
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
	gs.player = nil
	gs.objectsPool = make([]game_models.IGameObject, 0)
	gs.arena = make([][]game_models.ICell, 0)
	gs.stop = true
	gs.resetScore()
}

func (gs *AppGameSrv) SpawnPlayer(player game_models.IPlayer) {
	gs.player = player
}

func (gs *AppGameSrv) runLoop() {
	for !gs.stop {
		time.Sleep(16 * time.Millisecond)
		gs.injector.TeaProgram().Send(models.UpdateTrigger{})

		for _, cell := range gs.player.PrevCells() {
			coords := cell.Coords()
			if (coords.X+coords.Y)%2 == 0 {
				gs.arena[coords.Y][coords.X] = game_objects.NewCell(cellParams(coords.X, coords.Y, game_srv_styles.DarkBgColor))
			} else {
				gs.arena[coords.Y][coords.X] = game_objects.NewCell(cellParams(coords.X, coords.Y, game_srv_styles.SemiDarkBgColor))
			}
		}

		for _, cell := range gs.player.Cells() {
			coords := cell.Coords()
			gs.arena[coords.Y][coords.X] = cell
		}
	}
}

func (gs *AppGameSrv) increaseScore() {
	gs.score++
}

func (gs *AppGameSrv) resetScore() {
	gs.score = 0
}
