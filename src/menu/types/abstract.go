package menu_types

import "github.com/pseudoelement/galaga/src/menu/ui"

type IPage interface {
	View() string
	SelectedItem() *ui.SelectableElement
	Select(element *ui.SelectableElement)
	SelectNext()
	SelectPrev()
}
