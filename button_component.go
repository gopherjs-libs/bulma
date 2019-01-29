package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type ButtonComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *ButtonComponent) Render() vecty.ComponentOrHTML {
	return elem.Anchor(vecty.Markup(vecty.Class("button")), t.Markup, vecty.List(t.Slot))
}

type ButtonsComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *ButtonsComponent) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("buttons"), t.Markup), t.Slot)
}

type GroupButtonComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *GroupButtonComponent) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("field", "is-grouped"), t.Markup,
	), MapElem(t.Slot, func(c vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Paragraph(vecty.Markup(vecty.Class("control")), c)
	}))
}

type AddonsButtonComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *AddonsButtonComponent) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("field", "has-addons"), t.Markup,
	), MapElem(t.Slot, func(c vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Paragraph(vecty.Markup(vecty.Class("control")), c)
	}))
}

type IconButtonComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *IconButtonComponent) Render() vecty.ComponentOrHTML {
	var _p []vecty.ComponentOrHTML
	for _, _c := range t.Slot {
		switch _i := _c.(type) {
		case *IconComponent:
			_p = append(_p, _i)
		case *TextComponent:
			_p = append(_p, elem.Span(_i))
		}
	}
	return elem.Anchor(vecty.Markup(vecty.Class("button")), t.Markup, vecty.List(_p))
}
