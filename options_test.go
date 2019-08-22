package charts

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestOptions(test *testing.T) {
	var o Options
	o = make(map[string]interface{})
	o.SetTitle("Learn ChartJS")
	fmt.Println(o)
	c, _ := json.Marshal(o)
	fmt.Println(string(c))
}
