package game

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styles for the game
var (
	playerStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
	enemyStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	bulletStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00"))
	scoreStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	gameOverStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true)
	difficultyStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF"))
	farewellStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFAA00")).Bold(true)
)

// Model is the UI model for BubbleTea
type Model struct {
	game         Game
	keyPressed   string
	showFarewell bool
	finalScore   int
}

// NewModel create new instance of the model
func NewModel(difficulty string) Model {
	return Model{
		game:         NewGame(difficulty),
		showFarewell: false,
		finalScore:   0,
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return tea.Tick(time.Second/15, func(t time.Time) tea.Msg {
		return TickMsg{}
	})
}

// TickMsg is a message that tells the program to update the game
type TickMsg struct{}

// Update updates the game state
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			if m.showFarewell {
				return m, tea.Quit
			}
			// Activar pantalla de despedida
			m.showFarewell = true
			m.finalScore = m.game.Score
			// Mostrar despedida por 2 segundos antes de salir
			return m, tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
				return QuitMsg{}
			})
		case "left", "a":
			if !m.showFarewell && !m.game.GameOver {
				if m.game.Player.X > 2 {
					m.game.Player.X -= playerSpeed
				}
			}
		case "right", "d":
			if !m.showFarewell && !m.game.GameOver {
				if m.game.Player.X < m.game.Width-3 {
					m.game.Player.X += playerSpeed
				}
			}
		case " ", "up", "w":
			if !m.showFarewell && !m.game.GameOver {
				m.game.Player.Shoot(&m.game)
			}
		case "r":
			if !m.showFarewell && m.game.GameOver {
				m.game = NewGame(m.game.Difficulty)
			}
		}
	case TickMsg:
		if !m.showFarewell {
			m.game.Update()
			return m, tea.Tick(time.Second/15, func(t time.Time) tea.Msg {
				return TickMsg{}
			})
		}
	case QuitMsg:
		return m, tea.Quit
	}

	return m, nil
}

// QuitMsg outputs a message to quit the game
type QuitMsg struct{}

// View renders the game
func (m Model) View() string {
	if m.showFarewell {
		return renderFarewell(m.finalScore)
	}

	if m.game.GameOver {
		return renderGameOver(m.game)
	}

	// Create a screen buffer
	screen := make([][]string, m.game.Height)
	for i := range screen {
		screen[i] = make([]string, m.game.Width)
		for j := range screen[i] {
			screen[i][j] = " "
		}
	}

	// Render the player
	playerX := m.game.Player.X
	playerY := m.game.Player.Y
	screen[playerY][playerX-1] = "/"
	screen[playerY][playerX] = playerChar
	screen[playerY][playerX+1] = "\\"

	// Render enemies
	for _, enemy := range m.game.Enemies {
		x, y := int(enemy.X), int(enemy.Y)
		if x >= 0 && x < m.game.Width && y >= 0 && y < m.game.Height {
			screen[y][x] = enemyChar
		}
	}

	// Render bullets
	for _, bullet := range m.game.Bullets {
		x, y := int(bullet.X), int(bullet.Y)
		if x >= 0 && x < m.game.Width && y >= 0 && y < m.game.Height {
			screen[y][x] = bulletChar
		}
	}

	// Build the screen
	var sb strings.Builder

	// Status bar
	difficultyText := fmt.Sprintf("Dificultad: %s", getDifficultyName(m.game.Difficulty))
	statusBar := fmt.Sprintf("Puntuación: %d | Vidas: %d | %s",
		m.game.Score, m.game.Lives, difficultyStyle.Render(difficultyText))
	sb.WriteString(scoreStyle.Render(statusBar))
	sb.WriteString("\n")

	// Top frame
	sb.WriteString("+" + strings.Repeat("-", m.game.Width-2) + "+\n")

	// Game content
	for y := 0; y < m.game.Height-2; y++ {
		sb.WriteString("|")
		for x := 0; x < m.game.Width-2; x++ {
			char := screen[y+1][x+1]
			switch char {
			case playerChar, "/", "\\":
				sb.WriteString(playerStyle.Render(char))
			case enemyChar:
				sb.WriteString(enemyStyle.Render(char))
			case bulletChar:
				sb.WriteString(bulletStyle.Render(char))
			default:
				sb.WriteString(char)
			}
		}
		sb.WriteString("|\n")
	}

	// Bottom frame
	sb.WriteString("+" + strings.Repeat("-", m.game.Width-2) + "+\n")

	// Controls
	sb.WriteString("Controles: [←/→ o A/D] Mover, [Espacio o W] Disparar, [Q] Salir")

	return sb.String()
}

// getDifficultyName returns the name of the difficulty
/*
 @param difficulty string
 @return string
*/
func getDifficultyName(difficulty string) string {
	switch difficulty {
	case "easy":
		return "Fácil"
	case "hard":
		return "Difícil"
	default:
		return "Normal"
	}
}

// renderFarewell show bye message
func renderFarewell(score int) string {
	var sb strings.Builder

	// Write the farewell message
	width := 50

	sb.WriteString("\n\n")
	sb.WriteString(strings.Repeat(" ", (screenWidth-width)/2))
	sb.WriteString("+" + strings.Repeat("=", width-2) + "+\n")

	// White line
	sb.WriteString(strings.Repeat(" ", (screenWidth-width)/2))
	sb.WriteString("|" + strings.Repeat(" ", width-2) + "|\n")

	// Thank you message
	message := "¡GRACIAS POR JUGAR!"
	padding := (width - 2 - len(message)) / 2
	sb.WriteString(strings.Repeat(" ", (screenWidth-width)/2))
	sb.WriteString("|" + strings.Repeat(" ", padding))
	sb.WriteString(farewellStyle.Render(message))
	sb.WriteString(strings.Repeat(" ", width-2-padding-len(message)) + "|\n")

	// White line
	sb.WriteString(strings.Repeat(" ", (screenWidth-width)/2))
	sb.WriteString("|" + strings.Repeat(" ", width-2) + "|\n")

	// Score message
	scoreMsg := fmt.Sprintf("Tu puntuación final: %d", score)
	scorePadding := (width - 2 - len(scoreMsg)) / 2
	sb.WriteString(strings.Repeat(" ", (screenWidth-width)/2))
	sb.WriteString("|" + strings.Repeat(" ", scorePadding))
	sb.WriteString(scoreMsg)
	sb.WriteString(strings.Repeat(" ", width-2-scorePadding-len(scoreMsg)) + "|\n")

	// White line
	sb.WriteString(strings.Repeat(" ", (screenWidth-width)/2))
	sb.WriteString("|" + strings.Repeat(" ", width-2) + "|\n")

	// Closing line
	sb.WriteString(strings.Repeat(" ", (screenWidth-width)/2))
	sb.WriteString("+" + strings.Repeat("=", width-2) + "+\n")

	sb.WriteString("\n\nSaliendo...")

	return sb.String()
}

// renderGameOver shows the game over screen
func renderGameOver(g Game) string {
	var sb strings.Builder

	sb.WriteString("\n\n")
	sb.WriteString(gameOverStyle.Render("  GAME OVER  "))
	sb.WriteString("\n\n")
	sb.WriteString(fmt.Sprintf("  Puntuación final: %d\n", g.Score))
	sb.WriteString(fmt.Sprintf("  Dificultad: %s\n\n", getDifficultyName(g.Difficulty)))
	sb.WriteString("  Presiona 'R' para reiniciar o 'Q' para salir\n")

	return sb.String()
}
