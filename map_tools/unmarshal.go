package maptools

import (
	"encoding/json"
)

func Unmarshal(in []byte) (map[string]any, error) {
	out := make(map[string]any)
	if err := json.Unmarshal(in, &out); err != nil {
		return nil, err
	}

	return out, nil
}
