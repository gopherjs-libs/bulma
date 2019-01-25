package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

var ContentType = struct {
	IsLowerAlpha string
	IsLowerRoman string
	IsUpperAlpha string
	IsUpperRoman string
}{
	IsLowerAlpha: "is-lower-alpha",
	IsLowerRoman: "is-lower-roman",
	IsUpperAlpha: "is-upper-alpha",
	IsUpperRoman: "is-upper-roman",
}

var ContentState = struct {
	T1 string
	TA string
	Ta string
	TI string
	Ti string
}{
	T1: "1",
	TA: "A",
	Ta: "a",
	TI: "I",
	Ti: "i",
}

type Content struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *Content) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("content"),
		),
		t.Markup, t.Slot,
	)
}
