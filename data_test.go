package charts

import (
	"encoding/json"
	"fmt"
	"testing"
)

type v struct {
	Value map[string]interface{}
}

func New() *v {
	return &v{
		Value: make(map[string]interface{}),
	}
}

func (V *v) SetValue(key string, v interface{}) {
	V.Value[key] = v
}

func TestV(test *testing.T) {
	v := New()
	v.SetValue("labels", []string{"123", "2", "3"})
	v.SetValue("data", struct {
		Name  string `json:"name"`
		Label string `json:"label"`
	}{
		Name:  "d",
		Label: "da",
	})
	c, _ := json.MarshalIndent(v.Value, " ", "")
	fmt.Println(string(c))
}

func TestData(tests *testing.T) {
	d := DefaultData
	d.SetLabels([]string{"1", "2", "3", "4"})
	var da OneDataSet
	da = make(map[string]interface{})
	da.SetDefault(DataSet{})

	d.SetDataSet([]interface{}{da})
	c, _ := json.MarshalIndent(d, " ", " ")
	fmt.Println(string(c))
}
