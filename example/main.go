package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/goecharts/go-chart"
)

func logger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		Message := fmt.Sprintf("%s | %s | %s | %s", req.Host, req.Method, req.RequestURI, time.Now().Format("2006-01-02 15:04:05"))
		log.Printf("%s", Message)
		next.ServeHTTP(w, req)
	})
}
func exampleLine() *charts.Line {
	line := charts.NewLine("Chart.js Line Chart")
	line.Data.SetLabels([]interface{}{"Jan", "Feb", "Mar", "April", "May", "June", "July"})

	dataset := []charts.DataSet{
		{
			Label:           "My First DataSet",
			Data:            []interface{}{11, 2, 30, 14, 5, 23, 7},
			BackgroundColor: charts.Purple(),
			BorderColor:     charts.Purple(),
		},
		{
			Label:           "My Second DataSet",
			Data:            []interface{}{12, 22, 10, 14, 5, 13, 6},
			BackgroundColor: charts.BlueAlpha(0.4),
			BorderColor:     charts.Blue(),
			Fill:            "origin",
			BorderDash:      []int{5, 5},
		},
	}
	one := charts.OneDataSet{}
	one.SetDefault(dataset[0])
	one.SetProperty("borderWidth", 2)

	line.SetDataSet([]interface{}{one, dataset[1]})

	line.SetXAxes([]string{"Month"})
	line.SetYAxes([]string{"Value"})
	return line
}

func exampleBar() charts.Bar {
	//bar := charts.NewHorizontalBar("Chart.js HorizontalBar Chart")
	bar := charts.NewBar("Chart.js Bar Chart")
	bar.SetLabels([]interface{}{"Jan", "Feb", "Mar", "Apr", "May", "June", "July"})
	dataset := []charts.DataSet{
		{
			Label:           "DataSet 1",
			BackgroundColor: charts.Red(),
			BorderColor:     charts.RedAlpha(0.7),
			Data:            []interface{}{-15, 51, 36, 24, -17, -64, 80},
		},
		{
			Label:           "DataSet 2",
			BorderColor:     charts.BlueAlpha(0.7),
			BackgroundColor: charts.BlueAlpha(0.7),
			Data:            []interface{}{25, -25, 35, 62, 35, -26, -24},
		},
	}
	bar.SetDataSetDefault(dataset)
	bar.SetLegend(charts.TOP, false)
	bar.SetXYAxes([]string{"Month"}, []string{"V"})
	return *bar
}

func exampleRadar() charts.Radar {
	radar := charts.NewRadar("Chart js Radar")
	radar.SetLabels([]interface{}{[]string{"Eating", "Dinner"},
		[]string{"Drinking", "Water"},
		"sleeping", []string{"Designing", "Graphics"},
		"Coding", "Cycling", "Running"})
	datasets := []charts.DataSet{
		{
			Label:                "# My First DataSet",
			BackgroundColor:      charts.RedAlpha(0.4),
			BorderColor:          charts.RedAlpha(0.4),
			Data:                 []interface{}{10, 20, 50, 20, 40, 30, 80, 60},
			PointBackgroundColor: charts.Red(),
		},
		{
			Label:                "# My Second DataSet",
			BorderColor:          charts.PurpleAlpha(0.4),
			BackgroundColor:      charts.PurpleAlpha(0.4),
			Data:                 []interface{}{90, 10, 30, 20, 40, 60, 20},
			PointBackgroundColor: charts.Purple(),
		},
	}
	radar.SetDataSetDefault(datasets)
	radar.SetLegend(charts.TOP, false)
	type scale struct {
		Ticks struct {
			BeginAtZero bool `json:"beginAtZero"`
		} `json:"ticks"`
	}
	a := scale{}
	a.Ticks.BeginAtZero = true
	radar.SetProperty("scales", a)
	return *radar
}

