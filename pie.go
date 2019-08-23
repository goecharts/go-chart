package charts

import "net/http"

type Pie struct {
	Base
}

func (P Pie) Plot(w http.ResponseWriter, r *http.Request) {
	toHandle(w, r, P.Base)
}

func (P Pie) Render(w http.ResponseWriter, r *http.Request) {
	P.Plot(w, r)
}

func (P Pie) Save(name string) bool {
	if name == "" {
		name = "pie"
	}
	return toLocal(name, P.Base)
}

func (P Pie) Name() string {
	return P.Type
}

func (P Pie) String() string {
	return PieTypeType
}

func NewPie(title string) Pie {
	var option Options
	option = make(map[string]interface{})
	option.SetTitle(title)
	return Pie{
		Base: Base{
			Type:    PieTypeType,
			Data:    defaultData(),
			Options: option,
		},
	}
}
