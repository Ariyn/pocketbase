package pocketbase

type ParamsList struct {
	Page      int
	Size      int
	Filters   string
	Sort      string
	Expand    string
	Fields    string
	SkipTotal bool

	queryParams map[string]string // additional query parameters

	hackResponseRef any //hack for collection list
}

func (p *ParamsList) SetQueryParam(key, value string) {
	if p.queryParams == nil {
		p.queryParams = make(map[string]string)
	}
	p.queryParams[key] = value
}