func exampleScatter() charts.Scatter {
	scatter := charts.NewScatter("Chart js scatter")
	datasets := []charts.DataSet{
		{
			Label: "# My First Data",
			Data: []interface{}{
				charts.Points{X: 0, Y: 1},
				charts.Points{X: 2, Y: 3},
				charts.Points{X: 3, Y: 5},
				charts.Points{X: 5, Y: 10},
			},
			BackgroundColor: charts.Red(),
			BorderColor:     charts.Red(),
		}, {
			Label: "# My Second Data",
			Data: []interface{}{
				charts.Points{X: 1, Y: 0},
				charts.Points{X: 2, Y: 1},
				charts.Points{X: 4, Y: 3},
				charts.Points{X: 3, Y: 6},
			},
			BorderColor:     charts.Purple(),
			BackgroundColor: charts.Purple(),
		},
	}
	scatter.SetDataSetDefault(datasets)
	scatter.SetXAxes([]string{"XAxes"})
	scatter.SetYAxes([]string{"YAxes"})
	return scatter
}

func examplePie() charts.Pie {
	pie := charts.NewPie("chart js pie")
	pie.SetLabels([]interface{}{"Red", "Orange", "Yellow", "Green", "Blue"})
	datasets := []charts.DataSet{
		{
			Label: "# My First Pie Data",
			BackgroundColor: []interface{}{
				charts.Red(),
				charts.Orange(),
				charts.Yellow(),
				charts.Green(),
				charts.Blue(),
			},
			Data: []interface{}{10, 20, 30, 40, 50},
		},
	}
	pie.SetDataSetDefault(datasets)
	pie.SetResponsive(true)
	return pie
}

func examplePolarArea() charts.PolarArea {
	polarArea := charts.NewPolarArea("chart js polarArea")
	polarArea.SetLabels([]interface{}{"Red", "Orange", "Yellow", "Green", "Blue"})
	dataSets := []charts.DataSet{
		{
			Data: []interface{}{86, 49, 83, 31, 36},
			BackgroundColor: []interface{}{
				charts.Red(), charts.Orange(), charts.Yellow(), charts.Green(), charts.Blue(),
			},
			Label: "# My First DataSet",
		},
	}
	polarArea.SetDataSetDefault(dataSets)
	polarArea.SetResponsive(true)
	polarArea.SetLegend(charts.RIGHT, false)
	return polarArea
}

func exampleDoughnut() charts.Doughnut {
	doughnut := charts.NewDoughnut("chart js doughnut")
	doughnut.SetLabels([]interface{}{"Red", "Orange", "Yellow", "Green", "Blue"})
	datasets := []charts.DataSet{
		{
			Data: []interface{}{64, 27, 42, 2, 94},
			BackgroundColor: []interface{}{
				charts.Red(), charts.Orange(), charts.Yellow(), charts.Green(), charts.Blue(),
			},
			Label: "# My First DataSet",
		},
	}
	doughnut.SetDataSetDefault(datasets)
	doughnut.SetResponsive(true)
	doughnut.SetLegend(charts.RIGHT, false)
	return doughnut

}

func main() {
	line := exampleLine()
	bar := exampleBar()
	radar := exampleRadar()
	scatter := exampleScatter()
	pie := examplePie()
	polarArea := examplePolarArea()
	doughnut := exampleDoughnut()

	//c, _ := json.MarshalIndent(scatter.Base, " ", " ")
	//fmt.Println(string(c))

	http.HandleFunc(fmt.Sprintf("/%s", line.Type), logger(line.Plot))
	http.HandleFunc(fmt.Sprintf("/%s", bar.Type), logger(bar.Plot))
	http.HandleFunc(fmt.Sprintf("/%s", radar.Type), logger(radar.Plot))
	http.HandleFunc(fmt.Sprintf("/%s", scatter.Type), logger(scatter.Plot))
	http.HandleFunc(fmt.Sprintf("/%s", pie.Type), logger(pie.Plot))
	http.HandleFunc(fmt.Sprintf("/%s", polarArea.Type), logger(polarArea.Plot))
	http.HandleFunc(fmt.Sprintf("/%s", doughnut.Type), logger(doughnut.Plot))
	log.Fatal(http.ListenAndServe(":9090", nil))

}
