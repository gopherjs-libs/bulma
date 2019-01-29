package bulma

import (
	"github.com/gopherjs/vecty"
)

func MultilineGrouped(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return &MultilineGroupedComponent{Slot: vecty.List(c)}
}

func Field(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return &FieldComponent{Slot: vecty.List(c)}
}

func Block(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return &BlockComponent{Slot: vecty.List(c)}
}
