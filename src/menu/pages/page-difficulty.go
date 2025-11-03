package pages

import (
	"fmt"

	consts "github.com/pseudoelement/galaga/src/constants"
	pages_styles "github.com/pseudoelement/galaga/src/menu/pages/styles"
	"github.com/pseudoelement/galaga/src/menu/ui"
	"github.com/pseudoelement/galaga/src/models"
)

type PageDifficultyLvl struct {
	*Page
}

func NewPageDifficultyLvl(injector models.IAppInjector) models.IPage {
	p := &PageDifficultyLvl{Page: NewPage(injector)}
	currDifficulty := consts.DIFFICULTY_TO_TRANSLATE_MAP[injector.Storage().GameDifficulty()]

	// init list
	p.elementsList.PushBack(ui.NewBaseViewElement(models.BaseViewElementParams{
		InitialStyle: pages_styles.BoldTextStyle,
		Text:         injector.LanguageSrv().Translate("menu.difficulty.selectDifficulty"),
	}))
	p.elementsList.PushBack(ui.NewBaseViewElement(models.BaseViewElementParams{
		InitialStyle: pages_styles.AccentTextStyle,
		Text: fmt.Sprintf(
			"%v: %v\n",
			injector.LanguageSrv().Translate("menu.difficulty.currDifficulty"),
			injector.LanguageSrv().Translate(fmt.Sprintf("menu.difficulty.%s", currDifficulty)),
		),
	}))

	p.elementsList.PushBack(NewEasyDifficultyButton(p.injector))
	p.elementsList.PushBack(NewMediumDifficultyButton(p.injector))
	p.elementsList.PushBack(NewHardDifficultyButton(p.injector))

	p.elementsList.PushBack(NewMenuRedirectButton(
		injector.LanguageSrv().Translate("menu.buttons.back"),
		func() models.IPage { return NewPageFirst(injector) },
	),
	)
	p.SelectFirstSelectable()

	return p
}
