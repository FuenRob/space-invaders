package game

// Bullet represent a bullet in the game
type Bullet struct {
	X          float64
	Y          float64
	FromPlayer bool
	Velocity   float64
}

// Update updates the bullet position
func (b *Bullet) Update() {
	b.Y += b.Velocity
}
