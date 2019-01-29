package bulma

import (
	"github.com/gopherjs/vecty"
)

func Title(css vecty.ClassMap, value string) vecty.ComponentOrHTML {
	return &TitleComponent{title: "title", css: css, value: value}
}

func Subtitle(css vecty.ClassMap, value string) vecty.ComponentOrHTML {
	return &TitleComponent{title: "subtitle", css: css, value: value}
}
