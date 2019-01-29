package bulma

import (
	"github.com/gopherjs/vecty"
)

func Image(css vecty.ClassMap, applyer ...vecty.Applyer) func(src string) vecty.ComponentOrHTML {
	return func(src string) vecty.ComponentOrHTML {
		return &ImageComponent{Markup: vecty.Markup(append(applyer, css)...), IsRounded: false, Src: src}
	}
}

func RoundedImage(css vecty.ClassMap, applyer ...vecty.Applyer) func(url string) vecty.ComponentOrHTML {
	return func(src string) vecty.ComponentOrHTML {
		return &ImageComponent{Markup: vecty.Markup(append(applyer, css)...), IsRounded: true, Src: src}
	}
}
