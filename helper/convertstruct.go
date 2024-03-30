package helper

import "encoding/json"

func ConvertStruct(from, to interface{}) error {
	js, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(js, to)
}
