package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type MultilineGroupedComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *MultilineGroupedComponent) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("field", "is-grouped", "is-grouped-multiline")), t.Markup,
		MapElem(t.Slot, func(c vecty.ComponentOrHTML) vecty.ComponentOrHTML {
			return elem.Div(vecty.Markup(vecty.Class("control")), c)
		}))
}

type FieldComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *FieldComponent) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("field")), t.Slot)
}

type BlockComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *BlockComponent) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("block")), t.Markup, t.Slot)
}
