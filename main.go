package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	app "github.com/pseudoelement/galaga/src"
)

func main() {
	p := tea.NewProgram(app.NewApp())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
