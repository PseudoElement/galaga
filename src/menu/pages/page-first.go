package pages

import (
	"container/list"

	"github.com/charmbracelet/lipgloss"
	menu_types "github.com/pseudoelement/galaga/src/menu/types"
	"github.com/pseudoelement/galaga/src/menu/ui"
)

type PageFirst struct {
	elementsList *list.List
}

func NewPageFirst() menu_types.IPage {
	p := &PageFirst{elementsList: list.New()}

	// init list
	p.elementsList.PushBack(ui.NewSelectableElement("Start"))
	p.elementsList.PushBack(ui.NewSelectableElement("Language"))
	p.elementsList.PushBack(ui.NewSelectableElement("Settings"))
	p.elementsList.PushBack(ui.NewSelectableElement("Quit"))

	// select first element by default
	firstEl := p.elementsList.Front().Value.(*ui.SelectableElement)
	firstEl.SetSelected(true)

	return p
}

func (p *PageFirst) View() string {
	elementsSlice := make([]*ui.SelectableElement, 0, 4)
	for el := p.elementsList.Front(); el != nil; el = el.Next() {
		selectableEl, ok := el.Value.(*ui.SelectableElement)
		if ok {
			prevStyle := selectableEl.LipglossElement()
			if selectableEl.IsSelected() {
				selectableEl.UpdateStyle(prevStyle.Foreground(lipgloss.Color("#07c3f7")))
				selectableEl.UpdateText("❯ " + selectableEl.InitialValue())
			} else {
				selectableEl.UpdateStyle(prevStyle.Foreground(lipgloss.Color("#d9d9d9")))
				selectableEl.UpdateText(selectableEl.InitialValue())
			}
		}
		elementsSlice = append(elementsSlice, selectableEl)
	}

	return lipgloss.JoinVertical(lipgloss.Left, p.SelectableElsToStrings(elementsSlice)...)
}

func (p *PageFirst) ElementsList() *list.List {
	return p.elementsList
}

func (p *PageFirst) SelectableElsToStrings(elements []*ui.SelectableElement) []string {
	var sl []string
	for _, el := range elements {
		sl = append(sl, el.String())
	}

	return sl
}

func (p *PageFirst) SelectedItem() *ui.SelectableElement {
	for el := p.elementsList.Front(); el != nil; el = el.Next() {
		selectableEl, ok := el.Value.(*ui.SelectableElement)
		if ok && selectableEl.IsSelected() {
			return selectableEl
		}
	}

	return nil
}

func (p *PageFirst) Select(newSelectedEl *ui.SelectableElement) {
	for el := p.elementsList.Front(); el != nil; el = el.Next() {
		currentEl := el.Value.(*ui.SelectableElement)
		if currentEl.IsSelected() {
			currentEl.SetSelected(false)
		}
	}
	newSelectedEl.SetSelected(true)
}

func (p *PageFirst) SelectNext() {
	for el := p.elementsList.Front(); el != nil; el = el.Next() {
		currentEl := el.Value.(*ui.SelectableElement)
		if currentEl.IsSelected() {
			var newSelectedEl *ui.SelectableElement
			if el.Next() != nil {
				newSelectedEl = el.Next().Value.(*ui.SelectableElement)
			} else {
				newSelectedEl = p.elementsList.Front().Value.(*ui.SelectableElement)
			}
			currentEl.SetSelected(false)
			newSelectedEl.SetSelected(true)
			return
		}
	}
}

func (p *PageFirst) SelectPrev() {
	for el := p.elementsList.Front(); el != nil; el = el.Next() {
		currentEl := el.Value.(*ui.SelectableElement)
		if currentEl.IsSelected() {
			var newSelectedEl *ui.SelectableElement
			if el.Prev() != nil {
				newSelectedEl = el.Prev().Value.(*ui.SelectableElement)
			} else {
				newSelectedEl = p.elementsList.Back().Value.(*ui.SelectableElement)
			}
			currentEl.SetSelected(false)
			newSelectedEl.SetSelected(true)
			return
		}
	}
}
