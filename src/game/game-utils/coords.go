package game_utils

import (
	"math"

	game_models "github.com/pseudoelement/galaga/src/game/models"
)

func GetObjectMidCoords(obj game_models.IGameObject) game_models.Coords {
	var xMin, xMax int16 = math.MaxInt16, 0
	var yMin, yMax int16 = math.MaxInt16, 0
	for _, cell := range obj.Cells() {
		xCoord := cell.Coords().X
		yCoord := cell.Coords().Y

		if xCoord > xMax {
			xMax = xCoord
		}
		if xCoord < xMin {
			xMin = xCoord
		}
		if yCoord > yMax {
			yMax = yCoord
		}
		if yCoord < yMin {
			yMin = yCoord
		}
	}

	midX := math.Round(float64(xMin+xMax) / 2)
	midY := math.Round(float64(yMin+yMax) / 2)

	return game_models.Coords{X: int16(midX), Y: int16(midY)}
}
