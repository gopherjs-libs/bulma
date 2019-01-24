package bulma

type _colorMap struct {
	IsWhite bool
	IsLight bool
	IsDark  bool
	IsBlack bool
	IsText  bool
}

type _color struct {
	IsWhite string
	IsLight string
	IsDark  string
	IsBlack string
	IsText  string
}

var Color = _color{
	IsWhite: "is-white",
	IsLight: "is-light",
	IsDark:  "is-dark",
	IsBlack: "is-black",
	IsText:  "is-text",
}
