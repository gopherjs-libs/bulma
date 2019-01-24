package bulma

type _stateMap struct {
	IsPrimary  bool
	IsLink     bool
	IsInfo     bool
	IsSuccess  bool
	IsWarning  bool
	IsDanger   bool
	IsHovered  bool
	IsLoading  bool
	IsFocused  bool
	IsActive   bool
	IsStatic   bool
	IsOutlined bool
	Disabled   bool
	IsInverted bool
	IsRounded  bool
}

type _state struct {
	IsPrimary  string
	IsLink     string
	IsInfo     string
	IsSuccess  string
	IsWarning  string
	IsLoading  string
	IsDanger   string
	IsHovered  string
	IsFocused  string
	IsActive   string
	IsStatic   string
	IsOutlined string
	Disabled   string
	IsInverted string
	IsRounded  string
}

var State = _state{
	IsPrimary:  "is-primary",
	IsLink:     "is-link",
	IsInfo:     "is-info",
	IsSuccess:  "is-success",
	IsWarning:  "is-warning",
	IsDanger:   "is-danger",
	IsHovered:  "is-hovered",
	IsFocused:  "is-focused",
	IsActive:   "is-active",
	IsLoading:  "is-loading",
	IsStatic:   "is-static",
	IsOutlined: "is-outlined",
	Disabled:   "disabled",
	IsInverted: "is-inverted",
	IsRounded:  "is-rounded",
}
