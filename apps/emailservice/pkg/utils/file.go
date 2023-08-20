package utils

import "io/ioutil"

func ReadFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(content)
}
