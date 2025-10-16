package app

import (
	tea "github.com/charmbracelet/bubbletea"
	g_m "github.com/pseudoelement/galaga/src/game/models"
	"github.com/pseudoelement/galaga/src/injector"
	"github.com/pseudoelement/galaga/src/menu/pages"
	"github.com/pseudoelement/galaga/src/models"
	game_srv "github.com/pseudoelement/galaga/src/services/game-service"
	language_srv "github.com/pseudoelement/galaga/src/services/language-service"
	"github.com/pseudoelement/galaga/src/storage"
	app_view "github.com/pseudoelement/galaga/src/view"
)

type App struct {
	view     models.IAppView
	gameSrv  models.IAppGameSrv
	injector models.IAppInjector
}

func NewApp() *tea.Program {
	injector := injector.NewAppInjector()

	appStorage := storage.NewAppStorage()
	injector.SetStorage(appStorage)

	appLanguageSrv := language_srv.NewAppLanguageSrv(appStorage)
	injector.SetLanguageSrv(appLanguageSrv)

	appView := app_view.NewAppView(pages.NewPageFirst(injector))
	injector.SetView(appView)

	appView.SetPage(pages.NewPageFirst(injector))

	appGameSrv := game_srv.NewAppGameSrv(injector, nil)
	injector.SetGameSrv(appGameSrv)

	app := &App{view: appView, gameSrv: appGameSrv, injector: injector}

	teaProgram := tea.NewProgram(app)
	injector.SetTeaProgram(teaProgram)

	return teaProgram
}

func (app *App) Init() tea.Cmd {
	return nil
}

func (app *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case models.UpdateTrigger:
		return app, nil

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

		case "w":
			moveDir := g_m.MoveTopX0_Y3()
			if game_srv.CanMoveTop(app.gameSrv.Player(), moveDir) {
				app.gameSrv.Player().Move(moveDir)
			}
			return app, nil
		case "a":
			moveDir := g_m.MoveLeftX3_Y0()
			if game_srv.CanMoveLeft(app.gameSrv.Player(), moveDir) {
				app.gameSrv.Player().Move(moveDir)
			}
			return app, nil
		case "s":
			moveDir := g_m.MoveBottomX0_Y3()
			if game_srv.CanMoveBottom(app.gameSrv.Player(), moveDir, app.injector) {
				app.gameSrv.Player().Move(moveDir)
			}
			return app, nil
		case "d":
			moveDir := g_m.MoveRightX3_Y0()
			if game_srv.CanMoveRight(app.gameSrv.Player(), moveDir, app.injector) {
				app.gameSrv.Player().Move(moveDir)
			}
			return app, nil
		case "esc":
			app.gameSrv.EndGame()
			app.view.SetPage(pages.NewPageFirst(app.injector))
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
