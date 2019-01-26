package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Content struct {
	vecty.Core

	Type     string
	Modifier string
	Size     string
	Markup   vecty.MarkupList
	Slot     vecty.List
}

func (t *Content) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("content"),
		vecty.MarkupIf(t.Type != "", vecty.Attribute("type", t.Type)),
		ClassMap(t.Size, t.Modifier)),
		t.Markup,
		t.Slot)
}

func SimpleContent() {

}
