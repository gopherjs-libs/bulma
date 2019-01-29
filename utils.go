package bulma

import (
	"fmt"
	"github.com/gopherjs/vecty"
	"strings"
)

func Components(c ...vecty.ComponentOrHTML) (l vecty.List) {
	return c
}

func Str(c ...string) []string {
	return c
}

func Css(m ...string) vecty.ClassMap {
	return ClassMap(m...)
}

func ClassMap(m ...string) vecty.ClassMap {
	_css := vecty.ClassMap{}
	for _, i := range m {
		if i == "" {
			continue
		}

		for _, _s := range strings.Split(i, " ") {
			_css[_s] = true
		}

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

type M map[string]interface{}

func (t M) Map(fn func(k string, v interface{}) vecty.ComponentOrHTML) MapComponent {
	_dm := make(MapComponent)
	for k, v := range t {
		_dm[k] = fn(k, v)
	}
	return _dm
}

func Kv(k, v string) StrM {
	return StrM{k: v}
}

func init() {
	fmt.Println([]StrM{{"s": "22"}})
}

type MapComponent map[string]vecty.ComponentOrHTML

func BodyMap(dm []M, fn func(m M) MapComponent) []MapComponent {
	var _ds []MapComponent
	for _, m := range dm {
		_ds = append(_ds, fn(m))
	}
	return _ds
}

type ComponentFn func(style ...string) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML

func MapElem(cpt []vecty.ComponentOrHTML, fn func(c vecty.ComponentOrHTML) vecty.ComponentOrHTML) vecty.List {
	var _cs vecty.List
	for _, _i := range cpt {
		_cs = append(_cs, fn(_i))
	}
	return _cs
}
