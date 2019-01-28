package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Del(applyer ...vecty.Applyer) vecty.ComponentOrHTML {
	return elem.Button(vecty.Markup(append(applyer, vecty.Class("delete"))...))
}
