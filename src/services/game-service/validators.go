package game_srv

import (
	game_models "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/models"
)

func CanMoveTop(player game_models.IPlayer, moveDir game_models.MoveDir) bool {
	for _, cell := range player.Cells() {
		if cell.Coords().Y+moveDir.Y < 0 {
			return false
		}
	}
	return true
}

func CanMoveBottom(player game_models.IPlayer, moveDir game_models.MoveDir, injector models.IAppInjector) bool {
	bottom := injector.Storage().WindowSize().Height - 1
	for _, cell := range player.Cells() {
		if cell.Coords().Y+moveDir.Y > int16(bottom) {
			return false
		}
	}
	return true
}

func CanMoveLeft(player game_models.IPlayer, moveDir game_models.MoveDir) bool {
	for _, cell := range player.Cells() {
		if cell.Coords().X+moveDir.X < 0 {
			return false
		}
	}
	return true
}

func CanMoveRight(player game_models.IPlayer, moveDir game_models.MoveDir, injector models.IAppInjector) bool {
	right := injector.Storage().WindowSize().Width
	for _, cell := range player.Cells() {
		if cell.Coords().X+moveDir.X >= int16(right) {
			return false
		}
	}
	return true
}
