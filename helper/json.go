package helper

import "encoding/json"

func FillStruct(source interface{}, dest interface{}) {
	byteSource, _ := json.Marshal(source)
	json.Unmarshal(byteSource, &dest)
}