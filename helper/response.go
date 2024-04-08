package helper

import "reflect"

func ResponseFormat(status int, message any, data ...any) (int, map[string]any) {
	result := map[string]any{
		"code":    status,
		"message": message,
	}

	if len(data) >= 1 {
		if data[0] != nil {
			result["data"] = data[0]
		}
	}

	if len(data) >= 2 {
		if reflect.ValueOf(data[1]).Kind() == reflect.Map {
			for key, value := range data[1].(map[string]any) {
				result[key] = value
			}
		}
	}
	return status, result
}
