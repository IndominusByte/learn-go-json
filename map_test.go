package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	streamString := `{"id": 1, "name": "mac book", "price": 20000}`
	streamByte := []byte(streamString)

	var results map[string]interface{}
	json.Unmarshal(streamByte, &results)

	fmt.Println(results)
	fmt.Println(results["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    1,
		"name":  "mac book",
		"price": 20000,
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}
