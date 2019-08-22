package charts

import "net/http"

type Bubble struct {
	Base
}

func (B Bubble) Plot(w http.ResponseWriter, r *http.Request) {
	toHandle(w, r, B.Base)
}
func (B Bubble) Save(name string) bool {
	if name == "" {
		name = "bubble"
	}
	return toLocal(name, B.Base)
}

func (B Bubble) Name() string {
	return B.Type
}
func (B Bubble) String() string {
	return BubbleType
}

func NewBubble(title string) Bubble {
	var option Options
	option = make(map[string]interface{})
	option.SetTitle(title)
	return Bubble{
		Base: Base{
			Type:    BubbleType,
			Data:    defaultData(),
			Options: option,
		},
	}
}
