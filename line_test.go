package charts

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestLine(test *testing.T) {
	line := NewLine("Chart.js Line Chart")
	line.Data.SetLabels([]string{"Jan", "Feb", "Mar", "April", "May", "June", "July"})

	dataset := []DataSet{
		{
			Label:           "My First DataSet",
			Data:            []interface{}{11, 2, 30, 14, 5, 23, 7},
			BackgroundColor: Purple(),
			BorderColor:     Purple(),
		},
		{
			Label:           "My Second DataSet",
			Data:            []interface{}{12, 22, 10, 14, 5, 13, 6},
			BackgroundColor: Black(),
			BorderColor:     Black(),
			Fill:            false,
		},
	}
	one := OneDataSet{}
	one.SetDefault(dataset[0])
	one.SetProperty("borderWidth", 2)

	line.Data.SetDataSet([]interface{}{one, dataset[1]})

	line.Options.SetXAxes("Month")
	line.Options.SetYAxes("Value")
	//line.Options.SetXYAxes("Month", "Value")
	c, _ := json.MarshalIndent(line.Base, " ", " ")
	fmt.Println(string(c))
	http.HandleFunc("/line", line.Plot)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
