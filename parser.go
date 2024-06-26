package dbcore

import (
	"encoding/json"
	"errors"
)

func parseJson(data string) map[string]interface{} {
	var jsonData map[string]interface{}

	if err := json.Unmarshal([]byte(data), &jsonData); err == nil {
		return jsonData
	}

	Error(errors.New(ERROR_PARSE_JSON))
	return nil
}
