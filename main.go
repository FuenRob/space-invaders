package main

import (
	"fmt"
	"os"

	"space-invaders/cmd"
)

/**
 * Main function of the program that use
 * the Execute function of the CMD package
 */
func main() {
	if err := cmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
