package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pseudoelement/galaga/src/injector"
	"github.com/pseudoelement/galaga/src/menu/models"
	"github.com/pseudoelement/galaga/src/menu/pages"
	"github.com/pseudoelement/galaga/src/storage"
	app_view "github.com/pseudoelement/galaga/src/view"
)

type App struct {
	view     models.IAppView
	injector models.IAppInjector
}

func NewApp() *App {
	appStorage := storage.NewAppStorage()
	appView := app_view.NewAppView(pages.NewPageFirst(nil))
	injector := injector.NewAppInjector(appStorage, appView)
	println("start NewApp")

	appView.SetPage(pages.NewPageFirst(injector))

	return &App{view: appView, injector: injector}
}

func (app *App) Init() tea.Cmd {

	// app.view = appView
	// app.injector = injector.NewAppInjector(appStorage, appView)

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
		case "enter":
			currPage := app.view.CurrentPage()
			selectedEl := currPage.SelectedItem()
			if selectedEl != nil {
				actionEl, withAction := selectedEl.(models.IActionElement)
				if withAction {
					actionEl.Action()
				}
				redirectEl, withRedirect := selectedEl.(models.IRedirectableElement)
				if withRedirect {
					nextPage := redirectEl.PageToRedirect()
					app.view.SetPage(nextPage)
				}
			}

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
