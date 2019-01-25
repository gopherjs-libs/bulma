package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

type Delete struct {
	vecty.Core

	Size string

	OnClick EventHandle

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *Delete) Render() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			vecty.Class("delete"),
			ClassMap(t.Size),
			vecty.MarkupIf(t.OnClick != nil, event.Click(t.OnClick)),
		), t.Markup, t.Slot)
}
