package charts

import "net/http"

type Scatter struct {
	Base
}

func (S Scatter) Plot(w http.ResponseWriter, r *http.Request) {
	toHandle(w, r, S.Base)
}

func (S Scatter) Save(name string) bool {
	if name == "" {
		name = "scatter"
	}
	return toLocal(name, S.Base)
}

func (S Scatter) Name() string {
	return S.Type
}
func (S Scatter) String() string {
	return ScatterType
}

func NewScatter(title string) Scatter {
	var option Options
	option = make(map[string]interface{})
	option.SetTitle(title)
	return Scatter{
		Base: Base{
			Type:    ScatterType,
			Data:    defaultData(),
			Options: option,
		},
	}
}
