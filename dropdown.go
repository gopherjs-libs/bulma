package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

type Dropdown struct {
	vecty.Core

	Title    string
	IsActive bool
	OnClick  func(bool)
	Slot     vecty.Component
}

func (t *Dropdown) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("dropdown"),
			vecty.ClassMap{"is-active": t.IsActive},
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("dropdown-trigger"),
			),
			elem.Button(
				vecty.Markup(
					vecty.Class("button"),
					vecty.Attribute("aria-haspopup", "true"),
					vecty.Attribute("aria-controls", "dropdown-menu"),
					event.Click(func(i *vecty.Event) {
						t.IsActive = !t.IsActive
						vecty.Rerender(t)
					}),
				),
				elem.Span(
					Text("Dropdown button"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("icon", "is-small"),
					),
					elem.Italic(
						vecty.Markup(
							vecty.Class("fas", "fa-angle-down"),
							vecty.Attribute("aria-hidden", "true"),
						),
					),
				),
			),
		),
		t.Slot,
	)
}
