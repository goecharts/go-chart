package charts

import "net/http"

type Bar struct {
	Base
}

func (B Bar) Plot(w http.ResponseWriter, req *http.Request) {
	toHandle(w, req, B.Base)
}

func (B Bar) Render(w http.ResponseWriter, req *http.Request) {
	B.Plot(w, req)
}

func (B Bar) Save(name string) bool {
	if name == "" {
		name = "bar"
	}
	return toLocal(name, B.Base)
}
func (B Bar) Name() string {
	return B.Type
}

func (B Bar) String() string {
	return BarType
}

func NewBar(title string) *Bar {
	return &Bar{
		Base: Base{
			Type:    BarType,
			Data:    defaultData(),
			Options: defaultOptions(title),
		},
	}
}

func NewHorizontalBar(title string) *Bar {
	return &Bar{
		Base: Base{
			Type:    BarHorizonType,
			Data:    defaultData(),
			Options: defaultOptions(title),
		},
	}
}
