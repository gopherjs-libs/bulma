package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Buttons struct {
	vecty.Core

	AreSmall  bool
	AreMedium bool
	AreLarge  bool

	Markup   vecty.MarkupList
	Children vecty.ComponentOrHTML
}

func (t *Buttons) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("buttons"),
			vecty.ClassMap{
				"are-small":  t.AreSmall,
				"are-medium": t.AreMedium,
				"are-large":  t.AreLarge,
			},
		),
		t.Markup,
		t.Children,
	)
}
