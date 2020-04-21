package json

import (
	"fmt"
	"testing"
)

func TestConvertToJsonString(t *testing.T) {
	data := Book{Title: "Learning Go", Author: "Somebody"}
	want := "{\"title\":\"Learning Go\",\"author\":\"Somebody\"}"
	if got := ConvertToJsonString(data); got != want {
		t.Errorf("ConvertToJsonString() = %v, want %v", got, want)
	}
}

func TestConvertJsonToObject(t *testing.T) {
	input := "{\"title\":\"Learning Go\",\"author\":\"Somebody\"}"
	book := Book{}
	ConvertJsonToObject(input, &book)
	fmt.Println(book.Title, ",", book.Author)
}
