package bulma

import (
	"github.com/gopherjs/vecty"
)

func CpsOf(c ...vecty.ComponentOrHTML) (l vecty.List) {
	return c
}

func ClassMap(m ...string) vecty.ClassMap {
	_css := vecty.ClassMap{}
	for _, i := range m {
		if i == "" {
			continue
		}
		_css[i] = true
	}
	return _css
}
