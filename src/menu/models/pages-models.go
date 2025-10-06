package models

type IPage interface {
	View() string
	SelectedItem() ISelectableElement
	Select(element ISelectableElement)
	SelectNext()
	SelectPrev()
}
