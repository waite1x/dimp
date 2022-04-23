package file

import (
	"encoding/json"
	"os"
)

func ReadJson(path string, v interface{}) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, v)
}
