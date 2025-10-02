package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pseudoelement/galaga/src/menu/pages"
)

type App struct {
	view AppView
}

func NewApp() *App {
	return &App{view: NewAppView()}
}

func (app *App) Init() tea.Cmd {
	app.view.SetPage(pages.NewPageFirst())
	return nil
}

func (app *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return app, tea.Quit
		case "up":
			app.view.CurrentPage().SelectPrev()
			return app, nil
		case "down":
			app.view.CurrentPage().SelectNext()
			return app, nil
		}

	case tea.WindowSizeMsg:
		// width := msg.Width
		// height := msg.Height
	}

	return app, nil
}

func (app *App) View() string {
	return app.view.View()
}
