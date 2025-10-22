package enemy

import (
	g_m "github.com/pseudoelement/galaga/src/game/models"
)

type Enemy struct {
	*g_m.GameObject

	health          int16
	movementPattern []g_m.MoveDir
	movementIdx     int
}

func NewEnemy(health int16, cells []g_m.ICell, movementPattern []g_m.MoveDir) *Enemy {
	return &Enemy{
		GameObject:      g_m.NewGameObject(cells),
		health:          health,
		movementPattern: movementPattern,
	}
}

// movement direction on current game loop tick
func (e *Enemy) MoveDir(arenaWidth, arenaHeight int) g_m.MoveDir {
	if e.movementIdx >= len(e.movementPattern) {
		e.movementIdx = 0
	}

	moveDir := e.movementPattern[e.movementIdx]
	e.movementIdx++

	// check next moveDir if wants to move right being on right edge
	if e.staysOnArenaLeftEdge() && moveDir.X < 0 {
		return e.MoveDir(arenaWidth, arenaHeight)
	}
	// check next moveDir if wants to move left being on left edge
	if e.staysOnArenaRightEdge(arenaWidth) && moveDir.X > 0 {
		return e.MoveDir(arenaWidth, arenaHeight)
	}

	return moveDir
}

func (e *Enemy) Health() int16 {
	return e.health
}

func (e *Enemy) GetHeal(amount int16) {
	if amount < 0 {
		panic("[Enemy_GetHeal] amount should be positive.")
	}
	e.health += amount
}

func (e *Enemy) GetDamage(amount int16) {
	if amount < 0 {
		panic("[Enemy_GetDamage] amount should be positive.")
	}
	e.health -= amount

	if e.health <= 0 {
		e.Destroy()
	}
}

func (e *Enemy) staysOnArenaLeftEdge() bool {
	for _, cell := range e.Cells() {
		if cell.Coords().X == 0 {
			return true
		}
	}
	return false
}

func (e *Enemy) staysOnArenaRightEdge(arenaWidth int) bool {
	for _, cell := range e.Cells() {
		if cell.Coords().X == int16(arenaWidth-1) {
			return true
		}
	}
	return false
}

var _ g_m.IHealable = (*Enemy)(nil)
var _ g_m.IDamageable = (*Enemy)(nil)
