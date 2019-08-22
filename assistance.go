package charts

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func toLocal(name string, v interface{}) bool {
	if name == "" {
		return false
	}
	name = fmt.Sprintf("%s.html", name)
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
		return false
	}
	tpl := template.Must(template.New("").Parse(PlotText()))
	err = tpl.Execute(file, v)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func toHandle(w http.ResponseWriter, req *http.Request, v interface{}) {
	tpl := template.Must(template.New("").Parse(PlotText()))
	err := tpl.Execute(w, v)
	if err != nil {
		log.Println(err)
		return
	}
}
