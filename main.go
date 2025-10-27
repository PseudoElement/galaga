package main

import (
	"fmt"
	"os"

	app "github.com/pseudoelement/galaga/src"
)

func main() {
	p := app.NewApp()
	if _, err := p.Run(); err != nil {
		fmt.Printf("Start error: %v", err)
		os.Exit(1)
	}
}
