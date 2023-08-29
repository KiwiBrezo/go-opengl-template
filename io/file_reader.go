package io

import (
	"os"
)

func ReadShaderFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data) + "\x00", nil
}
