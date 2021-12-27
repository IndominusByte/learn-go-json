package learngojson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoderStream(t *testing.T) {
	writer, _ := os.Create("product_out.json")
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	product := Product{
		Id:       1,
		Name:     "macbook pro",
		ImageUrl: "http://localhost/default.jpg",
	}

	encoder.Encode(product)
}
