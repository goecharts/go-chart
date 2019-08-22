package charts

import "net/http"

const (
	VERSION  = "0.1"
	RELEASE  = "0.1.1"
	AUTHOR   = "wuxiaoshen@shu.edu.cn"
	LOCATION = "ShangHai"
)

type Charts interface {
	Plot(writer http.ResponseWriter, request *http.Request)
	Save(string) bool
	Name() string
}

var DefaultWidth = 800
var DefaultHeight = 800

type Theme struct {
	Width  int  `json:"width"`
	Height int  `json:"height"`
	Base   Base `json:"base"`
}

type Base struct {
	Type    string `json:"type,omitempty"`
	Data    `json:"data"`
	Options `json:"options"`
}
