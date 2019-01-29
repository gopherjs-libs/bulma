package bulma

import (
	"github.com/gopherjs/vecty"
)

func Del(applyer ...vecty.Applyer) vecty.ComponentOrHTML {
	return &DeleteComponent{Markup: vecty.Markup(applyer...)}
}
