package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type ImageComponent struct {
	vecty.Core

	IsRounded bool
	Src       string
	Markup    vecty.MarkupList
	Slot      vecty.List
}

func (t *ImageComponent) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("image")), t.Markup,
		elem.Image(vecty.Markup(vecty.Class("image"), vecty.ClassMap{"is-rounded": t.IsRounded}, prop.Src(t.Src))))
}
