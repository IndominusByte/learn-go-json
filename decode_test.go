package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

var Customer struct {
	FirstName string
	LastName  string
	Age       int
	Married   bool
}

func TestDecode(t *testing.T) {
	streamString := `{"FirstName": "oman", "LastName": "pradipta", "Age": 12, "lol": 12}`
	streamByte := []byte(streamString)

	c := &Customer

	if err := json.Unmarshal(streamByte, c); err != nil {
		panic(err)
	}

	fmt.Println(c)
}
