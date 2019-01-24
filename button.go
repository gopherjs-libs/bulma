package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Button struct {
	vecty.Core

	_stateMap
	_colorMap
	_sizeMap

	Markup   vecty.MarkupList
	Children vecty.ComponentOrHTML
}

func (t *Button) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("button"),
			vecty.ClassMap{
				State.IsPrimary:  t.IsPrimary,
				State.IsLink:     t.IsLink,
				State.IsInfo:     t.IsInfo,
				State.IsSuccess:  t.IsSuccess,
				State.IsWarning:  t.IsWarning,
				State.IsDanger:   t.IsDanger,
				State.IsHovered:  t.IsHovered,
				State.IsLoading:  t.IsLoading,
				State.IsFocused:  t.IsFocused,
				State.IsActive:   t.IsActive,
				State.IsStatic:   t.IsStatic,
				State.IsOutlined: t.IsOutlined,
				State.Disabled:   t.Disabled,
				State.IsInverted: t.IsInverted,
				State.IsRounded:  t.IsRounded,

				Color.IsWhite: t.IsWhite,
				Color.IsLight: t.IsLight,
				Color.IsDark:  t.IsDark,
				Color.IsBlack: t.IsBlack,
				Color.IsText:  t.IsText,

				Size.IsSmall:     t.IsSmall,
				Size.IsMedium:    t.IsMedium,
				Size.IsLarge:     t.IsLarge,
				Size.IsNormal:    t.IsNormal,
				Size.IsFullwidth: t.IsFullwidth,
			},
		),
	)
}
