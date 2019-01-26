package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Notification struct {
	vecty.Core

	Color string

	OnClick EventHandle
	Value   string
}

func (t *Notification) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("notification"),
		ClassMap(t.Color)),
		&Delete{OnClick: t.OnClick},
		vecty.Text(t.Value))
}
