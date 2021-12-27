package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(i interface{}) {
	bytes, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson(true)
	logJson(1)
	logJson("oman")
	logJson([...]string{"nyoman", "pradipta", "dewantara"})
	logJson([]struct {
		Name string
		Age  bool
	}{
		{
			Name: "Oman",
		},
	})
}
