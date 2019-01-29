package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type TitleComponent struct {
	vecty.Core

	title  string
	css    vecty.ClassMap
	value  string
	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *TitleComponent) Render() vecty.ComponentOrHTML {
	return elem.Paragraph(vecty.Markup(vecty.Class(t.title), t.Markup, t.css), Text(t.value))
}
