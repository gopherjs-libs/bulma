package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type StackedIconComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *StackedIconComponent) Render() vecty.ComponentOrHTML {
	return elem.Span(vecty.Markup(vecty.Class("fa-stack"), t.Markup), t.Slot)
}

type IComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *IComponent) Render() vecty.ComponentOrHTML {
	return elem.Span(vecty.Markup(vecty.Class("fas")), t.Markup, t.Slot)
}

type IconComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *IconComponent) Render() vecty.ComponentOrHTML {
	return elem.Span(vecty.Markup(vecty.Class("icon")), t.Markup, MapElem(t.Slot, func(c vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		switch _c := c.(type) {
		case *IconComponent:
			return _c
		case *TextComponent:
			return elem.Span(_c)
		default:
			return _c
		}
	}))
}
