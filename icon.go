package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

var IconSize = struct {
	IsSmall  string
	IsMedium string
	IsLarge  string
}{
	IsSmall:  "is-small",
	IsMedium: "is-medium",
	IsLarge:  "is-large",
}

var IconColors = struct {
	HasTextInfo    string
	HasTextSuccess string
	HasTextWarning string
	HasTextDanger  string
}{
	HasTextInfo:    "has-text-info",
	HasTextSuccess: "has-text-success",
	HasTextWarning: "has-text-warning",
	HasTextDanger:  "has-text-danger",
}

type I struct {
	vecty.Core

	Value string
	Size  string
	Color string
}

func (t *I) Render() vecty.ComponentOrHTML {
	return elem.Span(vecty.Markup(vecty.Class("fas"), ClassMap(t.Size, t.Color,t.Value)))
}

type Icon struct {
	vecty.Core

	Size     string
	Color    string
	Slot vecty.ComponentOrHTML
}

func (t *Icon) Render() vecty.ComponentOrHTML {
	return elem.Span(
		vecty.Markup(
			vecty.Class("icon"),
			ClassMap(t.Size, t.Color),
		),
		t.Slot,
	)
}

type IconStacked struct {
	vecty.Core

	Size  string
	Color string
	Slot  vecty.ComponentOrHTML
}

func (t *IconStacked) Render() vecty.ComponentOrHTML {
	return elem.Span(
		vecty.Markup(
			vecty.Class("fa-stack"),
			ClassMap(t.Size, t.Color),
		),
		t.Slot,
	)
}
