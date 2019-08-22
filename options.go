package charts

type Options map[string]interface{}

// 支持自定义结构体设置 options
func (O Options) SetProperty(key string, v interface{}) {
	O[key] = v
}

func (O Options) SetResponsive(v bool) {
	O["responsive"] = v
}

func (O Options) SetTitle(title string) {
	O["title"] = Title{
		Display: true,
		Text:    title,
	}
}

type Title struct {
	Display  bool   `json:"display,omitempty"`
	Text     string `json:"text,omitempty"`
	Position string `json:"position,omitempty"`
}

func (O Options) SetToolTips(tips ToolTips) {
	O["tooltips"] = tips
}

type ToolTips struct {
	Mode      string `json:"mode,omitempty"`
	Intersect bool   `json:"intersect,omitempty"`
}

func (O Options) SetHover(hover Hover) {
	O["hover"] = hover
}

type Hover struct {
	Mode      string `json:"mode,omitempty"`
	Intersect bool   `json:"intersect,omitempty"`
}

func (O Options) SetSales(scales Scales) {
	O["scales"] = scales
}

type Scales struct {
	X []Axes `json:"xAxes,omitempty"`
	Y []Axes `json:"yAxes,omitempty"`
}

type AxesOptions map[string]interface{}

func (A AxesOptions) AddAxesOption(key string, val interface{}) {
	A[key] = val
}

type Axes struct {
	Display    bool       `json:"display,omitempty"`
	ScaleLabel ScaleLabel `json:"scaleLabel,omitempty"`
	Ticks      Ticks      `json:"ticks,omitempty"`
	Position   string     `json:"position,omitempty"`
}

type ScaleLabel struct {
	Display     bool   `json:"display,omitempty"`
	LabelString string `json:"labelString,omitempty"`
}

type Ticks struct {
	BeginAtZero bool `json:"beginAtZero,omitempty"`
}

func (O Options) AxesInit() {
	if _, ok := O["scales"]; !ok {
		O["scales"] = Scales{
			X: []Axes{
				{
					ScaleLabel: ScaleLabel{
						Display:     true,
						LabelString: "",
					},
				},
			},
			Y: []Axes{
				{
					ScaleLabel: ScaleLabel{
						Display:     true,
						LabelString: "",
					},
				},
			},
		}
	}
}

// SetXAxes 替换默认的 X 轴的显示
func (O Options) SetXAxes(label []string) bool {
	O.AxesInit()
	X := O["scales"].(Scales).X
	if len(label) != len(X) {
		return false
	}
	tempX := X
	for index, _ := range tempX {
		tempX[index].ScaleLabel.LabelString = label[index]
	}
	X = tempX
	return true
}

func (O Options) SetXAxesBeginAtZero() {
	O.AxesInit()
	X := O["scales"].(Scales).X
	tempX := X
	for index, _ := range X {
		tempX[index].Ticks.BeginAtZero = true
	}
	X = tempX
}

func (O Options) SetYAxes(label []string) bool {
	O.AxesInit()
	Y := O["scales"].(Scales).Y
	if len(label) != len(Y) {
		return false
	}
	tempY := Y
	for index, _ := range Y {
		tempY[index].ScaleLabel.LabelString = label[index]
	}
	Y = tempY
	return true
}

func (O Options) SetYAxesBeginAtZero() {
	O.AxesInit()
	Y := O["scales"].(Scales).Y
	tempY := Y
	for index, _ := range Y {
		tempY[index].Ticks.BeginAtZero = true
	}
	Y = tempY
}

func (O Options) SetXYAxes(Xlabel []string, Ylable []string) {
	O.SetXAxes(Xlabel)
	O.SetYAxes(Ylable)
}

func (O Options) SetXYAxesBeginAtZero() {
	O.SetXAxesBeginAtZero()
	O.SetYAxesBeginAtZero()
}

func (O Options) SetAppendAxes(x bool, label []string, zero bool) {
	if x {
		X := O["scales"].(Scales).X
		tempX := X
		for _, i := range label {
			tempX = append(tempX, Axes{
				Display: true,
				ScaleLabel: ScaleLabel{
					Display:     true,
					LabelString: i,
				},
				Ticks: Ticks{
					BeginAtZero: zero,
				},
			})
		}
		X = tempX
	} else {
		Y := O["scales"].(Scales).Y
		tempY := Y
		for _, i := range label {
			tempY = append(tempY, Axes{
				Display: true,
				ScaleLabel: ScaleLabel{
					Display:     true,
					LabelString: i,
				},
				Ticks: Ticks{
					BeginAtZero: zero,
				},
			})
		}
		Y = tempY

	}
}

type LayOut struct {
	Padding struct {
		Left   int `json:"left"`
		Right  int `json:"right"`
		Top    int `json:"top"`
		Bottom int `json:"bottom"`
	} `json:"padding"`
}

func (O Options) SetLegend(position string, reverse bool) {
	O["legend"] = Legend{
		Display:  true,
		Position: position,
		Reverse:  reverse,
	}
}

type Legend struct {
	Display  bool   `json:"display,omitempty"`
	Position string `json:"position,omitempty"`
	Reverse  bool   `json:"reverse,omitempty"`
}

func defaultOptions(title string) Options {
	var options Options
	options = make(map[string]interface{})
	options.SetTitle(title)
	options.SetResponsive(true)

	options.defaultAxes("", "", false, false)
	return options
}

func (O Options) defaultAxes(x string, y string, xZero bool, yZero bool) {
	X := []Axes{
		{
			Display: true,
			ScaleLabel: ScaleLabel{
				Display:     true,
				LabelString: x,
			},
			Ticks: Ticks{
				BeginAtZero: xZero,
			},
			Position: LEFT,
		},
	}
	Y := []Axes{
		{
			Display: true,
			ScaleLabel: ScaleLabel{
				Display:     true,
				LabelString: y,
			},
			Ticks: Ticks{
				BeginAtZero: yZero,
			},
			Position: BOTTOM,
		},
	}
	O["scales"] = Scales{
		X: X,
		Y: Y,
	}
}
