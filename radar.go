package charts

import "net/http"

type Radar struct {
	Base
}

func (R Radar) Plot(w http.ResponseWriter, r *http.Request) {
	toHandle(w, r, R.Base)
}

func (R Radar) Name() string {
	return R.Type
}
func (R Radar) Save(name string) bool {
	if name == "" {
		name = "radar"
	}
	return toLocal(name, R.Base)
}

func (R Radar) String() string {
	return RadarType
}

func NewRadar(title string) *Radar {
	var option Options
	option = make(map[string]interface{})
	option.SetTitle(title)
	return &Radar{
		Base: Base{
			Type:    RadarType,
			Data:    defaultData(),
			Options: option,
		},
	}
}
