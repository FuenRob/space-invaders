package game

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"math/rand"
)

// Contants for the game
const (
	screenWidth  = 80
	screenHeight = 24
	playerChar   = "W"
	enemyChar    = "M"
	bulletChar   = "|"
)

// Difficult configuration
var (
	// Values by default (Normal difficulty mode)
	playerSpeed      = 1
	enemySpeed       = 0.2
	bulletSpeed      = 1.0
	enemyShootChance = 1.0 // % chance that an enemy will shoot
	playerLives      = 3
)

// Game represents the state of the game
type Game struct {
	Player        Player
	Enemies       []Enemy
	Bullets       []Bullet
	Score         int
	Lives         int
	GameOver      bool
	Width         int
	Height        int
	Tick          int
	LastSpawned   int
	EnemyDirRight bool
	Difficulty    string
}

// NewGame creates a new instance of the game
/*
 @param difficulty string
 @return Game
*/
func NewGame(difficulty string) Game {
	// Config values according to the difficulty
	setDifficulty(difficulty)

	g := Game{
		Player: Player{
			X:     screenWidth / 2,
			Y:     screenHeight - 2,
			Width: 3, // The player is 3 characters wide
		},
		Enemies:       []Enemy{},
		Bullets:       []Bullet{},
		Score:         0,
		Lives:         playerLives,
		GameOver:      false,
		Width:         screenWidth,
		Height:        screenHeight,
		Tick:          0,
		EnemyDirRight: true,
		Difficulty:    difficulty,
	}

	// Initial enemies
	g.SpawnEnemies(getDifficultyRows(difficulty), 8) // rows according to difficulty, 8 columns

	return g
}

// setDifficulty sets the game parameters according to the difficulty
/*
 @param difficulty string
*/
func setDifficulty(difficulty string) {
	switch difficulty {
	case "easy":
		playerSpeed = 2
		enemySpeed = 0.1
		bulletSpeed = 1
		enemyShootChance = 0.5
		playerLives = 5
	case "hard":
		playerSpeed = 1
		enemySpeed = 0.3
		bulletSpeed = 1.5
		enemyShootChance = 2
		playerLives = 2
	default: // normal
		playerSpeed = 1
		enemySpeed = 0.2
		bulletSpeed = 1
		enemyShootChance = 1
		playerLives = 3
	}
}

// getDifficultyRows returns the number of rows of enemies according to the difficulty
/*
 @param difficulty string
 @return int
*/
func getDifficultyRows(difficulty string) int {
	switch difficulty {
	case "easy":
		return 2
	case "hard":
		return 4
	default: // normal
		return 3
	}
}

// SpawnEnemies generates a group of enemies
/*
 @param rows int
 @param cols int
*/
func (g *Game) SpawnEnemies(rows, cols int) {
	spacingX := 6
	spacingY := 2

	startX := (g.Width - (cols * spacingX)) / 2
	startY := 3

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			x := startX + (c * spacingX)
			y := startY + (r * spacingY)
			g.Enemies = append(g.Enemies, Enemy{
				X: float64(x),
				Y: float64(y),
			})
		}
	}
}

// Update updates the game state
func (g *Game) Update() {
	if g.GameOver {
		return
	}

	g.Tick++

	// Move enemies
	moveX := enemySpeed
	if !g.EnemyDirRight {
		moveX = -enemySpeed
	}

	needReverse := false
	for i := range g.Enemies {
		g.Enemies[i].X += moveX

		// Verify if enemies are at the edge
		if g.Enemies[i].X <= 0 || g.Enemies[i].X >= float64(g.Width-2) {
			needReverse = true
		}

		// Random chance that an enemy will shoot
		if rand.Float64() < enemyShootChance/100 { // probabilidad según dificultad
			g.Enemies[i].Shoot(g)
		}
	}

	// Change direction if needed and move down
	if needReverse {
		g.EnemyDirRight = !g.EnemyDirRight
		for i := range g.Enemies {
			g.Enemies[i].Y += 1

			// Game over si los enemigos llegan abajo
			if g.Enemies[i].Y >= float64(g.Player.Y) {
				g.GameOver = true
			}
		}
	}

	// Update bullets
	for i := 0; i < len(g.Bullets); i++ {
		g.Bullets[i].Update()

		// Delete bullets out of bounds
		if g.Bullets[i].Y < 0 || g.Bullets[i].Y >= float64(g.Height) {
			g.Bullets = append(g.Bullets[:i], g.Bullets[i+1:]...)
			i--
			continue
		}

		// Detect collisions
		if g.Bullets[i].FromPlayer {
			// Bullets from player against enemies
			for j := 0; j < len(g.Enemies); j++ {
				if g.Bullets[i].X >= g.Enemies[j].X-1 &&
					g.Bullets[i].X <= g.Enemies[j].X+1 &&
					g.Bullets[i].Y <= g.Enemies[j].Y+1 &&
					g.Bullets[i].Y >= g.Enemies[j].Y-1 {

					// Destroy enemy
					g.Score += 10
					g.Enemies = append(g.Enemies[:j], g.Enemies[j+1:]...)
					g.Bullets = append(g.Bullets[:i], g.Bullets[i+1:]...)
					i--
					break
				}
			}
		} else {
			// Bullets from enemies against player
			if g.Bullets[i].X >= float64(g.Player.X-g.Player.Width/2) &&
				g.Bullets[i].X <= float64(g.Player.X+g.Player.Width/2) &&
				g.Bullets[i].Y >= float64(g.Player.Y-1) &&
				g.Bullets[i].Y <= float64(g.Player.Y+1) {

				g.Lives--
				g.Bullets = append(g.Bullets[:i], g.Bullets[i+1:]...)
				i--

				if g.Lives <= 0 {
					g.GameOver = true
				}
			}
		}
	}

	// Generate new enemies
	if len(g.Enemies) == 0 {
		g.SpawnEnemies(getDifficultyRows(g.Difficulty), 8)
	}
}

// Start init the game
/*
 @param difficulty string
*/
func Start(difficulty string) {
	p := tea.NewProgram(NewModel(difficulty))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error al iniciar el juego: %v\n", err)
	}
}
