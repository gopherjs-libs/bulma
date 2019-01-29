package bulma

import (
	"github.com/gopherjs/vecty"
)

func Progress(css vecty.ClassMap, value string, max string, indeterminate bool) vecty.ComponentOrHTML {
	return &ProgressComponent{css: css, value: value, max: max, indeterminate: indeterminate}
}
