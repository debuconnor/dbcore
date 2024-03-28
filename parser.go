package dbcore

import (
	"encoding/json"
)

func parseJson(data string) map[string]interface{} {
	var jsonData map[string]interface{}

	if err := json.Unmarshal([]byte(data), &jsonData); err == nil {
		return jsonData
	}

	Error(ERROR_CODE_PARSE_JSON)
	return nil
}
