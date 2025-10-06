package pages

import (
	"container/list"

	"github.com/charmbracelet/lipgloss"
	"github.com/pseudoelement/galaga/src/menu/models"
)

type Page struct {
	elementsList *list.List
	injector     models.IAppInjector
}

func NewPage(injector models.IAppInjector) *Page {
	return &Page{elementsList: list.New(), injector: injector}
}

func (p *Page) View() string {
	elementsSlice := make([]models.IBaseViewElement, 0, p.elementsList.Len())
	for el := p.elementsList.Front(); el != nil; el = el.Next() {
		selectableEl, ok := el.Value.(models.ISelectableElement)
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
		elementsSlice = append(elementsSlice, el.Value.(models.IBaseViewElement))
	}

	return lipgloss.JoinVertical(lipgloss.Left, p.ElementsToViews(elementsSlice)...)
}

func (p *Page) ElementsList() *list.List {
	return p.elementsList
}

func (p *Page) ElementsToViews(elements []models.IBaseViewElement) []string {
	var sl []string
	for _, el := range elements {
		sl = append(sl, el.ViewString())
	}

	return sl
}

func (p *Page) SelectedItem() models.ISelectableElement {
	for el := p.elementsList.Front(); el != nil; el = el.Next() {
		selectableEl, ok := el.Value.(models.ISelectableElement)
		if ok && selectableEl.IsSelected() {
			return selectableEl
		}
	}

	return nil
}

func (p *Page) SelectFirstSelectable() {
	for el := p.elementsList.Front(); el != nil; el = el.Next() {
		currentEl, selectable := el.Value.(models.ISelectableElement)
		if selectable {
			currentEl.SetSelected(true)
			break
		}
	}
}

func (p *Page) SelectLastSelectable() {
	var lastSelectableEl models.ISelectableElement = nil
	for el := p.elementsList.Front(); el != nil; el = el.Next() {
		currentEl, selectable := el.Value.(models.ISelectableElement)
		if selectable {
			lastSelectableEl = currentEl
			break
		}
	}
	if lastSelectableEl != nil {
		lastSelectableEl.SetSelected(true)
	}
}

func (p *Page) Select(newSelectedEl models.ISelectableElement) {
	for el := p.elementsList.Front(); el != nil; el = el.Next() {
		currentEl, selectable := el.Value.(models.ISelectableElement)
		if selectable && currentEl.IsSelected() {
			currentEl.SetSelected(false)
			break
		}
	}
	newSelectedEl.SetSelected(true)
}

func (p *Page) SelectNext() {
	if p.elementsList.Len() <= 1 {
		return
	}
	for el := p.elementsList.Front(); el != nil; el = el.Next() {
		currentEl, selectable := el.Value.(models.ISelectableElement)
		if selectable && currentEl.IsSelected() {
			var newSelectedEl models.ISelectableElement
			if el.Next() != nil {
				nextElSelectable, isNextSelectable := el.Next().Value.(models.ISelectableElement)
				if isNextSelectable {
					newSelectedEl = nextElSelectable
				}
			} else {
				firstElSelectable, isFirstSelectable := p.elementsList.Front().Value.(models.ISelectableElement)
				if isFirstSelectable {
					newSelectedEl = firstElSelectable
				}
			}

			currentEl.SetSelected(false)

			if newSelectedEl != nil {
				newSelectedEl.SetSelected(true)
			} else {
				p.SelectFirstSelectable()
			}

			return
		}
	}
}

func (p *Page) SelectPrev() {
	if p.elementsList.Len() <= 1 {
		return
	}
	for el := p.elementsList.Front(); el != nil; el = el.Next() {
		currentEl, selectable := el.Value.(models.ISelectableElement)
		if selectable && currentEl.IsSelected() {
			var newSelectedEl models.ISelectableElement
			if el.Prev() != nil {
				// @TODO fix when first element not selectable, it should transfer selector to last element
				prevElSelectable, isPrevSelectable := el.Prev().Value.(models.ISelectableElement)
				if isPrevSelectable {
					newSelectedEl = prevElSelectable
				}
			} else {
				lastElSelectable, isLastSelectable := p.elementsList.Back().Value.(models.ISelectableElement)
				if isLastSelectable {
					newSelectedEl = lastElSelectable
				}
			}
			currentEl.SetSelected(false)

			if newSelectedEl != nil {
				newSelectedEl.SetSelected(true)
			} else {
				p.SelectLastSelectable()
			}

			return
		}
	}
}
