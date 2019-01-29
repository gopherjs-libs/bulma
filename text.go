package bulma

import "github.com/gopherjs/vecty"

func Text(text string) vecty.ComponentOrHTML {
	return &TextComponent{text: text}
}
