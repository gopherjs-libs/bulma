package bulma

import (
	"github.com/gopherjs/vecty"
)

func Notification(css vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &NotificationComponent{Css: css, Markup: vecty.Markup(applyer...), Slot: vecty.List(c)}
	}
}
