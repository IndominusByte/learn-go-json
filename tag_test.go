package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestEncodeTag(t *testing.T) {
	bytes, err := json.Marshal(Product{
		Id:       2,
		Name:     "macbook pro",
		ImageUrl: "http://example.com/default.jpg",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestDecodeTag(t *testing.T) {
	streamString := `{"ID":2,"NAME":"macbook pro","IMAGE_URL":"http://example.com/default.jpg"}`
	streamByte := []byte(streamString)

	p := &Product{}

	if err := json.Unmarshal(streamByte, p); err != nil {
		panic(err)
	}

	fmt.Println(p)
}
