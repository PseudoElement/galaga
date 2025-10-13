package game_srv

import (
	"time"

	"github.com/charmbracelet/lipgloss"
	game_models "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
	game_srv_styles "github.com/pseudoelement/galaga/src/services/game-service/styles"
)

type AppGameSrv struct {
	injector models.IAppInjector

	arena       [][]string
	objectsPool []game_models.IGameObject
	player      game_models.IPlayer
	score       int32
	stop        bool
}

func NewAppGameSrv(injector models.IAppInjector) models.IAppGameSrv {
	return &AppGameSrv{
		injector:    injector,
		arena:       make([][]string, 0),
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
		row := lipgloss.JoinHorizontal(lipgloss.Left, arenaRow...)
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
	for heightCount := range rowsCount {
		arenaRow := make([]string, 0, width)
		for widthCount := range width {
			if (heightCount+widthCount)%2 == 0 {
				arenaRow = append(arenaRow, game_srv_styles.DarkCell.Render())
			} else {
				arenaRow = append(arenaRow, game_srv_styles.SemiDarkCell.Render())
			}
		}

		gs.arena = append(gs.arena, arenaRow)
	}

	for _, cell := range gs.player.Cells() {
		coords := cell.Coords()
		playerCell := game_srv_styles.Cell.Background(lipgloss.Color(cell.Color()))
		gs.arena[coords.Y][coords.X] = playerCell.Render()
	}

	go gs.runLoop()
}

func (gs *AppGameSrv) EndGame() {
	gs.player = nil
	gs.objectsPool = make([]game_models.IGameObject, 0)
	gs.arena = make([][]string, 0)
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

		// @FIX use prevCells instead of redrawing whole arena
		for rowIdx, arenaRow := range gs.arena {
			for columnIdx := range arenaRow {
				if (rowIdx+columnIdx)%2 == 0 {
					arenaRow[columnIdx] = game_srv_styles.DarkCell.Render()
				} else {
					arenaRow[columnIdx] = game_srv_styles.SemiDarkCell.Render()
				}
			}
		}

		for _, cell := range gs.player.Cells() {
			coords := cell.Coords()
			playerCell := game_srv_styles.Cell.Background(lipgloss.Color(cell.Color()))
			gs.arena[coords.Y][coords.X] = playerCell.Render()
		}
	}
}

func (gs *AppGameSrv) increaseScore() {
	gs.score++
}

func (gs *AppGameSrv) resetScore() {
	gs.score = 0
}
