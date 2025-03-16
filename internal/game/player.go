package game

// Player represent a player in the game
type Player struct {
	X     int
	Y     int
	Width int
}

// Shoot does the player shoot
func (p *Player) Shoot(g *Game) {
	g.Bullets = append(g.Bullets, Bullet{
		X:          float64(p.X),
		Y:          float64(p.Y - 1),
		FromPlayer: true,
		Velocity:   -bulletSpeed,
	})
}
