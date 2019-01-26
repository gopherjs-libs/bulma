package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type Image struct {
	vecty.Core

	Url       string
	Size      string
	IsRounded bool

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *Image) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("image"),
		ClassMap(t.Size)),
		elem.Image(vecty.Markup(prop.Src(t.Url))),
		t.Markup, t.Slot)
}

func SimpleImage(m ...string) func(url string) vecty.ComponentOrHTML {
	return func(url string) vecty.ComponentOrHTML {
		return elem.Div(vecty.Markup(vecty.Class("image"),
			ClassMap(m...)),
			elem.Image(vecty.Markup(prop.Src(url))))
	}
}

func RoundedImage(m ...string) func(url string) vecty.ComponentOrHTML {
	return func(url string) vecty.ComponentOrHTML {
		return elem.Div(vecty.Markup(vecty.Class("image", "is-rounded"),
			ClassMap(m...)),
			elem.Image(vecty.Markup(prop.Src(url))))
	}
}
