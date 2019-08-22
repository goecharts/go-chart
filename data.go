package charts

type Data map[string]interface{}

func (D Data) SetLabels(labels Labels) {
	D["labels"] = labels
}

type Labels []interface{}

func (D Data) SetDataSet(dataSets []interface{}) {
	D["datasets"] = dataSets
}

func (D Data) SetDataSetDefault(dataSets []DataSet) {
	data := []interface{}{}
	for _, i := range dataSets {
		data = append(data, i)
	}
	D.SetDataSet(data)
}

type OneDataSet map[string]interface{}

func (O OneDataSet) SetDefault(data DataSet) {
	O["label"] = data.Label
	O["backgroundColor"] = data.BackgroundColor
	O["borderColor"] = data.BorderColor
	O["data"] = data.Data
	O["fill"] = data.Fill
}

func (O OneDataSet) SetProperty(key string, v interface{}) {
	O[key] = v
}

type DataSet struct {
	Type                 string        `json:"type,omitempty"`
	Label                string        `json:"label,omitempty"`
	BackgroundColor      interface{}   `json:"backgroundColor,omitempty"`
	BorderColor          interface{}   `json:"borderColor,omitempty"`
	Data                 []interface{} `json:"data"`
	Fill                 interface{}   `json:"fill,omitempty"`
	BorderDash           []int         `json:"borderDash,omitempty"`
	PointBackgroundColor interface{}   `json:"pointBackgroundColor,omitempty"`
	BorderWidth          interface{}   `json:"borderWidth,omitempty"`
}

func defaultData() Data {
	var data Data
	data = make(map[string]interface{})
	return data
}

type Points struct {
	X interface{} `json:"x"`
	Y interface{} `json:"y"`
}
