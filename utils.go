package bulma

import (
	"fmt"
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

type EleHandle func(markup ...vecty.MarkupOrChild) *vecty.HTML
type EventHandle func(*vecty.Event)

func EleIf(b bool, a, a1 EleHandle) EleHandle {
	if b {
		return a
	}

	return a1
}

type StrM map[string]string

func Kv(k, v string) StrM {
	return StrM{k: v}
}

func init() {
	fmt.Println([]StrM{{"s": "22"}})
}
