package charts

import (
	"net/http"
)

type Line struct {
	Base
}

func (L Line) Plot(w http.ResponseWriter, r *http.Request) {
	toHandle(w, r, L.Base)
}

func (L Line) Render(w http.ResponseWriter, r *http.Request) {
	L.Plot(w, r)
}

func (L Line) Save(name string) bool {
	if name == "" {
		name = "line"
	}
	return toLocal(name, L.Base)
}

func (L Line) Name() string {
	return L.Base.Type
}
func (L Line) String() string {
	return LineType
}

func NewLine(title string) *Line {

	return &Line{
		Base: Base{
			Type:    LineType,
			Data:    defaultData(),
			Options: defaultOptions(title),
		},
	}
}
