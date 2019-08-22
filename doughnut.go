package charts

import "net/http"

type Doughnut struct {
	Base
}

func (D Doughnut) Plot(w http.ResponseWriter, r *http.Request) {
	toHandle(w, r, D.Base)
}

func (D Doughnut) Save(name string) bool {
	if name == "" {
		name = "doughnut"
	}
	return toLocal(name, D.Base)
}
func (D Doughnut) Name() string {
	return D.Type
}
func (D Doughnut) String() string {
	return DoughnutType
}

func NewDoughnut(title string) Doughnut {
	var option Options
	option = make(map[string]interface{})
	option.SetTitle(title)
	return Doughnut{
		Base: Base{
			Type:    DoughnutType,
			Data:    defaultData(),
			Options: option,
		},
	}
}
