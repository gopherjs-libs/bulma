package attrs

import "github.com/gopherjs/vecty"

func Href(s string) vecty.Applyer {
	return vecty.Attribute("href", s)
}

func Title(s string) vecty.Applyer {
	return vecty.Attribute("title", s)
}

func Value(s string) vecty.Applyer {
	return vecty.Attribute("value", s)
}

func Max(s string) vecty.Applyer {
	return vecty.Attribute("max", s)
}
