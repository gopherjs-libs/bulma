package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type IconSize struct {
	IsSmall  bool
	IsMedium bool
	IsLarge  bool
}

type IconColors struct {
	HasTextInfo    bool
	HasTextSuccess bool
	HasTextWarning bool
	HasTextDanger  bool
}

type I struct {
	vecty.Core

	Class  []string
	Colors IconColors
}

func (t *I) Render() vecty.ComponentOrHTML {
	_css := vecty.ClassMap{
		"has-text-info":    t.Colors.HasTextInfo,
		"has-text-success": t.Colors.HasTextSuccess,
		"has-text-warning": t.Colors.HasTextWarning,
		"has-text-danger":  t.Colors.HasTextDanger,
	}

	for _, _c := range t.Class {
		_css[_c] = true
	}

	return elem.Span(
		vecty.Markup(
			vecty.Class("icon"),
			_css,
		),
	)
}

type Icon struct {
	vecty.Core

	Colors   IconColors
	Size     IconSize
	Markup   vecty.MarkupList
	Children vecty.ComponentOrHTML
}

func (t *Icon) Render() vecty.ComponentOrHTML {
	return elem.Span(
		vecty.Markup(
			vecty.Class("icon"),
			vecty.ClassMap{
				"has-text-info":    t.Colors.HasTextInfo,
				"has-text-success": t.Colors.HasTextSuccess,
				"has-text-warning": t.Colors.HasTextWarning,
				"has-text-danger":  t.Colors.HasTextDanger,
				"is-small":         t.Size.IsSmall,
				"is-medium":        t.Size.IsMedium,
				"is-large":         t.Size.IsLarge,
			},
		),
		t.Markup,
		t.Children,
	)
}
