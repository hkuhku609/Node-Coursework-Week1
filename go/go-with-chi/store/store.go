package store

import "os"

type Quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func ReadJSONFile(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}
