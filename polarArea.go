package charts

import "net/http"

type PolarArea struct {
	Base
}

func (P PolarArea) Plot(w http.ResponseWriter, r *http.Request) {
	toHandle(w, r, P.Base)
}

func (P PolarArea) Render(w http.ResponseWriter, r *http.Request) {
	P.Plot(w, r)
}

func (P PolarArea) Save(name string) bool {
	if name == "" {
		name = "polarArea"
	}
	return toLocal(name, P.Base)
}

func (P PolarArea) Name() string {
	return P.Type
}
func (P PolarArea) String() string {
	return PolarAreaType
}

func NewPolarArea(title string) PolarArea {
	var option Options
	option = make(map[string]interface{})
	option.SetTitle(title)
	return PolarArea{
		Base: Base{
			Type:    PolarAreaType,
			Data:    defaultData(),
			Options: option,
		},
	}
}
