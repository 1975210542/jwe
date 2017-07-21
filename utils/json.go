package utils

import (
	"encoding/json"
)

func JsonEncode(v interface{}) (string, error) {
	j, e := json.Marshal(v)
	return string(j), e
}

func JsonDecode(data string, v interface{}) error {
	if err := json.Unmarshal(([]byte)(data), &v); err != nil {
		return err
	}
	return nil
}
