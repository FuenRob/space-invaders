package game

// Enemy represent a enemy in the game
type Enemy struct {
	X float64
	Y float64
}

// Shoot does the enemy shoot
func (e *Enemy) Shoot(g *Game) {
	g.Bullets = append(g.Bullets, Bullet{
		X:          e.X,
		Y:          e.Y + 1,
		FromPlayer: false,
		Velocity:   bulletSpeed,
	})
}
