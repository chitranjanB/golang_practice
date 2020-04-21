package json

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func ConvertToJsonString(data interface{}) string {
	byteArray, err := json.Marshal(data)
	//byteArray, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	return string(byteArray)
}

func ConvertJsonToObject(jsonString string, object interface{}) {
	err := json.Unmarshal([]byte(jsonString), &object)
	if err != nil {
		fmt.Println(err)
	}
}
