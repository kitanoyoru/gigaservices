package utils

import "os"

func ReadFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(content)
}
