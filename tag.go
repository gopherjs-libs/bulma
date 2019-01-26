package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)





type Tag struct {
	vecty.Core

	Color    string
	Size     string
	Modifier string
	Value    string
}

func (t *Tag) Render() vecty.ComponentOrHTML {
	return elem.Span(
		vecty.Markup(
			vecty.Class("tag"),
			ClassMap(t.Color, t.Size, t.Modifier),
		),
		vecty.Text(t.Value),
	)
}

var TagsSize = struct {
	AreSmall  string
	AreMedium string
	AreLarge  string
}{
	AreSmall:  "are-small",
	AreMedium: "are-medium",
	AreLarge:  "are-large",
}

type Tags struct {
	vecty.Core
	Size string

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *Tags) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(
		vecty.Class("tags"),
		ClassMap(t.Size),
	), t.Markup, t.Slot)
}

type AddonsTags struct {
	vecty.Core

	Size string

	OnClick EventHandle

	Markup vecty.MarkupList
	Slot   *Tags
}

func (t *AddonsTags) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("field", "is-grouped", "is-grouped-multiline"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("control"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("tags", "has-addons"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("tag", "is-dark"),
					),
					vecty.Text("npm"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("tag", "is-info"),
					),
					vecty.Text("0.5.0"),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("control"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("tags", "has-addons"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("tag", "is-dark"),
					),
					vecty.Text("build"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("tag", "is-success"),
					),
					vecty.Text("passing"),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("control"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("tags", "has-addons"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("tag", "is-dark"),
					),
					vecty.Text("chat"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("tag", "is-primary"),
					),
					vecty.Text("on gitter"),
				),
			),
		),
	)
}
