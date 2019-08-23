package charts

import "net/http"

type Chart struct {
	Base
}

func (C Chart) Plot(w http.ResponseWriter, r *http.Request) {
	toHandle(w, r, C.Base)
}
func (C Chart) Render(w http.ResponseWriter, r *http.Request) {
	C.Plot(w, r)
}
func (C Chart) Save(name string) bool {
	if name == "" {
		name = "chart"
	}
	return toLocal(name, C.Base)
}

func (C Chart) Name() string {
	return "chart"
}
func (C Chart) String() string {
	return "chart"
}

// 自定义类型的：即组合的形式
func NewChart(title string, t string) Chart {
	var option Options
	option = make(map[string]interface{})
	option.SetTitle(title)
	return Chart{
		Base: Base{
			Type:    t,
			Data:    defaultData(),
			Options: option,
		},
	}
}
