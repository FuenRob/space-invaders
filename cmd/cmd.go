package cmd

import (
	"github.com/spf13/cobra"
	"space-invaders/internal/game"
)

// Flags to choose the difficulty
var (
	easyMode bool
	hardMode bool
)

// The root command of the program
var rootCmd = &cobra.Command{
	Use:   "space-invaders",
	Short: "Un juego de Space Invaders en la terminal",
	Long: `Un juego de Space Invaders implementado como CLI en Go
usando Cobra para comandos y BubbleTea para la interfaz de usuario.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Determinar la dificultad
		var difficulty string
		if easyMode {
			difficulty = "easy"
		} else if hardMode {
			difficulty = "hard"
		} else {
			difficulty = "normal" // By default
		}
		game.Start(difficulty)
	},
}

// Execute is the main command of the program
func Execute() error {
	return rootCmd.Execute()
}

// Initializes the flags of the root command
func init() {
	// Flags to choose the difficulty
	rootCmd.Flags().BoolVarP(&easyMode, "easy", "e", false, "Modo fácil")
	rootCmd.Flags().BoolVarP(&hardMode, "hard", "d", false, "Modo difícil")
}
