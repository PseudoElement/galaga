package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pseudoelement/galaga/src/injector"
	"github.com/pseudoelement/galaga/src/menu/pages"
	"github.com/pseudoelement/galaga/src/models"
	language_srv "github.com/pseudoelement/galaga/src/services/language-service"
	"github.com/pseudoelement/galaga/src/storage"
	app_view "github.com/pseudoelement/galaga/src/view"
)

type App struct {
	view     models.IAppView
	injector models.IAppInjector
}

func NewApp() *App {
	injector := injector.NewAppInjector()

	appStorage := storage.NewAppStorage()
	injector.SetStorage(appStorage)

	appLanguageSrv := language_srv.NewAppLanguageSrv(appStorage)
	injector.SetLanguageSrv(appLanguageSrv)

	appView := app_view.NewAppView(pages.NewPageFirst(injector))
	injector.SetView(appView)

	appView.SetPage(pages.NewPageFirst(injector))

	return &App{view: appView, injector: injector}
}

func (app *App) Init() tea.Cmd {
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
		windowSize := models.WindowSize{
			Width:  msg.Width,
			Height: msg.Height,
		}
		app.injector.Storage().SetWindowSize(windowSize)
	}

	return app, nil
}

func (app *App) View() string {
	return app.view.View()
}
