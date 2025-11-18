package game_srv

import (
	"fmt"
	"runtime"
	"time"

	g_o "github.com/pseudoelement/galaga/src/game/game-objects"
	game_models "github.com/pseudoelement/galaga/src/game/models"
)

func cellParams(x, y int16, color string) game_models.CellConstructorParams {
	return game_models.CellConstructorParams{
		Color:  color,
		Coords: game_models.Coords{X: x, Y: y},
	}
}

func cellsToView(cells []game_models.ICell) []string {
	view := make([]string, len(cells))
	for idx, cell := range cells {
		view[idx] = cell.Style().Render()
	}
	return view
}

func isObjectOutOfArena(obj game_models.IGameObject, arenaWidth, arenaHeight int) bool {
	for _, cell := range obj.Cells() {
		coords := cell.Coords()
		if coords.X < 0 ||
			coords.Y < 0 ||
			coords.X > int16(arenaWidth) ||
			coords.Y > int16(arenaHeight) {
			return true
		}
	}
	return false
}

func isCellOutOfArena(cell game_models.ICell, arenaWidth, arenaHeight int) bool {
	coords := cell.Coords()
	if coords.X < 0 ||
		coords.Y < 0 ||
		coords.X > int16(arenaWidth-1) ||
		coords.Y > int16(arenaHeight-1) {
		return true
	}
	return false
}

func throwOnError(err error, tag string) {
	if err != nil {
		message := fmt.Sprintf("[%s] error ==> %s", tag, err.Error())
		panic(message)
	}
}

func runBackgroundProcesses(gameSrv *AppGameSrv) {
	var m runtime.MemStats
	for !gameSrv.stop {
		time.Sleep(200 * time.Millisecond)
		runtime.ReadMemStats(&m)
		usedMemoryKB := m.Alloc / 1024
		gameSrv.memoryUsed = usedMemoryKB

		// _clearUnusedCells(gameSrv)
	}
}

func _clearUnusedCells(gameSrv *AppGameSrv) {
	// activePlayerCells := make(map[game_models.Coords]bool, 300)

	// for _, cell := range gameSrv.player.Cells() {
	// 	activePlayerCells[cell.Coords()] = true
	// }
	for _, arenaRow := range gameSrv.arena {
		for _, arenaCell := range arenaRow {
			if g_o.IsPlayer(arenaCell.Owner()) {
				found := false
				for _, playerCell := range gameSrv.player.Cells() {
					if playerCell.Coords().X == arenaCell.Coords().X && playerCell.Coords().Y == arenaCell.Coords().Y {
						found = true
						break
					}
				}

				if !found {
					arenaCell.Destroy()
					// coords := arenaCell.Coords()
					// gameSrv.arena[coords.Y][coords.X] = blackCell(arenaCell)
				}
			}
			// if _, foundAsActive := activePlayerCells[arenaCell.Coords()]; !foundAsActive && g_o.IsPlayer(arenaCell.Owner()) {
			// 	// arenaCell.Destroy()
			// 	println("incative_player_cell")
			// 	// coords := arenaCell.Coords()
			// 	// gameSrv.arena[coords.Y][coords.X] = blackCell(arenaCell)
			// }
		}
	}
}

func _clearPlayerDuplicates(gameSrv *AppGameSrv) {
	width := gameSrv.injector.Storage().WindowSize().Width
	height := gameSrv.injector.Storage().WindowSize().Height
	gameObjectsCellsCoords := make(map[game_models.Coords]bool, width*height)

	for _, obj := range gameSrv.objectsPool {
		for _, cell := range obj.Cells() {
			gameObjectsCellsCoords[cell.Coords()] = true
		}
	}

	for _, arenaRow := range gameSrv.arena {
		for _, cell := range arenaRow {
			coords := cell.Coords()
			if _, ok := gameObjectsCellsCoords[coords]; !ok {
				gameSrv.arena[coords.Y][coords.X] = blackCell(cell)
			}
		}
	}
}
