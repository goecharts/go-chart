package charts

import (
	"log"

	"github.com/gobuffalo/packr"
)

func PlotText() string {
	box := packr.NewBox(".")
	text, err := box.FindString("plot.html")
	if err != nil {
		log.Println(err)
		return ""
	}
	return text
}
