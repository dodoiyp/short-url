package utils

import (
	"errors"

	"github.com/lytics/base62"
)

func Base62Encode(rowData string) (string, error) {
	encodedStr := base62.StdEncoding.EncodeToString([]byte(rowData))
	if encodedStr == "" {
		return "", errors.New("base 62 problem")
	}
	return encodedStr, nil
}
