package pocketbase

type ParamsList struct {
	Page      int
	Size      int
	Filters   string
	Sort      string
	Expand    string
	Fields    string
	SkipTotal bool

	hackResponseRef any //hack for collection list
}
