package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

func Image(css vecty.ClassMap, applyer ...vecty.Applyer) func(src string) vecty.ComponentOrHTML {
	return func(src string) vecty.ComponentOrHTML {
		return elem.Div(vecty.Markup(vecty.Class("image"), css),
			vecty.Markup(applyer...),
			elem.Image(vecty.Markup(prop.Src(src))))
	}
}

func RoundedImage(css vecty.ClassMap, applyer ...vecty.Applyer) func(url string) vecty.ComponentOrHTML {
	return func(src string) vecty.ComponentOrHTML {
		return elem.Div(vecty.Markup(vecty.Class("image"), css),
			vecty.Markup(applyer...),
			elem.Image(vecty.Markup(vecty.Class("image", "is-rounded"), prop.Src(src))))
	}
}
