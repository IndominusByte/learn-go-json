package learngojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("product.json")
	decoder := json.NewDecoder(reader)

	p := &Product{}
	decoder.Decode(p)

	fmt.Println(p)
}
